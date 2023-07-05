package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"we-connect-test/config"
	"we-connect-test/internal/financial"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	DefaultAddr  = "0.0.0.0:8000"
	ReadTimeout  = 30 * time.Second
	WriteTimeout = 30 * time.Second
)

type InternalServerErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type HttpServer struct {
	server   *http.Server
	engine   *gin.Engine
	services Services
}

type Services struct {
	Cfg              *config.Cfg
	FinancialService *financial.Service
}

func (s *HttpServer) ListenAndServe(address string) error {
	s.server.Addr = address
	return s.server.ListenAndServe()
}

func (s *HttpServer) GetEngine() http.Handler {
	return s.engine
}

func NewHttpServer(services Services, logger *zap.Logger) *HttpServer {
	apiRouter := gin.New()
	apiRouter.Use(globalRecover(logger, services.Cfg))
	env := services.Cfg.GetEnv()
	if strings.ToUpper(env) == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}
	server := &http.Server{
		Addr:           DefaultAddr,
		Handler:        apiRouter,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s := &HttpServer{
		server:   server,
		engine:   apiRouter,
		services: services,
	}
	s.registerRoutes()
	return s
}

func (s *HttpServer) registerRoutes() {
	r := s.engine
	v1 := r.Group("/api/v1")
	{
		financialRoutes := v1.Group("/financial")
		{
			financialRoutes.GET("", FinancialIndex(s.services.FinancialService))
			financialRoutes.POST("/create", CreateFinancialData(s.services.FinancialService))
			financialRoutes.POST("/update", UpdateFinancialData(s.services.FinancialService))
			financialRoutes.POST("/delete", DeleteFinancialData(s.services.FinancialService))
		}
	}
}

func globalRecover(logger *zap.Logger, cfg *config.Cfg) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {
			message := http.StatusText(http.StatusInternalServerError)
			if rec := recover(); rec != nil {
				if cfg.GetEnv() != config.EnvProd {
					fmt.Println("rec  =>", rec)
				}
				err := errors.New("error 500")
				logger.Error(fmt.Sprintf("error  500 in global recover %v", rec),
					zap.Error(err),
					zap.String("service", "httpServer"),
					zap.String("method", "globalRecover"),
				)
				response := InternalServerErrorResponse{
					Status:  false,
					Message: message,
				}
				c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			}
		}(c)
		c.Next()
	}
}
