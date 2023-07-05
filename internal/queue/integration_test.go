package queue_test

import (
	"context"
	"testing"
	"time"
	"we-connect-test/internal/di"
	"we-connect-test/internal/financial"
	"we-connect-test/internal/queue"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestManager_Run(t *testing.T) {
	//we do this to be sure that when tests run cuncurrently
	//the db data changes in other tests does not affect this one
	time.Sleep(3 * time.Second)
	container := di.NewContainer()
	cfg := container.GetCfg()
	dbName := cfg.GetString("mongodb.dbname")
	mongoDBClient, err := container.GetMongoDBClient()
	assert.Nil(t, err)
	//first we empty the db collection
	ctx := context.Background()
	coll := mongoDBClient.Database(dbName).Collection("financialData")
	_, err = coll.DeleteMany(ctx, bson.M{})
	assert.Nil(t, err)

	logger, err := container.GetLogger()
	assert.Nil(t, err)
	financialService := container.GetFinancialService()
	manager := queue.NewManager(financialService, logger)
	filePath := "./data_test.csv"
	err = manager.Run(ctx, filePath, 5)
	assert.Nil(t, err)

	//sleep here so all goroutine finish their job
	time.Sleep(5 * time.Second)
	coll = mongoDBClient.Database(dbName).Collection("financialData")
	cursor, err := coll.Find(ctx, bson.M{})
	assert.Nil(t, err)
	var results []financial.FinacialModel
	err = cursor.All(ctx, &results)
	assert.Nil(t, err)
	assert.Equal(t, len(results), 10)

	for _, res := range results {
		assert.NotEmpty(t, res.Period)
		assert.NotEmpty(t, res.DataValue)
		assert.NotEmpty(t, res.Status)
		assert.NotEmpty(t, res.Units)
		assert.NotEmpty(t, res.Magnitude)
		assert.NotEmpty(t, res.Subject)
		assert.NotEmpty(t, res.Group)
		assert.NotEmpty(t, res.SeriesTitle1)
		assert.NotEmpty(t, res.SeriesTitle2)
		assert.NotEmpty(t, res.SeriesTitle3)
		assert.NotEmpty(t, res.SeriesTitle4)
	}
}
