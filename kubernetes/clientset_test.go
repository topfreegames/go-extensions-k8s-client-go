package kubernetes_test

import (
	"context"
	"testing"

	kubernetesExtensions "github.com/topfreegames/go-extensions-k8s-client-go/kubernetes"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
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
	if c.Instrumented() == true {
		t.Fatal("Expected c *Clientset not to be instrumented")
	}
	cc, err := c.WithContext(context.Background())
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	if c.Instrumented() == true {
		t.Fatal("Expected c *Clientset not to be instrumented")
	}
	if cc.Instrumented() == false {
		t.Fatal("Expected cc *Clientset to be instrumented")
	}
	if cc == c {
		t.Fatal("Expected WithContext to return a new instance of *kubernetesExtensions.Clientset")
	}
	if cc.CoreV1() == c.CoreV1() {
		t.Fatal("Expected cc.CoreV1() to be != c.CoreV1()")
	}
}

type clientsetMock struct {
	*kubernetes.Clientset
}

func newClientsetMock(c *rest.Config) (*clientsetMock, error) {
	client, err := rest.RESTClientFor(c)
	if err != nil {
		return nil, err
	}
	return &clientsetMock{Clientset: kubernetes.New(client)}, nil
}

func TestTryWithContextWithValidCast(t *testing.T) {
	cs, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	k := kubernetesExtensions.TryWithContext(cs, context.Background())
	if k == cs {
		t.Fatal("Expected TryWithContext to return a new instance of *kubernetesExtensions.Clientset")
	}
}

func TestTryWithContextWithInvalidCast(t *testing.T) {
	m, err := newClientsetMock(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	k := kubernetesExtensions.TryWithContext(m, context.Background())
	if k != m {
		t.Fatal("Expected TryWithContext to return the same clientsetMock instance")
	}
}
