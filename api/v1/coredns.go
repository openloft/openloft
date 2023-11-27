package v1

import corev1 "k8s.io/api/core/v1"

/* Copied from https://github.com/loft-sh/vcluster/blob/main/charts/k3s/values.yaml#L337

# If enabled will deploy the coredns configmap
coredns:
  integrated: false
  plugin:
    enabled: false
    config: []
    # example configuration for plugin syntax, will be documented in detail
    # - record:
    #     fqdn: google.com
    #   target:
    #     mode: url
    #     url: google.co.in
    # - record:
    #     service: my-namespace/my-svc    # dns-test/nginx-svc
    #   target:
    #     mode: host
    #     service: dns-test/nginx-svc
    # - record:
    #     service: my-namespace-lb/my-svc-lb
    #   target:
    #     mode: host
    #     service: dns-test-exposed-lb/nginx-svc-exposed-lb
    # - record:
    #     service: my-ns-external-name/my-svc-external-name
    #   target:
    #     mode: host
    #     service: dns-test-external-name/nginx-svc-external-name
    # - record:
    #     service: my-ns-in-vcluster/my-svc-vcluster
    #   target:
    #     mode: vcluster              # can be tested only manually for now
    #     vcluster: test-vcluster-ns/test-vcluster
    #     service: dns-test-in-vcluster-ns/test-in-vcluster-service
    # - record:
    #     service: my-ns-in-vcluster-mns/my-svc-mns
    #   target:
    #     mode: vcluster              # can be tested only manually for now
    #     service: dns-test-in-vcluster-mns/test-in-vcluster-svc-mns
    #     vcluster: test-vcluster-ns-mns/test-vcluster-mns
    # - record:
    #     service: my-self-vc-ns/my-self-vc-svc
    #   target:
    #     mode: self
    #     service: dns-test/nginx-svc
  enabled: true
  replicas: 1
  # The nodeSelector example below specifices that coredns should only be scheduled to nodes with the arm64 label
  # nodeSelector:
  #   kubernetes.io/arch: arm64
  # image: my-core-dns-image:latest
  # config: |-
  #   .:1053 {
  #      ...
  # CoreDNS service configurations
  service:
    type: ClusterIP
    # Configuration for LoadBalancer service type
    externalIPs: []
    externalTrafficPolicy: ""
    # Extra Annotations
    annotations: {}
  resources:
    limits:
      cpu: 1000m
      memory: 170Mi
    requests:
      cpu: 3m
      memory: 16Mi
# if below option is configured, it will override the coredns manifests with the following string
#  manifests: |-
#    apiVersion: ...
#    ...
  podAnnotations: {}
  podLabels: {}
*/

// CoreDNS defines the configuration for CoreDNS
type CoreDNS struct {
	Enabled        bool                         `json:"enabled,omitempty"`
	Integrated     bool                         `json:"integrated,omitempty"`
	Plugin         *CoreDNSPluginConfig         `json:"plugin,omitempty"`
	PodAnnotations map[string]string            `json:"podAnnotations,omitempty"`
	PodLabels      map[string]string            `json:"podLabels,omitempty"`
	Replicas       int                          `json:"replicas,omitempty"`
	Resources      *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// CoreDNSPluginConfig defines the configuration for CoreDNS plugins
type CoreDNSPluginConfig struct {
	Enabled bool `json:"enabled,omitempty"`
}
