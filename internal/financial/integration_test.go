package financial_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"we-connect-test/internal/di"
	"we-connect-test/internal/financial"
	"we-connect-test/internal/handler/api"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestIndex(t *testing.T) {
	container := di.NewContainer()
	logger, err := container.GetLogger()
	assert.Nil(t, err)
	cfg := container.GetCfg()
	dbName := cfg.GetString("mongodb.dbname")
	mongoDBClient, err := container.GetMongoDBClient()
	assert.Nil(t, err)
	//first we empty the db collection
	ctx := context.Background()
	coll := mongoDBClient.Database(dbName).Collection("financialData")
	_, err = coll.DeleteMany(ctx, bson.M{})
	assert.Nil(t, err)

	financialService := container.GetFinancialService()
	httpServer := api.NewHttpServer(api.Services{
		Cfg:              cfg,
		FinancialService: financialService,
	}, logger)
	engine := httpServer.GetEngine()

	//first we insert 4 docs to db
	id1, err := financialService.CreateFinancialData(ctx, financial.FinacialModel{
		SeriesReference: "sr1",
		Period:          "period1",
		DataValue:       "dataValue1",
		Suppressed:      "suppressed1",
		Status:          "status1",
		Units:           "units1",
		Magnitude:       "magnitude1",
		Subject:         "subject1",
		Group:           "group1",
		SeriesTitle1:    "seriesTitle11",
		SeriesTitle2:    "seriesTitle21",
		SeriesTitle3:    "seriesTitle31",
		SeriesTitle4:    "seriesTitle41",
		SeriesTitle5:    "seriesTitle51",
	})
	assert.Nil(t, err)
	id2, err := financialService.CreateFinancialData(ctx, financial.FinacialModel{
		SeriesReference: "sr2",
		Period:          "period2",
		DataValue:       "dataValue2",
		Suppressed:      "suppressed2",
		Status:          "status2",
		Units:           "units2",
		Magnitude:       "magnitude2",
		Subject:         "subject2",
		Group:           "group2",
		SeriesTitle1:    "seriesTitle12",
		SeriesTitle2:    "seriesTitle22",
		SeriesTitle3:    "seriesTitle32",
		SeriesTitle4:    "seriesTitle42",
		SeriesTitle5:    "seriesTitle52",
	})
	assert.Nil(t, err)
	id3, err := financialService.CreateFinancialData(ctx, financial.FinacialModel{
		SeriesReference: "sr3",
		Period:          "period3",
		DataValue:       "dataValue3",
		Suppressed:      "suppressed3",
		Status:          "status3",
		Units:           "units3",
		Magnitude:       "magnitude3",
		Subject:         "subject3",
		Group:           "group3",
		SeriesTitle1:    "seriesTitle13",
		SeriesTitle2:    "seriesTitle23",
		SeriesTitle3:    "seriesTitle33",
		SeriesTitle4:    "seriesTitle43",
		SeriesTitle5:    "seriesTitle53",
	})
	assert.Nil(t, err)
	id4, err := financialService.CreateFinancialData(ctx, financial.FinacialModel{
		SeriesReference: "sr4",
		Period:          "period4",
		DataValue:       "dataValue4",
		Suppressed:      "suppressed4",
		Status:          "status4",
		Units:           "units4",
		Magnitude:       "magnitude4",
		Subject:         "subject4",
		Group:           "group4",
		SeriesTitle1:    "seriesTitle14",
		SeriesTitle2:    "seriesTitle24",
		SeriesTitle3:    "seriesTitle34",
		SeriesTitle4:    "seriesTitle44",
		SeriesTitle5:    "seriesTitle54",
	})
	assert.Nil(t, err)

	//for page 1
	queryParams := url.Values{}
	queryParams.Set("page", "0")
	queryParams.Set("pageSize", "2")
	paramsString := queryParams.Encode()

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/financial?"+paramsString, nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, res.Code, http.StatusOK)
	result := struct {
		Status  bool
		Message string
		Data    []financial.SingleFinancialDataResult
	}{}
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.Equal(t, len(result.Data), 2)
	fd1 := result.Data[0]
	assert.Equal(t, fd1.ID, id1)
	assert.Equal(t, fd1.SeriesReference, "sr1")
	assert.Equal(t, fd1.Period, "period1")
	assert.Equal(t, fd1.DataValue, "dataValue1")
	assert.Equal(t, fd1.Suppressed, "suppressed1")
	assert.Equal(t, fd1.Status, "status1")
	assert.Equal(t, fd1.Units, "units1")
	assert.Equal(t, fd1.Magnitude, "magnitude1")
	assert.Equal(t, fd1.Subject, "subject1")
	assert.Equal(t, fd1.Group, "group1")
	assert.Equal(t, fd1.SeriesTitle1, "seriesTitle11")
	assert.Equal(t, fd1.SeriesTitle2, "seriesTitle21")
	assert.Equal(t, fd1.SeriesTitle3, "seriesTitle31")
	assert.Equal(t, fd1.SeriesTitle4, "seriesTitle41")
	assert.Equal(t, fd1.SeriesTitle5, "seriesTitle51")
	fd2 := result.Data[1]
	assert.Equal(t, fd2.ID, id2)
	assert.Equal(t, fd2.SeriesReference, "sr2")
	assert.Equal(t, fd2.Period, "period2")
	assert.Equal(t, fd2.DataValue, "dataValue2")
	assert.Equal(t, fd2.Suppressed, "suppressed2")
	assert.Equal(t, fd2.Status, "status2")
	assert.Equal(t, fd2.Units, "units2")
	assert.Equal(t, fd2.Magnitude, "magnitude2")
	assert.Equal(t, fd2.Subject, "subject2")
	assert.Equal(t, fd2.Group, "group2")
	assert.Equal(t, fd2.SeriesTitle1, "seriesTitle12")
	assert.Equal(t, fd2.SeriesTitle2, "seriesTitle22")
	assert.Equal(t, fd2.SeriesTitle3, "seriesTitle32")
	assert.Equal(t, fd2.SeriesTitle4, "seriesTitle42")
	assert.Equal(t, fd2.SeriesTitle5, "seriesTitle52")

	//for page 2
	queryParams = url.Values{}
	queryParams.Set("page", "1")
	queryParams.Set("pageSize", "2")
	paramsString = queryParams.Encode()

	res = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/v1/financial?"+paramsString, nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, res.Code, http.StatusOK)
	result = struct {
		Status  bool
		Message string
		Data    []financial.SingleFinancialDataResult
	}{}
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.Equal(t, len(result.Data), 2)
	fd3 := result.Data[0]
	assert.Equal(t, fd3.ID, id3)
	assert.Equal(t, fd3.SeriesReference, "sr3")
	assert.Equal(t, fd3.Period, "period3")
	assert.Equal(t, fd3.DataValue, "dataValue3")
	assert.Equal(t, fd3.Suppressed, "suppressed3")
	assert.Equal(t, fd3.Status, "status3")
	assert.Equal(t, fd3.Units, "units3")
	assert.Equal(t, fd3.Magnitude, "magnitude3")
	assert.Equal(t, fd3.Subject, "subject3")
	assert.Equal(t, fd3.Group, "group3")
	assert.Equal(t, fd3.SeriesTitle1, "seriesTitle13")
	assert.Equal(t, fd3.SeriesTitle2, "seriesTitle23")
	assert.Equal(t, fd3.SeriesTitle3, "seriesTitle33")
	assert.Equal(t, fd3.SeriesTitle4, "seriesTitle43")
	assert.Equal(t, fd3.SeriesTitle5, "seriesTitle53")
	fd4 := result.Data[1]
	assert.Equal(t, fd4.ID, id4)
	assert.Equal(t, fd4.SeriesReference, "sr4")
	assert.Equal(t, fd4.Period, "period4")
	assert.Equal(t, fd4.DataValue, "dataValue4")
	assert.Equal(t, fd4.Suppressed, "suppressed4")
	assert.Equal(t, fd4.Status, "status4")
	assert.Equal(t, fd4.Units, "units4")
	assert.Equal(t, fd4.Magnitude, "magnitude4")
	assert.Equal(t, fd4.Subject, "subject4")
	assert.Equal(t, fd4.Group, "group4")
	assert.Equal(t, fd4.SeriesTitle1, "seriesTitle14")
	assert.Equal(t, fd4.SeriesTitle2, "seriesTitle24")
	assert.Equal(t, fd4.SeriesTitle3, "seriesTitle34")
	assert.Equal(t, fd4.SeriesTitle4, "seriesTitle44")
	assert.Equal(t, fd4.SeriesTitle5, "seriesTitle54")

	//for page 3
	queryParams = url.Values{}
	queryParams.Set("page", "2")
	queryParams.Set("pageSize", "2")
	paramsString = queryParams.Encode()

	res = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/v1/financial?"+paramsString, nil)
	engine.ServeHTTP(res, req)
	assert.Equal(t, res.Code, http.StatusOK)
	result = struct {
		Status  bool
		Message string
		Data    []financial.SingleFinancialDataResult
	}{}
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.Equal(t, len(result.Data), 0)
}

func TestCreate_Update_Delete(t *testing.T) {
	//first test is create
	container := di.NewContainer()
	logger, err := container.GetLogger()
	assert.Nil(t, err)
	cfg := container.GetCfg()
	dbName := cfg.GetString("mongodb.dbname")
	mongoDBClient, err := container.GetMongoDBClient()
	assert.Nil(t, err)
	//first we empty the db collection
	ctx := context.Background()
	coll := mongoDBClient.Database(dbName).Collection("financialData")
	_, err = coll.DeleteMany(ctx, bson.M{})
	assert.Nil(t, err)

	financialService := container.GetFinancialService()
	httpServer := api.NewHttpServer(api.Services{
		Cfg:              cfg,
		FinancialService: financialService,
	}, logger)
	engine := httpServer.GetEngine()
	res := httptest.NewRecorder()
	data := `{
		"seriesReference":"newSr",
		"period":"newPeriod",
		"dataValue":"newDataValue",
		"suppressed":"newSuppressed",
		"status":"newStatus",
		"units":"newUnits",
		"magnitude":"newMagnitude",
		"subject":"newSubject",
		"group":"newGroup",
		"seriesTitle1":"newSeriesTitle1",
		"seriesTitle2":"newSeriesTitle2",
		"seriesTitle3":"newSeriesTitle3",
		"seriesTitle4":"newSeriesTitle4",
		"seriesTitle5":"newSeriesTitle5"
	}`
	body := []byte(data)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/financial/create", bytes.NewReader(body))
	engine.ServeHTTP(res, req)
	assert.Equal(t, res.Code, http.StatusOK)
	result := struct {
		Status  bool
		Message string
		Data    map[string]string
	}{}
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	id, ok := result.Data["id"]
	assert.True(t, ok)
	//here we get the data by id from the db to check if it is inserted
	coll = mongoDBClient.Database(dbName).Collection("financialData")
	objectID, err := primitive.ObjectIDFromHex(id)
	assert.Nil(t, err)
	filter := bson.D{{"_id", objectID}}
	dbResult := coll.FindOne(ctx, filter)
	assert.Nil(t, dbResult.Err())
	m := financial.FinacialModel{}
	err = dbResult.Decode(&m)
	assert.Nil(t, err)
	assert.Equal(t, m.ID.Hex(), id)
	assert.Equal(t, m.SeriesReference, "newSr")
	assert.Equal(t, m.Period, "newPeriod")
	assert.Equal(t, m.DataValue, "newDataValue")
	assert.Equal(t, m.Suppressed, "newSuppressed")
	assert.Equal(t, m.Status, "newStatus")
	assert.Equal(t, m.Units, "newUnits")
	assert.Equal(t, m.Magnitude, "newMagnitude")
	assert.Equal(t, m.Subject, "newSubject")
	assert.Equal(t, m.Group, "newGroup")
	assert.Equal(t, m.SeriesTitle1, "newSeriesTitle1")
	assert.Equal(t, m.SeriesTitle2, "newSeriesTitle2")
	assert.Equal(t, m.SeriesTitle3, "newSeriesTitle3")
	assert.Equal(t, m.SeriesTitle4, "newSeriesTitle4")
	assert.Equal(t, m.SeriesTitle5, "newSeriesTitle5")

	//now we try to update the same doc
	res = httptest.NewRecorder()
	data = fmt.Sprintf(`{
		"id":"%s",
		"seriesReference":"updatedSr",
		"period":"updatedPeriod",
		"dataValue":"updatedDataValue",
		"suppressed":"updatedSuppressed",
		"status":"updatedStatus",
		"units":"updatedUnits",
		"magnitude":"updatedMagnitude",
		"subject":"updatedSubject",
		"group":"updatedGroup",
		"seriesTitle1":"updatedSeriesTitle1",
		"seriesTitle2":"updatedSeriesTitle2",
		"seriesTitle3":"updatedSeriesTitle3",
		"seriesTitle4":"updatedSeriesTitle4"
	}`, id)
	body = []byte(data)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/financial/update", bytes.NewReader(body))
	engine.ServeHTTP(res, req)
	assert.Equal(t, res.Code, http.StatusOK)
	result = struct {
		Status  bool
		Message string
		Data    map[string]string
	}{}
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	//here we get the data by id from the db to check if it is updated
	coll = mongoDBClient.Database(dbName).Collection("financialData")
	assert.Nil(t, err)
	dbResult = coll.FindOne(ctx, filter)
	assert.Nil(t, dbResult.Err())
	m = financial.FinacialModel{}
	err = dbResult.Decode(&m)
	assert.Nil(t, err)
	assert.Equal(t, m.ID.Hex(), id)
	assert.Equal(t, m.SeriesReference, "updatedSr")
	assert.Equal(t, m.Period, "updatedPeriod")
	assert.Equal(t, m.DataValue, "updatedDataValue")
	assert.Equal(t, m.Suppressed, "updatedSuppressed")
	assert.Equal(t, m.Status, "updatedStatus")
	assert.Equal(t, m.Units, "updatedUnits")
	assert.Equal(t, m.Magnitude, "updatedMagnitude")
	assert.Equal(t, m.Subject, "updatedSubject")
	assert.Equal(t, m.Group, "updatedGroup")
	assert.Equal(t, m.SeriesTitle1, "updatedSeriesTitle1")
	assert.Equal(t, m.SeriesTitle2, "updatedSeriesTitle2")
	assert.Equal(t, m.SeriesTitle3, "updatedSeriesTitle3")
	assert.Equal(t, m.SeriesTitle4, "updatedSeriesTitle4")
	//this should not be updated because we did not sent it
	assert.Equal(t, m.SeriesTitle5, "newSeriesTitle5")

	//now we try to delete the same doc
	res = httptest.NewRecorder()
	data = fmt.Sprintf(`{
		"id":"%s"
	}`, id)
	body = []byte(data)
	req = httptest.NewRequest(http.MethodPost, "/api/v1/financial/delete", bytes.NewReader(body))
	engine.ServeHTTP(res, req)
	assert.Equal(t, res.Code, http.StatusOK)
	result = struct {
		Status  bool
		Message string
		Data    map[string]string
	}{}
	err = json.Unmarshal(res.Body.Bytes(), &result)
	assert.Nil(t, err)
	//here we get the data by id from the db to check if it is deleted
	coll = mongoDBClient.Database(dbName).Collection("financialData")
	assert.Nil(t, err)
	dbResult = coll.FindOne(ctx, filter)
	assert.Equal(t, dbResult.Err(), mongo.ErrNoDocuments)
}
