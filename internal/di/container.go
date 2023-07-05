package di

import (
	"we-connect-test/config"
	"we-connect-test/internal/client"
	"we-connect-test/internal/financial"
	"we-connect-test/internal/handler/api"
	"we-connect-test/internal/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Container struct {
	httpServer       *api.HttpServer
	logger           *zap.Logger
	cfg              *config.Cfg
	financialService *financial.Service
	financialRepo    *financial.Repository
	mongoDBClient    *mongo.Client
}

func (c *Container) GetLogger() (*zap.Logger, error) {
	if c.logger == nil {
		cfg := c.GetCfg()
		logger, err := logger.NewLogger(cfg)
		if err != nil {
			return nil, err
		}
		c.logger = logger
	}
	return c.logger, nil
}

func (c *Container) GetMongoDBClient() (*mongo.Client, error) {
	if c.mongoDBClient == nil {
		cfg := c.GetCfg()
		client, err := client.NewMongoDBClient(cfg)
		if err != nil {
			return nil, err
		}
		c.mongoDBClient = client
	}
	return c.mongoDBClient, nil
}

func (c *Container) GetCfg() *config.Cfg {
	if c.cfg == nil {
		viper := config.SetCfg()
		cfg := config.NewConfigs(viper)
		c.cfg = cfg
	}
	return c.cfg
}

func (c *Container) GetFinancialRepository() *financial.Repository {
	if c.financialService == nil {
		cfg := c.GetCfg()
		mongoDBClient, _ := c.GetMongoDBClient()
		c.financialRepo = financial.NewRepository(cfg, mongoDBClient)
	}
	return c.financialRepo
}

func (c *Container) GetFinancialService() *financial.Service {
	if c.financialService == nil {
		repo := c.GetFinancialRepository()
		logger, _ := c.GetLogger()
		c.financialService = financial.NewService(repo, logger)
	}
	return c.financialService
}

func NewContainer() *Container {
	return &Container{}
}
