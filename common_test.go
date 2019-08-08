package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/micro/go-micro/client"
	casbinmw "github.com/paysuper/echo-casbin-middleware"
)

var (
	httpAddr      string
	getHandler    = func(ctx echo.Context) error { return ctx.String(http.StatusOK, "GET\n") }
	postHandler   = func(ctx echo.Context) error { return ctx.String(http.StatusCreated, "POST\n") }
	deleteHandler = func(ctx echo.Context) error { return ctx.NoContent(http.StatusNoContent) }
	patchHandler  = func(ctx echo.Context) error { return ctx.String(http.StatusOK, "PATCH\n") }
	putHandler    = func(ctx echo.Context) error { return ctx.String(http.StatusOK, "PUT\n") }
)

type (
	routerFn func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route
)

func TestMain(m *testing.M) {
	var err error

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	c := client.NewClient()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(casbinmw.Middleware(c))

	mapfn := make(map[string]routerFn)
	mapfn["GET"] = e.GET
	mapfn["POST"] = e.POST
	mapfn["PATCH"] = e.PATCH
	mapfn["DELETE"] = e.DELETE
	mapfn["PUT"] = e.PUT

	if err := setupHandlers(e, mapfn); err != nil {
		log.Fatal(err)
	}

	// Route => handler
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!\n")
	})

	port, err := getFreeTestPort()
	if err != nil {
		log.Fatal(err)
	}
	httpAddr := "localhost:" + port

	go func(err error) {
		// Start server
		err = e.Start(httpAddr)
	}(err)

	// close server
	defer e.Close()

	log.Printf("echo running on %s", httpAddr)
	ecode := m.Run()

	os.Exit(ecode)
}

func setupHandlers(e *echo.Echo, mapfn map[string]routerFn) error {
	f, err := os.OpenFile("policy.conf", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		fields := strings.Split(line, ",")
		if len(fields) < 2 {
			continue
		}

		for i, f := range fields {
			fields[i] = strings.TrimSpace(f)
		}

		method := fields[len(fields)-1]
		fn, ok := mapfn[method]
		if !ok {
			continue
		}

		switch method {
		case "GET":
			fn(fields[2], getHandler)
		case "POST":
			fn(fields[2], postHandler)
		case "PUT":
			fn(fields[2], putHandler)
		case "PATCH":
			fn(fields[2], patchHandler)
		case "DELETE":
			fn(fields[2], deleteHandler)
		}
	}

	return nil
}

func getFreeTestPort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = l.Close()
	}()
	return strconv.Itoa(l.Addr().(*net.TCPAddr).Port), nil
}
