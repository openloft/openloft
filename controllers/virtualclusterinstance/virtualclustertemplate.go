package virtualclusterinstance

import (
	"context"
	"fmt"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	networkingv1 "k8s.io/api/networking/v1"
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

	if vcSpec.Ingress != nil && vcSpec.Ingress.Enabled {
		ingressClass, err := r.defaultIngressClass(ctx)
		if err != nil {
			r.Log.Error(err, "Failed to get default IngressClass")
			return nil, err
		}

		// https://www.vcluster.com/docs/using-vclusters/access#via-ingress
		pathType := networkingv1.PathTypeImplementationSpecific

		annotations := map[string]string{
			fmt.Sprintf("%s.ingress.kubernetes.io/backend-protocol", ingressClass.Name):   "HTTPS",
			fmt.Sprintf("%s.ingress.kubernetes.io/force-ssl-redirect", ingressClass.Name): "true",
			fmt.Sprintf("%s.ingress.kubernetes.io/ssl-passthrough", ingressClass.Name):    "true",
			"kubernetes.io/ingress.class": ingressClass.Name,
		}

		vcSpec.Ingress.IngressClassName = &ingressClass.Name
		vcSpec.Ingress.APIVersion = ingressClass.APIVersion
		vcSpec.Ingress.PathType = pathType
		vcSpec.Ingress.Annotations = annotations
		vcSpec.Ingress.TLS = []networkingv1.IngressTLS{}

		// XXX: HOST
		vcSpec.Ingress.Host = vci.Name
		if vcSpec.Ingress.Host == "" {
			host, err := r.discoverHostFromService(ctx, vci)
			if err != nil {
				return nil, err
			}
			// write the discovered host back
			vcSpec.Ingress.Host = host
		}

		if vcSpec.Syncer == nil {
			vcSpec.Syncer = &openloftv1.Syncer{}
		}
		vcSpec.Syncer.ExtraArgs = append(vcSpec.Syncer.ExtraArgs,
			fmt.Sprintf("--tls-san=%s", vcSpec.Ingress.Host),
			fmt.Sprintf("--out-kube-config-server=https://%s", vcSpec.Ingress.Host),
			"--enforce-node-selector=true",
		)
	}

	return vcSpec, nil
}
