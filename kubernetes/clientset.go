/*
 * Copyright (c) 2019 TFG Co <backend@tfgco.com>
 * Author: TFG Co <backend@tfgco.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package kubernetes

import (
	"context"

	http "github.com/topfreegames/go-extensions-http"
	restWrapper "github.com/topfreegames/go-extensions-k8s-client-go/rest"
	discovery "k8s.io/client-go/discovery"
	kubernetes "k8s.io/client-go/kubernetes"
	admissionregistrationv1alpha1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	batchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	settingsv1alpha1 "k8s.io/client-go/kubernetes/typed/settings/v1alpha1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type NotRESTClientError struct{}

func (e *NotRESTClientError) Error() string {
	return "Not a *rest.RESTClient instance"
}

type NotClientsetError struct{}

func (e *NotClientsetError) Error() string {
	return "Not a *Clientset instance"
}

type Clientset struct {
	*discovery.DiscoveryClient
	admissionregistrationV1alpha1 *admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Client
	admissionregistrationV1beta1  *admissionregistrationv1beta1.AdmissionregistrationV1beta1Client
	appsV1                        *appsv1.AppsV1Client
	appsV1beta1                   *appsv1beta1.AppsV1beta1Client
	appsV1beta2                   *appsv1beta2.AppsV1beta2Client
	authenticationV1              *authenticationv1.AuthenticationV1Client
	authenticationV1beta1         *authenticationv1beta1.AuthenticationV1beta1Client
	authorizationV1               *authorizationv1.AuthorizationV1Client
	authorizationV1beta1          *authorizationv1beta1.AuthorizationV1beta1Client
	autoscalingV1                 *autoscalingv1.AutoscalingV1Client
	autoscalingV2beta1            *autoscalingv2beta1.AutoscalingV2beta1Client
	batchV1                       *batchv1.BatchV1Client
	batchV1beta1                  *batchv1beta1.BatchV1beta1Client
	batchV2alpha1                 *batchv2alpha1.BatchV2alpha1Client
	certificatesV1beta1           *certificatesv1beta1.CertificatesV1beta1Client
	coreV1                        *corev1.CoreV1Client
	eventsV1beta1                 *eventsv1beta1.EventsV1beta1Client
	extensionsV1beta1             *extensionsv1beta1.ExtensionsV1beta1Client
	networkingV1                  *networkingv1.NetworkingV1Client
	policyV1beta1                 *policyv1beta1.PolicyV1beta1Client
	rbacV1                        *rbacv1.RbacV1Client
	rbacV1beta1                   *rbacv1beta1.RbacV1beta1Client
	rbacV1alpha1                  *rbacv1alpha1.RbacV1alpha1Client
	schedulingV1alpha1            *schedulingv1alpha1.SchedulingV1alpha1Client
	schedulingV1beta1             *schedulingv1beta1.SchedulingV1beta1Client
	settingsV1alpha1              *settingsv1alpha1.SettingsV1alpha1Client
	storageV1beta1                *storagev1beta1.StorageV1beta1Client
	storageV1                     *storagev1.StorageV1Client
	storageV1alpha1               *storagev1alpha1.StorageV1alpha1Client
}

// NewForConfig creates a *Clientset that acts as a wrapper over an instance of
// *k8s.io/client-go/kubernetes.Clientset
// OpenTracing instrumentation is set for the underlying *http.Client
// of the rest.Interface sent to all instances created by the wrapped
// *kubernetes.Clientset
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.admissionregistrationV1alpha1, err = admissionregistrationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.admissionregistrationV1beta1, err = admissionregistrationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1, err = appsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta1, err = appsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta2, err = appsv1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1, err = authenticationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1beta1, err = authenticationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1, err = authorizationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1beta1, err = authorizationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV1, err = autoscalingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta1, err = autoscalingv2beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV1, err = batchv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV1beta1, err = batchv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV2alpha1, err = batchv2alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1beta1, err = certificatesv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coreV1, err = corev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.eventsV1beta1, err = eventsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.extensionsV1beta1, err = extensionsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.policyV1beta1, err = policyv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1, err = rbacv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1beta1, err = rbacv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1alpha1, err = rbacv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1beta1, err = schedulingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.settingsV1alpha1, err = settingsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1beta1, err = storagev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1, err = storagev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1alpha1, err = storagev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// WithContext creates a new instance of *Clientset with `ctx` propagated to
// it's components' RESTClient instances
func (c *Clientset) WithContext(ctx context.Context) (*Clientset, error) {
	cs := *c
	if casted, ok := cs.Discovery().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.DiscoveryClient = discovery.NewDiscoveryClient(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AdmissionregistrationV1alpha1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.admissionregistrationV1alpha1 = admissionregistrationv1alpha1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AdmissionregistrationV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.admissionregistrationV1beta1 = admissionregistrationv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AppsV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.appsV1 = appsv1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AppsV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.appsV1beta1 = appsv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AppsV1beta2().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.appsV1beta2 = appsv1beta2.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AuthenticationV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.authenticationV1 = authenticationv1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AuthenticationV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.authenticationV1beta1 = authenticationv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AuthorizationV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.authorizationV1 = authorizationv1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AuthorizationV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.authorizationV1beta1 = authorizationv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AutoscalingV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.autoscalingV1 = autoscalingv1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.AutoscalingV2beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.autoscalingV2beta1 = autoscalingv2beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.BatchV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.batchV1 = batchv1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.BatchV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.batchV1beta1 = batchv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.BatchV2alpha1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.batchV2alpha1 = batchv2alpha1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.CertificatesV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.certificatesV1beta1 = certificatesv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.CoreV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.coreV1 = corev1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.EventsV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.eventsV1beta1 = eventsv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.ExtensionsV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.extensionsV1beta1 = extensionsv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.NetworkingV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.networkingV1 = networkingv1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.PolicyV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.policyV1beta1 = policyv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.RbacV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.rbacV1 = rbacv1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.RbacV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.rbacV1beta1 = rbacv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.RbacV1alpha1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.rbacV1alpha1 = rbacv1alpha1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.SchedulingV1alpha1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.schedulingV1alpha1 = schedulingv1alpha1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.SchedulingV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.schedulingV1beta1 = schedulingv1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.SettingsV1alpha1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.settingsV1alpha1 = settingsv1alpha1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.StorageV1beta1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.storageV1beta1 = storagev1beta1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.StorageV1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.storageV1 = storagev1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	if casted, ok := cs.StorageV1alpha1().RESTClient().(*rest.RESTClient); ok {
		http.Instrument(casted.Client)
		rw := restWrapper.New(casted)
		cs.storageV1alpha1 = storagev1alpha1.New(rw)
	} else {
		return nil, &NotRESTClientError{}
	}
	return &cs, nil
}

// WithContext tries to cast the kubernetes.Interface sent to *Clientset
// and wrap it with `ctx`
func WithContext(c kubernetes.Interface, ctx context.Context) (kubernetes.Interface, error) {
	if casted, ok := c.(*Clientset); ok {
		return casted.WithContext(ctx)
	}
	return nil, &NotClientsetError{}
}

// TryWithContext will return either a *Clientset wrapping `ctx` or the original
// kubernetes.Interface if an error occurs
func TryWithContext(c kubernetes.Interface, ctx context.Context) kubernetes.Interface {
	k, err := WithContext(c, ctx)
	if err != nil {
		return c
	}
	return k
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.DiscoveryClient
}

func (c *Clientset) AdmissionregistrationV1alpha1() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface {
	return c.admissionregistrationV1alpha1
}

func (c *Clientset) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return c.admissionregistrationV1beta1
}

func (c *Clientset) Admissionregistration() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return c.admissionregistrationV1beta1
}

func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return c.appsV1
}

func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return c.appsV1beta1
}

func (c *Clientset) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return c.appsV1beta2
}

func (c *Clientset) Apps() appsv1.AppsV1Interface {
	return c.appsV1
}

func (c *Clientset) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	return c.authenticationV1
}

func (c *Clientset) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	return c.authenticationV1beta1
}

func (c *Clientset) Authentication() authenticationv1.AuthenticationV1Interface {
	return c.authenticationV1
}

func (c *Clientset) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	return c.authorizationV1
}

func (c *Clientset) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	return c.authorizationV1beta1
}

func (c *Clientset) Authorization() authorizationv1.AuthorizationV1Interface {
	return c.authorizationV1
}

func (c *Clientset) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	return c.autoscalingV1
}

func (c *Clientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return c.autoscalingV2beta1
}

func (c *Clientset) Autoscaling() autoscalingv1.AutoscalingV1Interface {
	return c.autoscalingV1
}

func (c *Clientset) BatchV1() batchv1.BatchV1Interface {
	return c.batchV1
}

func (c *Clientset) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return c.batchV1beta1
}

func (c *Clientset) BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface {
	return c.batchV2alpha1
}

func (c *Clientset) Batch() batchv1.BatchV1Interface {
	return c.batchV1
}

func (c *Clientset) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	return c.certificatesV1beta1
}

func (c *Clientset) Certificates() certificatesv1beta1.CertificatesV1beta1Interface {
	return c.certificatesV1beta1
}

func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return c.coreV1
}

func (c *Clientset) Core() corev1.CoreV1Interface {
	return c.coreV1
}

func (c *Clientset) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	return c.eventsV1beta1
}

func (c *Clientset) Events() eventsv1beta1.EventsV1beta1Interface {
	return c.eventsV1beta1
}

func (c *Clientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	return c.extensionsV1beta1
}

func (c *Clientset) Extensions() extensionsv1beta1.ExtensionsV1beta1Interface {
	return c.extensionsV1beta1
}

func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return c.networkingV1
}

func (c *Clientset) Networking() networkingv1.NetworkingV1Interface {
	return c.networkingV1
}

func (c *Clientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	return c.policyV1beta1
}

func (c *Clientset) Policy() policyv1beta1.PolicyV1beta1Interface {
	return c.policyV1beta1
}

func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return c.rbacV1
}

func (c *Clientset) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	return c.rbacV1beta1
}

func (c *Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return c.rbacV1alpha1
}

func (c *Clientset) Rbac() rbacv1.RbacV1Interface {
	return c.rbacV1
}

func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

func (c *Clientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return c.schedulingV1beta1
}

func (c *Clientset) Scheduling() schedulingv1beta1.SchedulingV1beta1Interface {
	return c.schedulingV1beta1
}

func (c *Clientset) SettingsV1alpha1() settingsv1alpha1.SettingsV1alpha1Interface {
	return c.settingsV1alpha1
}

func (c *Clientset) Settings() settingsv1alpha1.SettingsV1alpha1Interface {
	return c.settingsV1alpha1
}

func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return c.storageV1beta1
}

func (c *Clientset) StorageV1() storagev1.StorageV1Interface {
	return c.storageV1
}

func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return c.storageV1alpha1
}

func (c *Clientset) Storage() storagev1.StorageV1Interface {
	return c.storageV1
}
