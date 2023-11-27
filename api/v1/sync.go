package v1

/* Copied from https://github.com/loft-sh/vcluster/blob/main/charts/k3s/values.yaml#L27

# Resource syncers that should be enabled/disabled.
# Enabling syncers will impact RBAC Role and ClusterRole permissions.
# To disable a syncer set "enabled: false".
# See docs for details - https://www.vcluster.com/docs/architecture/synced-resources
sync:
  services:
    enabled: true
  configmaps:
    enabled: true
    all: false
  secrets:
    all: false
    enabled: true
  endpoints:
    enabled: true
  pods:
    enabled: true
    ephemeralContainers: false
    status: false
  events:
    enabled: true
  persistentvolumeclaims:
    enabled: true
  ingresses:
    enabled: false
  ingressclasses: {}
    # By default IngressClasses sync is enabled when the Ingress sync is enabled
    # but it can be explicitly disabled by setting:
    # enabled: false
  fake-nodes:
    enabled: true # will be ignored if nodes.enabled = true
  fake-persistentvolumes:
    enabled: true # will be ignored if persistentvolumes.enabled = true
  nodes:
    fakeKubeletIPs: true
    enabled: false
    # If nodes sync is enabled, and syncAllNodes = true, the virtual cluster
    # will sync all nodes instead of only the ones where some pods are running.
    syncAllNodes: false
    # nodeSelector is used to limit which nodes get synced to the vcluster,
    # and which nodes are used to run vcluster pods.
    # A valid string representation of a label selector must be used.
    nodeSelector: ""
    # if true, vcluster will run with a scheduler and node changes are possible
    # from within the virtual cluster. This is useful if you would like to
    # taint, drain and label nodes from within the virtual cluster
    enableScheduler: false
    # DEPRECATED: use enable scheduler instead
    # syncNodeChanges allows vcluster user edits of the nodes to be synced down to the host nodes.
    # Write permissions on node resource will be given to the vcluster.
    syncNodeChanges: false
  persistentvolumes:
    enabled: false
  storageclasses:
    enabled: false
  # formerly named - "legacy-storageclasses"
  hoststorageclasses:
    enabled: false
  priorityclasses:
    enabled: false
  networkpolicies:
    enabled: false
  volumesnapshots:
    enabled: false
  poddisruptionbudgets:
    enabled: false
  serviceaccounts:
    enabled: false
  # generic CRD configuration
  generic:
    config: |-
      ---
*/

// Sync defines the configuration for sync
type Sync struct {
	Ingresses *Ingresses `json:"ingresses,omitempty"`
	Nodes     *Nodes     `json:"nodes,omitempty"`
}

// Ingresses defines the configuration for ingress
type Ingresses struct {
	Enabled bool `json:"enabled,omitempty"`
}

type Nodes struct {
	// Enabled controls whether the nodes are Fake or Real
	Enabled      bool   `json:"enabled,omitempty"`
	NodeSelector string `json:"nodeSelector,omitempty"`
}
