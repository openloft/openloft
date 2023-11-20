package virtualclusterinstance

import (
	"context"
	"fmt"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	finalizer = "finalizer.virtualclusterinstance.storage.openloft.cn"
)

func (r *Reconciler) handleFinalizer(ctx context.Context, vci *loftv1.VirtualClusterInstance) (ctrl.Result, error) {
	if vci.GetDeletionTimestamp() != nil {
		if controllerutil.ContainsFinalizer(vci, finalizer) {
			r.Log.Info("Performing Finalizer Operations for VirtualClusterInstance before delete CR")

			// Perform all operations required before remove the finalizer and allow
			// the Kubernetes API to remove the custom resource.
			r.doFinalizerOperationsForVirtualClusterInstance(vci)

			r.Log.Info("Removing Finalizer for VirtualClusterInstance after successfully perform the operations")
			if ok := controllerutil.RemoveFinalizer(vci, finalizer); !ok {
				r.Log.Info("Failed to remove finalizer into the custom resource")
				return ctrl.Result{Requeue: true}, nil
			}

			if err := r.Update(ctx, vci); err != nil {
				r.Log.Error(err, "Failed to update custom resource to remove finalizer")
				return ctrl.Result{RequeueAfter: defaultErrRequeuePeriod}, err
			}
		}

		// Requeue until the object was properly deleted by Kubernetes
		return ctrl.Result{RequeueAfter: defaultRequeuePeriod}, nil
	}

	// Let's add a finalizer. Then, we can define some operations which should
	// occurs before the custom resource to be deleted.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/finalizers
	if !controllerutil.ContainsFinalizer(vci, finalizer) {
		r.Log.Info("Adding Finalizer for VirtualClusterInstance")
		if ok := controllerutil.AddFinalizer(vci, finalizer); !ok {
			r.Log.Info("Failed to add finalizer into the custom resource")
			return ctrl.Result{Requeue: true}, nil
		}

		if err := r.Update(ctx, vci); err != nil {
			r.Log.Error(err, "Failed to update custom resource to add finalizer")
			return ctrl.Result{RequeueAfter: defaultErrRequeuePeriod}, err
		}

		// Requeue until the object was properly deleted by Kubernetes
		return ctrl.Result{RequeueAfter: defaultRequeuePeriod}, nil
	}

	// Proceed in reconcile loop.
	return ctrl.Result{}, nil
}

// doFinalizerOperationsForVirtualClusterInstance will perform the required operations before delete the CR.
func (r *Reconciler) doFinalizerOperationsForVirtualClusterInstance(cr *loftv1.VirtualClusterInstance) {
	// TODO(user): Add the cleanup steps that the operator
	// needs to do before the CR can be deleted. Examples
	// of finalizers include performing backups and deleting
	// resources that are not owned by this CR, like a PVC.

	// Note: It is not recommended to use finalizers with the purpose of delete resources which are
	// created and managed in the reconciliation. These ones, such as the Deployment created on this reconcile,
	// are defined as depended of the custom resource. See that we use the method ctrl.SetControllerReference.
	// to set the ownerRef which means that the Deployment will be deleted by the Kubernetes API.
	// More info: https://kubernetes.io/docs/tasks/administer-cluster/use-cascading-deletion/

	// The following implementation will raise an event
	r.Recorder.Event(cr, "Warning", "Deleting",
		fmt.Sprintf("Custom Resource %s is being deleted from the namespace %s",
			cr.Name,
			cr.Namespace))
}
