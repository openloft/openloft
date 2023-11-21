/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package virtualclusterinstance

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	openloftv1 "github.com/openloft/openloft/api/v1"
	ctrlutils "github.com/openloft/openloft/pkg/controller/utils"
)

const (
	defaultRequeuePeriod    = 60 * time.Second
	defaultErrRequeuePeriod = 5 * time.Second
)

// Reconciler reconciles a VirtualClusterInstance object
type Reconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Log      logr.Logger
	Recorder record.EventRecorder
}

func generateNamespace(instance *loftv1.VirtualClusterInstance) string {
	return fmt.Sprintf("vc-%s", instance.Name)
}

//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclusterinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclusterinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclusterinstances/finalizers,verbs=update

//+kubebuilder:rbac:groups=storage.openloft.cn,resources=virtualclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=storage.openloft.cn,resources=virtualclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=storage.openloft.cn,resources=virtualclusters/finalizers,verbs=update

//+kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Reconciling VirtualClusterInstance", "req.NamespacedName", req.NamespacedName)

	vci := &loftv1.VirtualClusterInstance{}
	err := r.Client.Get(ctx, req.NamespacedName, vci)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.Log.Info("VirtualClusterInstance not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		r.Log.Error(err, "Failed to get VirtualClusterInstance")
		return ctrl.Result{}, err
	}

	if result, err := r.handleFinalizer(ctx, vci); ctrlutils.ShouldReturn(result, err) {
		return result, err
	}

	if result, err := r.ensureNamespaceExists(ctx, vci); ctrlutils.ShouldReturn(result, err) {
		return result, err
	}

	vcSpec, err := r.getVirtualClusterSpec(ctx, vci)
	if err != nil {
		r.Log.Error(err, "Failed to get VirtualClusterSpec")
		// Let's return the error for the reconciliation be re-triggered again
		return ctrl.Result{}, err
	}

	if result, err := r.ensureVirtualClusterExists(ctx, vci, vcSpec); ctrlutils.ShouldReturn(result, err) {
		return result, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&loftv1.VirtualClusterInstance{}).
		Owns(&corev1.Namespace{}).
		Owns(&openloftv1.VirtualCluster{}).
		Watches(
			&loftv1.VirtualClusterTemplate{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, o client.Object) []reconcile.Request {
				vct := o.(*loftv1.VirtualClusterTemplate)

				vciList := &loftv1.VirtualClusterInstanceList{}
				if err := mgr.GetClient().List(ctx, vciList); err != nil {
					r.Log.Error(err, "Failed to list VirtualClusterInstances")
					return []reconcile.Request{}
				}

				var requests []reconcile.Request
				for _, vci := range vciList.Items {
					if vci.Spec.TemplateRef.Name == vct.Name {
						r.Log.Info("Enqueueing VirtualClusterInstance", "VirtualClusterInstance.Name", vci.Name)
						requests = append(requests, reconcile.Request{NamespacedName: types.NamespacedName{Name: vci.Name}})
					}
				}

				return requests
			}),
			// Watch for NodeClaim deletion events
			builder.WithPredicates(predicate.Funcs{
				CreateFunc: func(e event.CreateEvent) bool { return false },
				UpdateFunc: func(e event.UpdateEvent) bool { return true },
				DeleteFunc: func(e event.DeleteEvent) bool { return false },
			}),
		).
		Complete(r)
}
