package v1

import corev1 "k8s.io/api/core/v1"

/* Coped from https://github.com/loft-sh/vcluster/blob/main/charts/k3s/values.yaml#L412

# If enabled will deploy vcluster in an isolated mode with pod security
# standards, limit ranges and resource quotas
isolation:
  enabled: false
  namespace: null

  podSecurityStandard: baseline

  # If enabled will add node/proxy permission to the cluster role
  # in isolation mode
  nodeProxyPermission:
    enabled: false

  resourceQuota:
    enabled: true
    quota:
      requests.cpu: 10
      requests.memory: 20Gi
      requests.storage: "100Gi"
      requests.ephemeral-storage: 60Gi
      limits.cpu: 20
      limits.memory: 40Gi
      limits.ephemeral-storage: 160Gi
      services.nodeports: 0
      services.loadbalancers: 1
      count/endpoints: 40
      count/pods: 20
      count/services: 20
      count/secrets: 100
      count/configmaps: 100
      count/persistentvolumeclaims: 20
    scopeSelector:
      matchExpressions:
    scopes:

  limitRange:
    enabled: true
    default:
      ephemeral-storage: 8Gi
      memory: 512Mi
      cpu: "1"
    defaultRequest:
      ephemeral-storage: 3Gi
      memory: 128Mi
      cpu: 100m

  networkPolicy:
    enabled: true
    outgoingConnections:
      ipBlock:
        cidr: 0.0.0.0/0
        except:
          - 100.64.0.0/10
          - 127.0.0.0/8
          - 10.0.0.0/8
          - 172.16.0.0/12
          - 192.168.0.0/16
*/

// IsolationConfig defines the configuration for isolation
type IsolationConfig struct {
	Enabled   bool    `json:"enabled,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	//+kubebuilder:validation:Enum=privileged;baseline;restricted
	PodSecurityStandard *string              `json:"podSecurityStandard,omitempty"`
	NodeProxyPermission *NodeProxyPermission `json:"nodeProxyPermission,omitempty"`
	ResourceQuota       *ResourceQuota       `json:"resourceQuota,omitempty"`
}

type NodeProxyPermission struct {
	Enabled bool `json:"enabled,omitempty"`
}
type ResourceQuota struct {
	Enabled       bool                        `json:"enabled,omitempty"`
	Quota         corev1.ResourceList         `json:"quota,omitempty"`
	ScopeSelector *corev1.ScopeSelector       `json:"scopeSelector,omitempty"`
	Scopes        []corev1.ResourceQuotaScope `json:"scopes,omitempty"`
}
