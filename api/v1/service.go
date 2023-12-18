package v1

import corev1 "k8s.io/api/core/v1"

/* Copied from https://github.com/loft-sh/vcluster/blob/main/charts/k3s/values.yaml#L273

# Service configurations
service:
  type: ClusterIP

  # Optional configuration
  # A list of IP addresses for which nodes in the cluster will also accept traffic for this service.
  # These IPs are not managed by Kubernetes; e.g., an external load balancer.
  externalIPs: []

  # Optional configuration for LoadBalancer & NodePort service types
  # Route external traffic to node-local or cluster-wide endpoints [ Local | Cluster ]
  externalTrafficPolicy: ""

  # Optional configuration for LoadBalancer service type
  # Specify IP of load balancer to be created
  loadBalancerIP: ""
  # CIDR block(s) for the service allowlist
  loadBalancerSourceRanges: []
  # Set the loadBalancerClass if using an external load balancer controller
  loadBalancerClass: ""
*/

type Service struct {
	//+kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer;ExternalName
	Type        corev1.ServiceType `json:"type,omitempty"`
	ExternalIPs []string           `json:"externalIPs,omitempty"`
	//+kubebuilder:validation:Enum=Local;Cluster
	ExternalTrafficPolicy    corev1.ServiceExternalTrafficPolicy `json:"externalTrafficPolicy,omitempty"`
	LoadBalancerIP           string                              `json:"loadBalancerIP,omitempty"`
	LoadBalancerSourceRanges []string                            `json:"loadBalancerSourceRanges,omitempty"`
	LoadBalancerClass        *string                             `json:"loadBalancerClass,omitempty"`
}
