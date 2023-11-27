package virtualclusterinstance

import (
	"context"

	openloftv1 "github.com/openloft/openloft/api/v1"
)

func (r *Reconciler) getClusterDomain(ctx context.Context) *openloftv1.ClusterDomain {
	clusterDomainList := &openloftv1.ClusterDomainList{}
	err := r.List(ctx, clusterDomainList)
	if err != nil {
		r.Log.Error(err, "Failed to list ClusterDomains")
		return nil
	}

	if len(clusterDomainList.Items) == 0 {
		r.Log.Error(err, "No ClusterDomain found")
		return nil
	}

	return &clusterDomainList.Items[0]
}
