package virtualclusterinstance

import (
	"context"
	"fmt"
	"time"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
)

func (r *Reconciler) defaultIngressClass(ctx context.Context) (*networkingv1.IngressClass, error) {
	ingressClasses := &networkingv1.IngressClassList{}
	err := r.List(ctx, ingressClasses)
	if err != nil {
		r.Log.Error(err, "Failed to list IngressClasses")
		return nil, err
	}

	if len(ingressClasses.Items) == 0 {
		return nil, fmt.Errorf("no IngressClass found")
	}

	for _, ingressClass := range ingressClasses.Items {
		if ingressClass.Annotations["ingressclass.kubernetes.io/is-default-class"] == "true" {
			return &ingressClass, nil
		}
	}

	// Incase no default ingress class is found, return the first one
	return &ingressClasses.Items[0], nil
}

func (r *Reconciler) discoverHostFromService(ctx context.Context, vci *loftv1.VirtualClusterInstance) (string, error) {
	name := normalizedName(vci)
	namespace := normalizedNamespace(vci)

	host := ""
	err := wait.PollImmediate(time.Second*2, time.Second*10, func() (done bool, err error) {
		service := &corev1.Service{}
		err = r.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, service)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return true, nil
			}

			return false, err
		}

		// not a load balancer? Then don't wait
		if service.Spec.Type != corev1.ServiceTypeLoadBalancer {
			return true, nil
		}

		if len(service.Status.LoadBalancer.Ingress) == 0 {
			// Waiting for vcluster LoadBalancer ip
			return false, nil
		}

		if service.Status.LoadBalancer.Ingress[0].Hostname != "" {
			host = service.Status.LoadBalancer.Ingress[0].Hostname
		} else if service.Status.LoadBalancer.Ingress[0].IP != "" {
			host = service.Status.LoadBalancer.Ingress[0].IP
		}

		if host == "" {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return "", fmt.Errorf("can not get vcluster service: %v", err)
	}

	if host == "" {
		host = fmt.Sprintf("%s.%s.svc", name, namespace)
	}
	return host, nil
}
