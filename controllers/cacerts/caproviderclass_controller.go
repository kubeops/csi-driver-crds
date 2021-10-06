/*
Copyright 2021.

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

package cacerts

import (
	"context"

	"github.com/jetstack/cert-manager/pkg/apis/certmanager"
	certmanagerv1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	cacertsv1alpha1 "kubeops.dev/csi-driver-ca-certificates-crds/apis/cacerts/v1alpha1"
)

// CAProviderClassReconciler reconciles a CAProviderClass object
type CAProviderClassReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cacerts.csi.appscode.com,resources=caproviderclasses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cacerts.csi.appscode.com,resources=caproviderclasses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cacerts.csi.appscode.com,resources=caproviderclasses/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CAProviderClass object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *CAProviderClassReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CAProviderClassReconciler) SetupWithManager(mgr ctrl.Manager) error {
	secretHandler := handler.EnqueueRequestsFromMapFunc(func(a client.Object) []reconcile.Request {
		providers := &cacertsv1alpha1.CAProviderClassList{}
		if err := r.List(context.Background(), providers); err != nil {
			return nil
		}
		var req []reconcile.Request

		var ns, name string
		for _, p := range providers.Items {
			for _, ref := range p.Spec.SecretRefs {
				ns = ref.Namespace
				if ns == "" {
					ns = p.Namespace
				}
				name = ref.Name

				if ns == a.GetNamespace() && name == a.GetName() {
					req = append(req, reconcile.Request{NamespacedName: client.ObjectKeyFromObject(&p)})
					break
				}
			}
		}
		return req
	})

	certHandler := handler.EnqueueRequestsFromMapFunc(func(a client.Object) []reconcile.Request {
		providers := &cacertsv1alpha1.CAProviderClassList{}
		if err := r.List(context.Background(), providers); err != nil {
			return nil
		}
		var req []reconcile.Request

		var ns, name string
		for _, p := range providers.Items {
			for _, ref := range p.Spec.CertificateRefs {
				ns = ref.Namespace
				if ns == "" {
					ns = p.Namespace
				}
				name = ref.Name

				if ns == a.GetNamespace() && name == a.GetName() {
					req = append(req, reconcile.Request{NamespacedName: client.ObjectKeyFromObject(&p)})
					break
				}
			}
		}
		return req
	})

	issuerHandler := handler.EnqueueRequestsFromMapFunc(func(a client.Object) []reconcile.Request {
		providers := &cacertsv1alpha1.CAProviderClassList{}
		if err := r.List(context.Background(), providers); err != nil {
			return nil
		}
		var req []reconcile.Request

		var ns, name string
		for _, p := range providers.Items {
			for _, ref := range p.Spec.IssuerRefs {
				if ref.APIGroup == nil || *ref.APIGroup != certmanager.GroupName || ref.Kind != "Issuer" {
					continue
				}
				ns = ref.Namespace
				if ns == "" {
					ns = p.Namespace
				}
				name = ref.Name

				if ns == a.GetNamespace() && name == a.GetName() {
					req = append(req, reconcile.Request{NamespacedName: client.ObjectKeyFromObject(&p)})
					break
				}
			}
		}
		return req
	})

	clusterIssuerHandler := handler.EnqueueRequestsFromMapFunc(func(a client.Object) []reconcile.Request {
		providers := &cacertsv1alpha1.CAProviderClassList{}
		if err := r.List(context.Background(), providers); err != nil {
			return nil
		}
		var req []reconcile.Request

		for _, p := range providers.Items {
			for _, ref := range p.Spec.IssuerRefs {
				if ref.APIGroup == nil || *ref.APIGroup != certmanager.GroupName || ref.Kind != "ClusterIssuer" {
					continue
				}

				if ref.Name == a.GetName() {
					req = append(req, reconcile.Request{NamespacedName: client.ObjectKeyFromObject(&p)})
					break
				}
			}
		}
		return req
	})

	return ctrl.NewControllerManagedBy(mgr).
		For(&cacertsv1alpha1.CAProviderClass{}).
		Watches(&source.Kind{Type: &core.Secret{}}, secretHandler).
		Watches(&source.Kind{Type: &certmanagerv1.Certificate{}}, certHandler).
		Watches(&source.Kind{Type: &certmanagerv1.Issuer{}}, issuerHandler).
		Watches(&source.Kind{Type: &certmanagerv1.ClusterIssuer{}}, clusterIssuerHandler).
		Complete(r)
}
