package virtualclusterinstance

import (
	"context"
	"fmt"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"

	openloftv1 "github.com/openloft/openloft/api/v1"
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

	extraArgs := []string{
		fmt.Sprintf("--tls-san=%s", vci.Name),
		fmt.Sprintf("--out-kube-config-server=https://%s", vci.Name),
	}

	if vcSpec.Syncer != nil {
		vcSpec.Syncer.ExtraArgs = append(vcSpec.Syncer.ExtraArgs, extraArgs...)
	} else {
		vcSpec.Syncer = &openloftv1.SyncerConfig{
			ExtraArgs: extraArgs,
		}
	}

	return vcSpec, nil
}
