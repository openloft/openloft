package v1

import corev1 "k8s.io/api/core/v1"

/* Copied from https://github.com/loft-sh/vcluster/blob/main/charts/k3s/values.yaml#L135

# Syncer configuration
syncer:
  # Image to use for the syncer
  # image: ghcr.io/loft-sh/vcluster
  extraArgs: []
  env: []
  livenessProbe:
    enabled: true
  readinessProbe:
    enabled: true
  volumeMounts:
    - mountPath: /data
      name: data
      # readOnly: true
  extraVolumeMounts: []
  resources:
    limits:
      cpu: 1000m
      memory: 512Mi
    requests:
      # ensure that cpu/memory requests are high enough.
      # for example gke wants minimum 10m/32Mi here!
      cpu: 20m
      memory: 64Mi
  kubeConfigContextName: "my-vcluster"
  serviceAnnotations: {}
*/

type Syncer struct {
	ExtraArgs []string        `json:"extraArgs,omitempty"`
	Env       []corev1.EnvVar `json:"env,omitempty"`
}
