package kubernetes

import (
	"context"
	"reflect"
	"unsafe"

	http "github.com/topfreegames/go-extensions-http"
	restWrapper "github.com/topfreegames/go-extensions-k8s-client-go/rest"
	"k8s.io/client-go/rest"
)

// for each field in *kubernetes.Clientset, instrument it's underlying `restClient`
func Instrument(c *Clientset, ctx context.Context) error {
	rs := reflect.ValueOf(c).Elem()
	rf := rs.FieldByName("Clientset")
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	for _, s := range fields {
		if err := instrumentStruct(rf, ctx, s); err != nil {
			return err
		}
	}
	return nil
}

// access `s`.`restClient` and instrument it. It's type is expected to be
// *rest.RESTClient. After instrumentation, it's reset wrapped in *restWrapper.Client
func instrumentStruct(c reflect.Value, ctx context.Context, s string) error {
	// TODO: test when field doesnt exist, handle error instead of panic
	rf := reflect.Indirect(c).FieldByName(s)
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	rf = reflect.Indirect(rf).FieldByName("restClient")
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	v, ok := rf.Interface().(*rest.RESTClient)
	if !ok {
		return &UnexpectedRestInterfaceImplError{}
	}
	http.Instrument(v.Client)
	wrapped := restWrapper.New(v).WithContext(ctx)
	rf.Set(reflect.ValueOf(wrapped))
	return nil
}
