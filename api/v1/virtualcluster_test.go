package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"
)

func TestSpec(t *testing.T) {
	in := "# Below you can configure the virtual cluster\nisolation:\n  enabled: true\n  resourceQuota:\n    enabled: true\n    quota:\n      requests.cpu: 10\n      requests.memory: 20Gi\n      requests.storage: \"100Gi\"\n      requests.ephemeral-storage: 60Gi\n      requests.nvidia.com/gpu: 2\n      limits.nvidia.com/gpu: 2\n      limits.cpu: 20\n      limits.memory: 40Gi\n      limits.ephemeral-storage: 160Gi\n      services.nodeports: 0\n      services.loadbalancers: 1\n      count/endpoints: 40\n      count/pods: 20\n      count/services: 20\n      count/secrets: 100\n      count/configmaps: 100\n      count/persistentvolumeclaims: 20\n    scopeSelector:\n      matchExpressions:\n    scopes:\n\nsync:\n  ingresses:\n    enabled: true\n\ncoredns:\n  enabled: true\n  integrated: false\n\n# Checkout https://vcluster.com/pro/docs/ for more config options"
	out := &VirtualClusterSpec{}
	err := yaml.Unmarshal([]byte(in), out)
	require.NoError(t, err)

	require.True(t, out.Isolation.Enabled)
	require.NotEmpty(t, out.Isolation.ResourceQuota)
	require.True(t, out.Isolation.ResourceQuota.Enabled)
}
