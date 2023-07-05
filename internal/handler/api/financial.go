package api

import (
	"net/http"
	"we-connect-test/internal/financial"

	"github.com/gin-gonic/gin"
)

func FinancialIndex(s *financial.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := financial.GetFinancialDataListParams{}
		err := c.ShouldBindQuery(&p)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		resp, statusCode := s.GetFinancialDataList(c, p)
		c.JSON(statusCode, resp)
	}
}

func CreateFinancialData(s *financial.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := financial.CreateFinancialDataParams{}
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		resp, statusCode := s.CreateFinancialDataByUser(c, p)
		c.JSON(statusCode, resp)
	}
}

func UpdateFinancialData(s *financial.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := financial.UpdateFinancialDataParams{}
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		resp, statusCode := s.UpdateFinancialData(c, p)
		c.JSON(statusCode, resp)
	}
}

func DeleteFinancialData(s *financial.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		p := financial.DeleteFinancialDataParams{}
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		resp, statusCode := s.DeleteFinancialData(c, p)
		c.JSON(statusCode, resp)
	}
}
