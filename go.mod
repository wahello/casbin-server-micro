module github.com/paysuper/casbin-server

go 1.12

require (
	github.com/InVisionApp/go-health v2.1.1-0.20181204182500-21e799eae2a9+incompatible
	github.com/InVisionApp/go-logger v1.0.1 // indirect
	github.com/casbin/casbin/v2 v2.0.2
	github.com/casbin/xorm-adapter v0.0.0-20190806085643-0629743c2857
	github.com/go-xorm/xorm v0.7.6 // indirect
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/hashicorp/consul/api v1.2.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.1.6
	github.com/lib/pq v1.2.0
	github.com/marten-seemann/qtls v0.3.3 // indirect
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.2.1-0.20190803141733-bc5828af5d4f
	github.com/miekg/dns v1.1.17 // indirect
	github.com/paysuper/echo-casbin-middleware v0.0.0-20190810212121-02efa137c0bb
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4 // indirect
	github.com/prometheus/procfs v0.0.4 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190911031432-227b76d455e7 // indirect
	golang.org/x/net v0.0.0-20190909003024-a7b16738d86b // indirect
	golang.org/x/sys v0.0.0-20190911201528-7ad0cfa0b7b5 // indirect
	google.golang.org/genproto v0.0.0-20190911173649-1774047e7e51 // indirect
	google.golang.org/grpc v1.23.1 // indirect
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.0.0-00010101000000-000000000000 // indirect
)

replace (
	github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.2
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.1.0
	github.com/unistack-org/casbin-micro => github.com/paysuper/casbin-server v0.0.0-20190809212717-30678425fdda
	gopkg.in/DATA-DOG/go-sqlmock.v1 => github.com/DATA-DOG/go-sqlmock v1.3.3
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.21.0
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)
