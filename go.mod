module github.com/paysuper/casbin-server

go 1.12

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/casbin/casbin/v2 v2.0.1
	github.com/casbin/xorm-adapter v0.0.0-20190806085643-0629743c2857
	github.com/golang/protobuf v1.3.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.1.6
	github.com/lib/pq v1.2.0
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/micro/go-micro v1.7.1-0.20190724203029-7ca8f8f0ab98
	github.com/micro/go-plugins v1.1.2-0.20190723131713-899a15dc9867
	github.com/paysuper/echo-casbin-middleware v0.0.0-20190808195909-2df368871e0e
	github.com/prometheus/client_golang v1.0.0
	github.com/stretchr/testify v1.3.0
	go.uber.org/zap v1.10.0
)

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.2
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.1.0
)
