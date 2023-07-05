package financial

import (
	"context"
	"errors"
	"net/http"
	"we-connect-test/internal/response"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service struct {
	repo   *Repository
	logger *zap.Logger
}

type GetFinancialDataListParams struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type SingleFinancialDataResult struct {
	ID              string `json:"id"`
	SeriesReference string `json:"seriesReference"`
	Period          string `json:"period"`
	DataValue       string `json:"dataValue"`
	Suppressed      string `json:"suppressed"`
	Status          string `json:"status"`
	Units           string `json:"units"`
	Magnitude       string `json:"magnitude"`
	Subject         string `json:"subject"`
	Group           string `json:"group"`
	SeriesTitle1    string `json:"seriesTitle1"`
	SeriesTitle2    string `json:"seriesTitle2"`
	SeriesTitle3    string `json:"seriesTitle3"`
	SeriesTitle4    string `json:"seriesTitle4"`
	SeriesTitle5    string `json:"seriesTitle5"`
}

type CreateFinancialDataParams struct {
	SeriesReference string `json:"seriesReference"`
	Period          string `json:"period"`
	DataValue       string `json:"dataValue"`
	Suppressed      string `json:"suppressed"`
	Status          string `json:"status"`
	Units           string `json:"units"`
	Magnitude       string `json:"magnitude"`
	Subject         string `json:"subject"`
	Group           string `json:"group"`
	SeriesTitle1    string `json:"seriesTitle1"`
	SeriesTitle2    string `json:"seriesTitle2"`
	SeriesTitle3    string `json:"seriesTitle3"`
	SeriesTitle4    string `json:"seriesTitle4"`
	SeriesTitle5    string `json:"seriesTitle5"`
}

type UpdateFinancialDataParams struct {
	ID              string  `json:"id"`
	SeriesReference *string `json:"seriesReference"`
	Period          *string `json:"period"`
	DataValue       *string `json:"dataValue"`
	Suppressed      *string `json:"suppressed"`
	Status          *string `json:"status"`
	Units           *string `json:"units"`
	Magnitude       *string `json:"magnitude"`
	Subject         *string `json:"subject"`
	Group           *string `json:"group"`
	SeriesTitle1    *string `json:"seriesTitle1"`
	SeriesTitle2    *string `json:"seriesTitle2"`
	SeriesTitle3    *string `json:"seriesTitle3"`
	SeriesTitle4    *string `json:"seriesTitle4"`
	SeriesTitle5    *string `json:"seriesTitle5"`
}

type DeleteFinancialDataParams struct {
	ID string `json:"id"`
}

func (s *Service) GetFinancialDataList(
	ctx context.Context,
	params GetFinancialDataListParams,
) (apiResponse response.ApiResponse, statusCode int) {
	if params.Page < 0 {
		params.Page = 0
	}
	if params.PageSize < 2 {
		params.PageSize = 2
	}
	if params.PageSize > 100 {
		params.PageSize = 100
	}
	models, err := s.repo.GetFinancialDataByPagination(ctx, params.Page, params.PageSize)
	if err != nil {
		s.logger.Error("cannot GetFinancialDataByPagination",
			zap.Error(err),
			zap.String("service", "financialService"),
			zap.String("method", "GetFinancialDataList"),
		)
		return response.Error("something went wrong", http.StatusInternalServerError, nil)
	}
	res := make([]SingleFinancialDataResult, len(models))
	for i, m := range models {
		res[i] = SingleFinancialDataResult{
			ID:              m.ID.Hex(),
			SeriesReference: m.SeriesReference,
			Period:          m.Period,
			DataValue:       m.DataValue,
			Suppressed:      m.Suppressed,
			Status:          m.Status,
			Units:           m.Units,
			Magnitude:       m.Magnitude,
			Subject:         m.Subject,
			Group:           m.Group,
			SeriesTitle1:    m.SeriesTitle1,
			SeriesTitle2:    m.SeriesTitle2,
			SeriesTitle3:    m.SeriesTitle3,
			SeriesTitle4:    m.SeriesTitle4,
			SeriesTitle5:    m.SeriesTitle5,
		}
	}
	return response.Success(res, "")
}

func (s *Service) CreateFinancialDataByUser(
	ctx context.Context,
	params CreateFinancialDataParams,
) (apiResponse response.ApiResponse, statusCode int) {
	id, err := s.CreateFinancialData(ctx, FinacialModel{
		SeriesReference: params.SeriesReference,
		Period:          params.Period,
		DataValue:       params.DataValue,
		Suppressed:      params.Suppressed,
		Status:          params.Status,
		Units:           params.Units,
		Magnitude:       params.Magnitude,
		Subject:         params.Subject,
		Group:           params.Group,
		SeriesTitle1:    params.SeriesTitle1,
		SeriesTitle2:    params.SeriesTitle2,
		SeriesTitle3:    params.SeriesTitle3,
		SeriesTitle4:    params.SeriesTitle4,
		SeriesTitle5:    params.SeriesTitle5,
	})
	if err != nil {
		s.logger.Error("cannot CreateFinancialData",
			zap.Error(err),
			zap.String("service", "financialService"),
			zap.String("method", "CreateFinancialDataByUser"),
		)
		return response.Error("something went wrong", http.StatusInternalServerError, nil)
	}
	res := make(map[string]string)
	res["id"] = id
	return response.Success(res, "")
}

func (s *Service) CreateFinancialData(
	ctx context.Context,
	data FinacialModel,
) (string, error) {
	return s.repo.CreateFinancialData(ctx, data)
}

func (s *Service) UpdateFinancialData(
	ctx context.Context,
	params UpdateFinancialDataParams,
) (apiResponse response.ApiResponse, statusCode int) {
	_, err := s.repo.GetFinancialDataByID(ctx, params.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		s.logger.Error("cannot GetFinancialDataByID",
			zap.Error(err),
			zap.String("service", "financialService"),
			zap.String("method", "UpdateFinancialDataByUser"),
			zap.String("id", params.ID),
		)
		return response.Error("something went wrong", http.StatusInternalServerError, nil)
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return response.Error("not found", http.StatusNotFound, nil)
	}
	err = s.repo.UpdateFinancialData(ctx, params.ID, FinancialUpdateModel{
		SeriesReference: params.SeriesReference,
		Period:          params.Period,
		DataValue:       params.DataValue,
		Suppressed:      params.Suppressed,
		Status:          params.Status,
		Units:           params.Units,
		Magnitude:       params.Magnitude,
		Subject:         params.Subject,
		Group:           params.Group,
		SeriesTitle1:    params.SeriesTitle1,
		SeriesTitle2:    params.SeriesTitle2,
		SeriesTitle3:    params.SeriesTitle3,
		SeriesTitle4:    params.SeriesTitle4,
		SeriesTitle5:    params.SeriesTitle5,
	})
	if err != nil {
		s.logger.Error("cannot UpdateFinancialData",
			zap.Error(err),
			zap.String("service", "financialService"),
			zap.String("method", "UpdateFinancialDataByUser"),
			zap.String("id", params.ID),
		)
		return response.Error("something went wrong", http.StatusInternalServerError, nil)
	}
	res := make(map[string]string)
	return response.Success(res, "")
}

func (s *Service) DeleteFinancialData(
	ctx context.Context,
	params DeleteFinancialDataParams,
) (apiResponse response.ApiResponse, statusCode int) {
	_, err := s.repo.GetFinancialDataByID(ctx, params.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		s.logger.Error("cannot GetFinancialDataByID",
			zap.Error(err),
			zap.String("service", "financialService"),
			zap.String("method", "DeleteFinancialDataByUser"),
			zap.String("id", params.ID),
		)
		return response.Error("something went wrong", http.StatusInternalServerError, nil)
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return response.Error("not found", http.StatusNotFound, nil)
	}
	err = s.repo.DeleteFinancialData(ctx, params.ID)
	if err != nil {
		s.logger.Error("cannot GetFinancialDataByID",
			zap.Error(err),
			zap.String("service", "financialService"),
			zap.String("method", "DeleteFinancialDataByUser"),
			zap.String("id", params.ID),
		)
		return response.Error("something went wrong", http.StatusInternalServerError, nil)
	}
	res := make(map[string]string)
	return response.Success(res, "")
}

func NewService(
	repo *Repository,
	logger *zap.Logger,
) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
