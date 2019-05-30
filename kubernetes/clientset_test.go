package kubernetes_test

import (
	"context"
	"testing"

	kubernetesExtensions "github.com/topfreegames/go-extensions-k8s-client-go/kubernetes"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

var c = &rest.Config{
	APIPath: "/apis",
	ContentConfig: rest.ContentConfig{
		NegotiatedSerializer: scheme.Codecs,
		GroupVersion:         &appsv1.SchemeGroupVersion,
	},
}

func TestNewForConfig(t *testing.T) {
	_, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
}

func TestWithContext(t *testing.T) {
	c, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	cc := c.WithContext(context.Background())
	if cc == c {
		t.Fatal("Expected WithContext to return a new instance of *kubernetesExtensions.Clientset")
	}
	if cc.Interface == c.Interface {
		t.Fatal("Expected cc.Interface to be != c.Interface")
	}
}

type clientsetMock struct {
	calls  map[string]int
	client rest.Interface
	*kubernetes.Clientset
}

func newClientsetMock(c *rest.Config) (*clientsetMock, error) {
	client, err := rest.RESTClientFor(c)
	if err != nil {
		return nil, err
	}
	return &clientsetMock{
		calls:     map[string]int{},
		client:    client,
		Clientset: kubernetes.New(client),
	}, nil
}

func (c *clientsetMock) CoreV1() corev1.CoreV1Interface {
	c.calls["CoreV1"]++
	return corev1.New(c.client)
}

func TestProxiedCalls(t *testing.T) {
	cs, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	m, err := newClientsetMock(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	cs.Interface = m
	if m.calls["CoreV1"] != 0 {
		t.Fatalf("Expected calls[CoreV1] to be 0. Got %d", m.calls["CoreV1"])
	}
	cs.CoreV1()
	if m.calls["CoreV1"] != 1 {
		t.Fatalf("Expected calls[CoreV1] to be 0. Got %d", m.calls["CoreV1"])
	}
}

func TestTryWithContextWithValidCast(t *testing.T) {
	cs, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	k, ok := kubernetesExtensions.TryWithContext(cs, context.Background())
	if k == cs {
		t.Fatal("Expected TryWithContext to return a new instance of *kubernetesExtensions.Clientset")
	}
	if ok != true {
		t.Fatal("Expected TryWithContext to return true")
	}
}

func TestTryWithContextWithInvalidCast(t *testing.T) {
	m, err := newClientsetMock(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	k, ok := kubernetesExtensions.TryWithContext(m, context.Background())
	if k != m {
		t.Fatal("Expected TryWithContext to return the same clientsetMock instance")
	}
	if ok != false {
		t.Fatal("Expected TryWithContext to return false")
	}
}
