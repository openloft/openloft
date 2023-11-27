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

package clusterdomain

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"
	openloftv1 "github.com/openloft/openloft/api/v1"
)

// ClusterDomainReconciler reconciles a ClusterDomain object
type ClusterDomainReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Log      logr.Logger
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=storage.openloft.cn,resources=clusterdomains,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=storage.openloft.cn,resources=clusterdomains/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=storage.openloft.cn,resources=clusterdomains/finalizers,verbs=update

func (r *ClusterDomainReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Reconciling VirtualClusterTemplateReconciler", "req.NamespacedName", req.NamespacedName)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ClusterDomainReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&openloftv1.ClusterDomain{}).
		Complete(r)
}
