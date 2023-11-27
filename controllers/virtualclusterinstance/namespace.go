package virtualclusterinstance

import (
	"context"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *Reconciler) namespaceForVirtualClusterInstance(
	vci *loftv1.VirtualClusterInstance) (*corev1.Namespace, error) {
	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: normalizedNamespace(vci),
		},
	}

	// Set the ownerRef for the Namespace
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/
	if err := ctrl.SetControllerReference(vci, ns, r.Scheme); err != nil {
		return nil, err
	}

	return ns, nil
}

func (r *Reconciler) ensureNamespaceExists(
	ctx context.Context, vci *loftv1.VirtualClusterInstance) (ctrl.Result, error) {

	found := &corev1.Namespace{}
	err := r.Get(ctx, types.NamespacedName{Name: normalizedNamespace(vci)}, found)
	if err != nil && apierrors.IsNotFound(err) {
		ns, err := r.namespaceForVirtualClusterInstance(vci)
		if err != nil {
			r.Log.Error(err, "Failed to define new Namespace resource for VirtualClusterInstance")
			return ctrl.Result{}, err
		}
		r.Log.Info("Creating a new Namespace", "Namespace.Name", ns.Name)
		if err = r.Create(ctx, ns); err != nil {
			r.Log.Error(err, "Failed to create new Namespace", "Namespace.Name", ns.Name)
			return ctrl.Result{}, err
		}

		// Namespace created successfully
		// We will requeue the reconciliation so that we can ensure the state
		// and move forward for the next operations
		return ctrl.Result{RequeueAfter: defaultRequeuePeriod}, nil
	} else if err != nil {
		r.Log.Error(err, "Failed to get Namespace")
		// Let's return the error for the reconciliation be re-triggered again
		return ctrl.Result{}, err
	}

	// Proceed in reconcile loop.
	return ctrl.Result{}, nil
}
