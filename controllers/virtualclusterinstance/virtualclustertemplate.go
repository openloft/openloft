package virtualclusterinstance

import (
	"context"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	openloftv1 "github.com/openloft/openloft/api/v1"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/types"
)

func (r *Reconciler) getVirtualClusterSpec(
	ctx context.Context, vci *loftv1.VirtualClusterInstance) (*openloftv1.VirtualClusterSpec, error) {

	found := &loftv1.VirtualClusterTemplate{}
	if err := r.Client.Get(ctx, types.NamespacedName{Name: vci.Spec.TemplateRef.Name}, found); err != nil {
		r.Log.Error(err, "Failed to get VirtualClusterTemplate")
		// Let's return the error for the reconciliation be re-triggered again
		return nil, err
	}

	vcSpec := &openloftv1.VirtualClusterSpec{}
	if err := yaml.Unmarshal([]byte(found.Spec.Template.HelmRelease.Values), vcSpec); err != nil {
		r.Log.Error(err, "Failed to unmarshal VirtualClusterSpec")
		// Let's return the error for the reconciliation be re-triggered again
		return nil, err
	}

	return vcSpec, nil
}
