module github.com/topfreegames/go-extensions-k8s-client-go/kubernetes

go 1.12

replace github.com/topfreegames/go-extensions-k8s-client-go/rest => ../rest

replace github.com/topfreegames/go-extensions-http => ../../go-extensions-http

replace github.com/topfreegames/go-extensions-tracing => ../../go-extensions-tracing

require (
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/topfreegames/go-extensions-http v1.0.0
	github.com/topfreegames/go-extensions-k8s-client-go/rest v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	k8s.io/api v0.0.0-20190126160459-e86510ea3fe7
	k8s.io/apimachinery v0.0.0-20190118094746-1525e4dadd2d
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/klog v0.3.2 // indirect
)
