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

package virtualclustertemplate

import (
	"context"

	"github.com/go-logr/logr"
	v1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Reconciler reconciles a VirtualClusterTemplate object
type Reconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Log      logr.Logger
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclustertemplates,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclustertemplates/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclustertemplates/finalizers,verbs=update

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("Reconciling VirtualClusterTemplateReconciler", "req.NamespacedName", req.NamespacedName)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.VirtualClusterTemplate{}).
		Complete(r)
}
