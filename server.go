package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/debmalya/gomad/log"
	"github.com/debmalya/gomad/router"
)

// SetupEcho to setup echo framework
func SetupEcho() *echo.Echo {
	e := echo.New()
	defer e.Close()
	e.Pre(middleware.RemoveTrailingSlash())

	e.HideBanner = true

	e.Use(middleware.Recover())
	return e
}

// main start application
func main() {

	log.Info("Applicaton starting")
	e := SetupEcho()
	router.InitRoutes(e)

	s := SetupServer(e)
	gracefulShutdown(s, 5*time.Second)
}

// SetupServer - To setup the server
func SetupServer(e *echo.Echo) *http.Server {
	s := &http.Server{Addr: ":1971", Handler: e}
	done := make(chan bool)
	go func() {
		log.Info("Server exit:", s.ListenAndServe())
		done <- true
	}()
	return s
}

// Graceful shutdown of server when receive SIGTERM signal
func gracefulShutdown(hs *http.Server, timeout time.Duration) {

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Infof("\nShutdown with timeout: %s\n", timeout)

	if err := hs.Shutdown(ctx); err != nil {
		log.Errorf("Error: %v\n", err)
	} else {
		//repositories.Close()
		//mq.Disconnect()
		log.Infof("Server stopped")
	}
}
