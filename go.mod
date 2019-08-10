module github.com/paysuper/casbin-server

go 1.12

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/InVisionApp/go-logger v1.0.1 // indirect
	github.com/casbin/casbin/v2 v2.0.1
	github.com/casbin/xorm-adapter v0.0.0-20190806085643-0629743c2857
	github.com/golang/protobuf v1.3.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.1.6
	github.com/lib/pq v1.2.0
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/micro/go-micro v1.8.3-0.20190810174654-8986b3135f72
	github.com/micro/go-plugins v1.2.1-0.20190803141733-bc5828af5d4f
	github.com/paysuper/echo-casbin-middleware v0.0.0-20190810212121-02efa137c0bb
	github.com/prometheus/client_golang v1.1.0
	github.com/stretchr/testify v1.3.0
	go.uber.org/zap v1.10.0
)

replace (
	github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.2
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.1.0
	github.com/unistack-org/casbin-micro => github.com/paysuper/casbin-server v0.0.0-20190809212717-30678425fdda
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.21.0
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)
