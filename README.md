# go-extensions-k8s-client-go

This go library is a wrapper over [k8s.io/client-go](https://godoc.org/k8s.io/client-go) that adds OpenTracing support through a new method `WithContext` added to our implementations of their kubernetes.Interface and rest.Interface.

All the other methods are proxies to the wrapped instances of *kubernetes.Clientset and *rest.RESTClient methods.
