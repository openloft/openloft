package utils

import "sigs.k8s.io/controller-runtime/pkg/reconcile"

// ShouldReturn returns if we should stop the reconcile loop based on result
func ShouldReturn(result reconcile.Result, err error) bool {
	return err != nil || result.Requeue || result.RequeueAfter > 0
}
