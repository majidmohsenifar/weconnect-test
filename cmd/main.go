package main

import (
	"context"
	"fmt"
	"log"
	"we-connect-test/internal/di"
	"we-connect-test/internal/handler/api"
	"we-connect-test/internal/queue"

	"go.uber.org/zap"
)

const (
	defaultPort = "8000"
)

func main() {
	container := di.NewContainer()
	logger, err := container.GetLogger()
	if err != nil {
		log.Fatal("cannot initialize logger")
	}
	_, err = container.GetMongoDBClient()
	if err != nil {
		log.Fatal("cannot initialize mongo")
	}
	financialService := container.GetFinancialService()
	//here we run queue
	ctx := context.Background()
	go func() {
		queueManager := queue.NewManager(financialService, logger)
		filePath := "./data.csv"
		workerCount := 5
		err = queueManager.Run(ctx, filePath, workerCount)
		if err != nil {
			logger.Fatal("err in queueManager", zap.Error(err))
		}
	}()

	//here we run httpServer
	httpServer := api.NewHttpServer(api.Services{
		Cfg:              container.GetCfg(),
		FinancialService: financialService,
	}, logger)
	err = httpServer.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", defaultPort))
	if err != nil {
		log.Fatal("can not run http server because of err" + err.Error())
	}
}
