package v1

import corev1 "k8s.io/api/core/v1"

/* Copied from https://github.com/loft-sh/vcluster/blob/main/charts/k3s/values.yaml#L162

# Virtual Cluster (k3s) configuration
vcluster:
  # Image to use for the virtual cluster
  image: rancher/k3s:v1.28.2-k3s1
  imagePullPolicy: ""
  command:
    - /k3s-binary/k3s
  baseArgs:
    - server
    - --write-kubeconfig=/data/k3s-config/kube-config.yaml
    - --data-dir=/data
    - --disable=traefik,servicelb,metrics-server,local-storage,coredns
    - --disable-network-policy
    - --disable-agent
    - --disable-cloud-controller
    - --egress-selector-mode=disabled
    - --flannel-backend=none
    - --kube-apiserver-arg=bind-address=127.0.0.1
  extraArgs: []
  # Deprecated: Use syncer.extraVolumeMounts instead
  extraVolumeMounts: []
  # Deprecated: Use syncer.volumeMounts instead
  volumeMounts: []
  # Deprecated: Use syncer.env instead
  env: []
  resources:
    limits:
      cpu: 100m
      memory: 128Mi
    requests:
      cpu: 20m
      memory: 64Mi
*/

// VCluster defines the configuration for vCluster
type VCluster struct {
	// XXX: Supported K3s versions are 1.25, 1.26, 1.27 and 1.28
	// See available images at https://github.com/loft-sh/vcluster/blob/main/pkg/values/k3s.go#L12
	Image string          `json:"image,omitempty"`
	Env   []corev1.EnvVar `json:"env,omitempty"`
}
