package v1

import networkingv1 "k8s.io/api/networking/v1"

/* Coped from https://github.com/loft-sh/vcluster/blob/main/charts/k3s/values.yaml#L294

# Configure the ingress resource that allows you to access the vcluster
ingress:
  # Enable ingress record generation
  enabled: false
  # Ingress path type
  pathType: ImplementationSpecific
  ingressClassName: ""
  host: vcluster.local
  annotations:
    nginx.ingress.kubernetes.io/backend-protocol: HTTPS
    nginx.ingress.kubernetes.io/ssl-passthrough: "true"
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
  # Ingress TLS configuration
  tls: []
    # - secretName: tls-vcluster.local
    #   hosts:
    #     - vcluster.local
*/

// Ingress defines the configuration for ingress
type Ingress struct {
	Enabled          bool                      `json:"enabled,omitempty"`
	PathType         networkingv1.PathType     `json:"pathType,omitempty"`
	APIVersion       string                    `json:"apiVersion,omitempty"`
	IngressClassName *string                   `json:"ingressClassName,omitempty"`
	Host             string                    `json:"host,omitempty"`
	Annotations      map[string]string         `json:"annotations,omitempty"`
	TLS              []networkingv1.IngressTLS `json:"tls,omitempty"`
}
