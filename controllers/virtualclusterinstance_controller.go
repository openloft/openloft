/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	loftv1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	v1 "github.com/loft-sh/api/v3/pkg/apis/storage/v1"
	openloftv1 "github.com/openloft/openloft/api/v1"
	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// VirtualClusterInstanceReconciler reconciles a VirtualClusterInstance object
type VirtualClusterInstanceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclusterinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclusterinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=storage.loft.sh,resources=virtualclusterinstances/finalizers,verbs=update

func (r *VirtualClusterInstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithName("VirtualClusterInstanceReconciler")

	logger.Info("Reconciling...", "instance", req.NamespacedName)

	instance := &loftv1.VirtualClusterInstance{}
	err := r.Client.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// 1. Create a namespace if it does not exist
	if instance.Namespace == "" || instance.Namespace == "default" {
		instance.Namespace = "virtualcluster-" + instance.Name

		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: instance.Namespace,
			},
		}

		err := r.Client.Create(ctx, namespace, &client.CreateOptions{})
		if err != nil && client.IgnoreAlreadyExists(err) != nil {
			return ctrl.Result{}, err
		}
	}

	// 2. Get the template from the template ref
	template := loftv1.VirtualClusterTemplate{}
	err = r.Client.Get(ctx, types.NamespacedName{Name: instance.Spec.TemplateRef.Name}, &template)
	if err != nil {
		return ctrl.Result{}, err
	}

	spec := &openloftv1.VirtualClusterSpec{}
	err = yaml.Unmarshal([]byte(template.Spec.Template.HelmRelease.Values), spec)
	if err != nil {
		return ctrl.Result{}, err
	}

	// 3. Create the virtual cluster
	virtualCluster := &openloftv1.VirtualCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		},
		Spec: *spec,
	}
	err = r.Client.Create(ctx, virtualCluster, &client.CreateOptions{})
	if err != nil {
		return ctrl.Result{}, client.IgnoreAlreadyExists(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VirtualClusterInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.VirtualClusterInstance{}).
		Complete(r)
}
