package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/adhityapp/go-starterkit/bootstrap"
	"github.com/adhityapp/go-starterkit/internal/gate.go"
)

func Serve(cfg *bootstrap.Container) {
	var (
		server = echo.New()
	)

	server.Debug = false
	server.HideBanner = true
	server.HidePort = true
	server.Pre(
		middleware.RemoveTrailingSlash(),
	)

	server.Use(
		// otelecho.Middleware(os.Getenv("SERVICE_NAME")),
		middleware.Recover(),
		middleware.RequestID(),
		middleware.CORS(),
		gate.RestLogger(),
	)

	var (
	// Init Repo

	// Init Service

	// Init Usecase

	// Init Controller
	)

	// Call Rest Register
	// health.RestRegister(cfg, server)
	//handler.RestRegister(server, namafungsihandler)

	go func() {
		port := viper.GetString("server.port")
		if err := server.Start((":" + port)); err != nil {
			if err == http.ErrServerClosed {
				logrus.Info("server stopped")
			} else {
				logrus.Fatal("failed to start server " + err.Error())
			}
		}
		logrus.Infof("server starting at %s", port)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// do other stuff
	logrus.Info("shutting down the server")
	cfg.Close()

	if err := server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		logrus.Fatal("failed to gracefully shut down the server " + err.Error())
	}
}
