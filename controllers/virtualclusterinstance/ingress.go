package virtualclusterinstance

import (
	"context"
	"fmt"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *Reconciler) defaultIngressClassName(ctx context.Context) (string, error) {
	ingressClasses := &networkingv1.IngressClassList{}
	err := r.List(ctx, ingressClasses)
	if err != nil {
		r.Log.Error(err, "Failed to list IngressClasses")
		return "", err
	}

	if len(ingressClasses.Items) == 0 {
		return "", fmt.Errorf("no IngressClass found")
	}

	for _, ingressClass := range ingressClasses.Items {
		if ingressClass.Annotations["ingressclass.kubernetes.io/is-default-class"] == "true" {
			return ingressClass.Name, nil
		}
	}

	// Incase no default ingress class is found, return the first one
	return ingressClasses.Items[0].Name, nil
}

func (r *Reconciler) ingressForVirtualClusterInstance(
	vci *loftv1.VirtualClusterInstance, ingressClassName *string) (*networkingv1.Ingress, error) {

	// https://www.vcluster.com/docs/using-vclusters/access#via-ingress
	pathType := networkingv1.PathTypeImplementationSpecific

	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      generateIngressName(vci),
			Namespace: generateNamespace(vci),
			Annotations: map[string]string{
				fmt.Sprintf("%s.ingress.kubernetes.io/backend-protocol", *ingressClassName):   "HTTPS",
				fmt.Sprintf("%s.ingress.kubernetes.io/force-ssl-redirect", *ingressClassName): "true",
				fmt.Sprintf("%s.ingress.kubernetes.io/ssl-passthrough", *ingressClassName):    "true",
				"kubernetes.io/ingress.class": *ingressClassName,
			},
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: ingressClassName,
			Rules: []networkingv1.IngressRule{
				{
					// TODO: How to get the domain name?
					Host: fmt.Sprintf("%s.%s", vci.Name, "openloft.cn"),
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: vci.Name,
											Port: networkingv1.ServiceBackendPort{
												Number: 443,
											},
										},
									},
									Path:     "/",
									PathType: &pathType,
								},
							},
						},
					},
				},
			},
		},
	}

	// Set the ownerRef for the Ingress
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/
	if err := ctrl.SetControllerReference(vci, ingress, r.Scheme); err != nil {
		return nil, err
	}

	return ingress, nil
}

func (r *Reconciler) ensureIngressExists(
	ctx context.Context, vci *loftv1.VirtualClusterInstance) (ctrl.Result, error) {

	found := &networkingv1.Ingress{}
	err := r.Get(ctx, types.NamespacedName{Name: generateIngressName(vci), Namespace: generateNamespace(vci)}, found)
	if err != nil && apierrors.IsNotFound(err) {
		ingressClassName, err := r.defaultIngressClassName(ctx)
		if err != nil {
			r.Log.Error(err, "Failed to get default IngressClass")
			return ctrl.Result{}, err
		}
		ingress, err := r.ingressForVirtualClusterInstance(vci, &ingressClassName)
		if err != nil {
			r.Log.Error(err, "Failed to define new Ingress resource for VirtualClusterInstance")
			return ctrl.Result{}, err
		}
		r.Log.Info("Creating a new Ingress",
			"Ingress.Namespace", ingress.Namespace, "Ingress.Name", ingress.Name)
		if err = r.Create(ctx, ingress); err != nil {
			r.Log.Error(err, "Failed to create new Ingress",
				"Ingress.Namespace", ingress.Namespace, "Ingress.Name", ingress.Name)
			return ctrl.Result{}, err
		}

		// Ingress created successfully
		// We will requeue the reconciliation so that we can ensure the state
		// and move forward for the next operations
		return ctrl.Result{RequeueAfter: defaultRequeuePeriod}, nil
	} else if err != nil {
		r.Log.Error(err, "Failed to get Ingress")
		// Let's return the error for the reconciliation be re-triggered again
		return ctrl.Result{}, err
	}

	// Proceed in reconcile loop.
	return ctrl.Result{}, nil
}
