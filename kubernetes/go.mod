module github.com/topfreegames/go-extensions-k8s-client-go/kubernetes

go 1.12

replace github.com/topfreegames/go-extensions-k8s-client-go/rest => ../rest

replace github.com/topfreegames/go-extensions-http => ../../go-extensions-http

replace github.com/topfreegames/go-extensions-tracing => ../../go-extensions-tracing

require (
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/topfreegames/go-extensions-http v0.0.0-00010101000000-000000000000
	github.com/topfreegames/go-extensions-k8s-client-go/rest v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092 // indirect
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apimachinery v0.0.0-20190313205120-d7deff9243b1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v0.3.2 // indirect
)
