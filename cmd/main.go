package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/salesforceanton/products-data-manager/internal/config"
	"github.com/salesforceanton/products-data-manager/internal/handler"
	"github.com/salesforceanton/products-data-manager/internal/logger"
	"github.com/salesforceanton/products-data-manager/internal/repository"
	"github.com/salesforceanton/products-data-manager/internal/server"
	"github.com/salesforceanton/products-data-manager/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	// Intialize App configuration
	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogServerIssue(err)
		return
	}

	// Connect to DB
	db, err := repository.NewMongoDB(cfg)
	if err != nil {
		logger.LogServerIssue(err)
		return
	}

	// Set dependenties
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)
	server := server.New(handler)

	logrus.Info(fmt.Sprintf("SERVER STARTED: %s", time.Now().Local().String()))
	go func() {
		if err := server.ListenAndServe(cfg.Port); err != nil {
			logger.LogServerIssue(err)
			return
		}

	}()

	// Gracefull shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	if err := db.Client().Disconnect(context.Background()); err != nil {
		logger.LogServerIssue(err)
		return
	}

	logrus.Info("Server shutdown successfully")
}
