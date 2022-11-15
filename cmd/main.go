package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pashapdev/applicationDesignTest/internal/application"
	"github.com/pashapdev/applicationDesignTest/internal/config"
	"github.com/pashapdev/applicationDesignTest/internal/logger"
)

func main() {
	logger.Info("reading config...")
	conf, err := config.LoadFromEnv()
	if err != nil {
		logger.Error("failed to read config")
		os.Exit(1)
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	app, err := application.New(conf)
	if err != nil {
		logger.Error("failed to read config")
		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		app.GracefulStop(serverCtx, sig, serverStopCtx)
	}()

	err = app.Run()
	if err != nil {
		logger.Error("failed to read config")
		os.Exit(1)
	}

	<-serverCtx.Done()
}
