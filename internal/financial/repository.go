package financial

import (
	"context"
	"we-connect-test/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	financialDataCollectionName = "financialData"
)

type FinacialModel struct {
	ID              primitive.ObjectID `bson:"_id"`
	SeriesReference string             `bson:"seriesReference"`
	Period          string             `bson:"period"`
	DataValue       string             `bson:"dataValue"`
	Suppressed      string             `bson:"suppressed"`
	Status          string             `bson:"status"`
	Units           string             `bson:"units"`
	Magnitude       string             `bson:"magnitude"`
	Subject         string             `bson:"subject"`
	Group           string             `bson:"group"`
	SeriesTitle1    string             `bson:"seriesTitle1"`
	SeriesTitle2    string             `bson:"seriesTitle2"`
	SeriesTitle3    string             `bson:"seriesTitle3"`
	SeriesTitle4    string             `bson:"seriesTitle4"`
	SeriesTitle5    string             `bson:"seriesTitle5"`
}

type FinancialUpdateModel struct {
	SeriesReference *string
	Period          *string
	DataValue       *string
	Suppressed      *string
	Status          *string
	Units           *string
	Magnitude       *string
	Subject         *string
	Group           *string
	SeriesTitle1    *string
	SeriesTitle2    *string
	SeriesTitle3    *string
	SeriesTitle4    *string
	SeriesTitle5    *string
}

type Repository struct {
	dbName        string
	mongoDBClient *mongo.Client
}

func (r *Repository) GetFinancialDataByPagination(ctx context.Context, page, pageSize int) ([]FinacialModel, error) {
	opts := options.Find().
		SetLimit(int64(pageSize)).
		SetSkip(int64(page * pageSize))
	coll := r.mongoDBClient.Database(r.dbName).Collection(financialDataCollectionName)
	cursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	var results []FinacialModel
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *Repository) CreateFinancialData(ctx context.Context, m FinacialModel) (string, error) {
	m.ID = primitive.NewObjectID()
	coll := r.mongoDBClient.Database(r.dbName).Collection(financialDataCollectionName)
	result, err := coll.InsertOne(ctx, m)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (r *Repository) GetFinancialDataByID(ctx context.Context, id string) (FinacialModel, error) {
	coll := r.mongoDBClient.Database(r.dbName).Collection(financialDataCollectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return FinacialModel{}, err
	}
	filter := bson.D{{"_id", objectID}}
	res := coll.FindOne(ctx, filter)
	if res.Err() != nil {
		return FinacialModel{}, res.Err()
	}
	m := FinacialModel{}
	err = res.Decode(&m)
	return m, err
}

func (r *Repository) UpdateFinancialData(ctx context.Context, id string, m FinancialUpdateModel) error {
	coll := r.mongoDBClient.Database(r.dbName).Collection(financialDataCollectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectID}}
	set := bson.D{}
	if m.SeriesReference != nil {
		set = append(set, bson.E{"seriesReference", m.SeriesReference})
	}
	if m.Period != nil {
		set = append(set, bson.E{"period", m.Period})
	}
	if m.DataValue != nil {
		set = append(set, bson.E{"dataValue", m.DataValue})
	}
	if m.Suppressed != nil {
		set = append(set, bson.E{"suppressed", m.Suppressed})
	}
	if m.Status != nil {
		set = append(set, bson.E{"status", m.Status})
	}
	if m.Units != nil {
		set = append(set, bson.E{"units", m.Units})
	}
	if m.Magnitude != nil {
		set = append(set, bson.E{"magnitude", m.Magnitude})
	}
	if m.Subject != nil {
		set = append(set, bson.E{"Subject", m.Subject})
	}
	if m.Group != nil {
		set = append(set, bson.E{"group", m.Group})
	}
	if m.SeriesTitle1 != nil {
		set = append(set, bson.E{"seriesTitle1", m.SeriesTitle1})
	}
	if m.SeriesTitle2 != nil {
		set = append(set, bson.E{"seriesTitle2", m.SeriesTitle2})
	}
	if m.SeriesTitle3 != nil {
		set = append(set, bson.E{"seriesTitle3", m.SeriesTitle3})
	}
	if m.SeriesTitle4 != nil {
		set = append(set, bson.E{"seriesTitle4", m.SeriesTitle4})
	}
	if m.SeriesTitle5 != nil {
		set = append(set, bson.E{"seriesTitle5", m.SeriesTitle5})
	}
	update := bson.D{{"$set", set}}
	_, err = coll.UpdateOne(ctx, filter, update)
	return err
}

func (r *Repository) DeleteFinancialData(ctx context.Context, id string) error {
	coll := r.mongoDBClient.Database(r.dbName).Collection(financialDataCollectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectID}}
	_, err = coll.DeleteOne(ctx, filter)
	return err
}

func NewRepository(cfg *config.Cfg, mongoDBClient *mongo.Client) *Repository {
	return &Repository{
		dbName:        cfg.GetString("mongodb.dbname"),
		mongoDBClient: mongoDBClient,
	}

}
