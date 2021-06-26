package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbRepository struct {
	DB  *mongo.Database
	Ctx context.Context
}

func NewMongoDbRepository(connection *mongo.Database, ctx context.Context) *MongoDbRepository {
	return &MongoDbRepository{DB: connection, Ctx: ctx}
}

func (r *MongoDbRepository) newCollection(collection string) *mongo.Collection {
	return r.DB.Collection(collection)
}

func (r *MongoDbRepository) CreateUser(email string, password string, role string) (*mongo.InsertOneResult, error) {
	collection := r.newCollection("users")
	res, err := collection.InsertOne(context.Background(), bson.D{{"email", email}, {"password", password}, {"role", role}})
	if err != nil {
		return nil, err
	}
	return res, nil
}
