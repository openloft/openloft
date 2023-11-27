package v1

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
}
