module github.com/unistack-org/casbin-micro

go 1.12

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/InVisionApp/go-logger v1.0.1 // indirect
	github.com/casbin/casbin/v2 v2.0.1
	github.com/casbin/gorm-adapter/v2 v2.0.0
	github.com/denisenkom/go-mssqldb v0.0.0-20190707035753-2be1aa521ff4 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.1.1
	github.com/lucas-clemente/quic-go v0.11.2 // indirect
	github.com/micro/go-micro v1.7.1-0.20190724203029-7ca8f8f0ab98
	github.com/micro/go-plugins v1.1.2-0.20190723131713-899a15dc9867
	github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d // indirect
	github.com/prometheus/client_golang v1.0.0
	github.com/stretchr/testify v1.3.0
	go.uber.org/zap v1.10.0
)

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.2
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.1.0
)
