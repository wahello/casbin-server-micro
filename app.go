package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	health "github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/handlers"
	"github.com/casbin/casbin/v2/model"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/client/selector/static"
	metrics "github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	casbinpb "github.com/paysuper/casbin-server/casbinpb"
	cserver "github.com/paysuper/casbin-server/server"
	"go.uber.org/zap"

	// database support
	xormadapter "github.com/casbin/xorm-adapter"
	// postgresql support
	_ "github.com/lib/pq"
	// remove
	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	cfg        *Config
	msvc       micro.Service
	httpServer *http.Server
	router     *http.ServeMux
	logger     *zap.Logger
	csvc       *cserver.Server
}

type appHealthCheck struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	app.initLogger()

	cfg, err := NewConfig()

	if err != nil {
		app.logger.Fatal("Config load failed", zap.Error(err))
	}

	app.cfg = cfg

	options := []micro.Option{
		micro.Name("casbin"),
		micro.Version("0.0.1"),
		micro.WrapHandler(metrics.NewHandlerWrapper()),
		micro.AfterStop(func() error {
			app.logger.Info("Micro service stopped")
			app.Stop()
			return nil
		}),
	}

	if os.Getenv("MICRO_SELECTOR") == "static" {
		log.Println("Use micro selector `static`")
		options = append(options, micro.Selector(static.NewSelector()))
	}

	app.logger.Info("Initialize micro service")

	app.msvc = micro.NewService(options...)
	app.msvc.Init()

	m, err := model.NewModelFromFile("model.conf")
	if err != nil {
		app.logger.Fatal("Failed to init casbin model", zap.Error(err))
	}

	xormAdapter, err := xormadapter.NewAdapter(cfg.CasbinAdapter, cfg.CasbinDSN)
	if err != nil {
		app.logger.Fatal("Failed to init casbin xorm adapter", zap.Error(err))
	}

	app.csvc, err = cserver.NewServer(m, xormAdapter)
	if err != nil {
		app.logger.Fatal("Create service instance failed", zap.Error(err))
	}

	err = app.csvc.LoadPolicyFile("policy.conf")
	if err != nil {
		app.logger.Fatal("Failed to load policy file to adapter", zap.Error(err))
	}

	err = casbinpb.RegisterCasbinHandler(app.msvc.Server(), app.csvc)
	if err != nil {
		app.logger.Fatal("Service init failed", zap.Error(err))
	}

	app.router = http.NewServeMux()
	app.initHealth()
	app.initMetrics()
}

func (app *Application) initLogger() {
	var err error

	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("Application logger initialization failed with error: %s\n", err)
	}
	app.logger = logger.Named("CASBIN_SERVER")
	zap.ReplaceGlobals(app.logger)
}

func (app *Application) initHealth() {
	h := health.New()
	err := h.AddChecks([]*health.Config{
		{
			Name:     "health-check",
			Checker:  &appHealthCheck{},
			Interval: time.Duration(1) * time.Second,
			Fatal:    true,
		},
	})

	if err != nil {
		app.logger.Fatal("Health check register failed", zap.Error(err))
	}

	if err = h.Start(); err != nil {
		app.logger.Fatal("Health check start failed", zap.Error(err))
	}

	app.logger.Info("Health check listener started", zap.String("port", app.cfg.MetricsPort))

	app.router.HandleFunc("/health", handlers.NewJSONHandlerFunc(h, nil))
}

func (app *Application) initMetrics() {
	app.router.Handle("/metrics", promhttp.Handler())
}

func (app *Application) Run() {
	app.httpServer = &http.Server{
		Addr:    ":" + app.cfg.MetricsPort,
		Handler: app.router,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Fatal("Http server starting failed", zap.Error(err))
		}
	}()

	if err := app.msvc.Run(); err != nil {
		app.logger.Fatal("Micro service starting failed", zap.Error(err))
	}
}

func (c *appHealthCheck) Status() (interface{}, error) {
	return "ok", nil
}

func (app *Application) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.httpServer.Shutdown(ctx); err != nil {
		app.logger.Error("Http server shutdown failed", zap.Error(err))
	}
	app.logger.Info("Http server stopped")

	if err := app.logger.Sync(); err != nil {
		app.logger.Error("Logger sync failed", zap.Error(err))
	} else {
		app.logger.Info("Logger synced")
	}
}
