module github.com/unistack-org/casbin-micro

go 1.12

require (
	github.com/ProtocolONE/go-micro-plugins v0.3.0
	github.com/casbin/casbin/v2 v2.0.1
	github.com/casbin/gorm-adapter v1.0.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/lucas-clemente/quic-go v0.11.2 // indirect
	github.com/micro/go-micro v1.7.1-0.20190722165734-49dcc3d1bdfe
)

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.2
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.1.0
)
