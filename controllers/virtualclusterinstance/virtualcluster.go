package virtualclusterinstance

import (
	"context"
	"reflect"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"

	openloftv1 "github.com/openloft/openloft/api/v1"
)

func (r *Reconciler) virtualClusterForVirtualClusterInstance(
	vci *loftv1.VirtualClusterInstance) (*openloftv1.VirtualCluster, error) {
	vc := &openloftv1.VirtualCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      normalizedName(vci),
			Namespace: normalizedNamespace(vci),
		},
	}

	// Set the ownerRef for the VirtualCluster
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/
	if err := ctrl.SetControllerReference(vci, vc, r.Scheme); err != nil {
		return nil, err
	}

	return vc, nil
}

func (r *Reconciler) ensureVirtualClusterExists(
	ctx context.Context, vci *loftv1.VirtualClusterInstance,
	vcSpec *openloftv1.VirtualClusterSpec) (ctrl.Result, error) {

	found := &openloftv1.VirtualCluster{}
	err := r.Get(ctx, types.NamespacedName{Name: normalizedName(vci), Namespace: normalizedNamespace(vci)}, found)
	if err != nil && apierrors.IsNotFound(err) {
		vc, err := r.virtualClusterForVirtualClusterInstance(vci)
		if err != nil {
			r.Log.Error(err, "Failed to define new VirtualCluster resource for VirtualClusterInstance")
			return ctrl.Result{}, err
		}
		vc.Spec = *vcSpec
		r.Log.Info("Creating a new VirtualCluster",
			"VirtualCluster.Namespace", vc.Namespace, "VirtualCluster.Name", vc.Name)
		if err = r.Create(ctx, vc); err != nil {
			r.Log.Error(err, "Failed to create new VirtualCluster",
				"VirtualCluster.Namespace", vc.Namespace, "VirtualCluster.Name", vc.Name)
			return ctrl.Result{}, err
		}

		// VirtualCluster created successfully
		// We will requeue the reconciliation so that we can ensure the state
		// and move forward for the next operations
		return ctrl.Result{RequeueAfter: defaultRequeuePeriod}, nil
	} else if err != nil {
		r.Log.Error(err, "Failed to get VirtualCluster")
		// Let's return the error for the reconciliation be re-triggered again
		return ctrl.Result{}, err
	}

	if !reflect.DeepEqual(*vcSpec, found.Spec) {
		r.Log.Info("Updating VirtualCluster", "VirtualCluster.Namespace", found.Namespace, "VirtualCluster.Name", found.Name)
		found.Spec = *vcSpec
		if err := r.Update(ctx, found); err != nil {
			r.Log.Error(err, "Failed to update VirtualCluster")
			return ctrl.Result{}, err
		}
	} else {
		r.Log.Info("Skip updating VirtualCluster", "VirtualCluster.Namespace", found.Namespace, "VirtualCluster.Name", found.Name)
	}

	// Proceed in reconcile loop.
	return ctrl.Result{}, nil
}
