package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/micro/go-micro/client"
	casbinmw "github.com/paysuper/echo-casbin-middleware"
)

var (
	httpAddr string
)

func TestMain(m *testing.M) {
	var err error

	// Echo instance
	e := echo.New()

	c := client.NewClient()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(casbinmw.Middleware(c))

	// Route => handler
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!\n")
	})

	go func(err error) {
		// Start server
		err = e.Start(":1323")
	}(err)

	// close server
	defer e.Close()

	httpAddr = e.Server.Addr

	ecode := m.Run()

	os.Exit(ecode)
}
