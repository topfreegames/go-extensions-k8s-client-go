package kubernetes_test

import (
	"context"
	"reflect"
	"testing"
	"unsafe"

	kubernetesExtensions "github.com/topfreegames/go-extensions-k8s-client-go/kubernetes"
	restWrapper "github.com/topfreegames/go-extensions-k8s-client-go/rest"
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

type wrongRest struct {
	*rest.RESTClient
}

func TestUnexpectedRestInterfaceImpl(t *testing.T) {
	k, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	cv1 := k.CoreV1()
	rf := reflect.ValueOf(cv1).Elem().FieldByName("restClient")
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	rf.Set(reflect.ValueOf(&wrongRest{}))
	err = kubernetesExtensions.Instrument(k, nil)
	if _, ok := err.(*kubernetesExtensions.UnexpectedRestInterfaceImplError); !ok {
		t.Fatal("Expected err to be UnexpectedRestInterfaceImplError")
	}
}

func TestInstrumentShouldReplaceByRestWrapper(t *testing.T) {
	k, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	client := k.CoreV1().RESTClient()
	if _, ok := client.(*restWrapper.Client); !ok {
		t.Fatal("Expected client to be *restWrapper.Client")
	}
}

func TestInstrumentWithContextShouldReturnWrapperWithCtx(t *testing.T) {
	k, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	ctx := context.Background()
	kk, err := k.WithContext(ctx)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	clientWithCtx := kk.CoreV1().RESTClient()
	req := clientWithCtx.Get()
	rf := reflect.ValueOf(req).Elem().FieldByName("ctx")
	reqCtx := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).
		Elem().Interface().(context.Context)
	if ctx != reqCtx {
		t.Fatal("Expected ctx and reqCtx to be the same")
	}
}

func TestClientWithContextShouldReturnNewClient(t *testing.T) {
	k, err := kubernetesExtensions.NewForConfig(c)
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	cc, err := k.WithContext(context.Background())
	if err != nil {
		t.Fatalf("Expected err not to have occurred. Err: %s", err.Error())
	}
	if cc == k {
		t.Fatal("Expected WithContext to return a new instance of *kubernetesExtensions.Clientset")
	}
	if cc.CoreV1() == k.CoreV1() {
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
