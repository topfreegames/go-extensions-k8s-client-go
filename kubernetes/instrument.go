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
func instrument(c *Clientset, ctx context.Context) error {
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

// access `s`.`restClient` and instrument it if it's the first time, meaning
// it's type is *rest.RESTClient, or change it's `ctx` if it's already a
// *restWrapper.Client
func instrumentStruct(c reflect.Value, ctx context.Context, s string) error {
	// TODO: test when field doesnt exist, handle error instead of panic
	rf := reflect.Indirect(c).FieldByName(s)
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	rf = reflect.Indirect(rf).FieldByName("restClient")
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	clientI := rf.Interface()
	switch v := clientI.(type) {
	case *rest.RESTClient:
		http.Instrument(v.Client)
	case *restWrapper.Client:
		clientI = v.WithContext(ctx)
	default:
		return &UnexpectedRestInterfaceImplError{}
	}
	rf.Set(reflect.ValueOf(clientI))
	return nil
}
