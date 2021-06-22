package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_api/models"
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

func (r *MongoDbRepository) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	collection := r.newCollection("users")
	res, err := collection.InsertOne(r.Ctx, bson.D{{"email", user.Email}, {"password", user.Password}, {"role", user.Role}})
	if err != nil {
		return nil, err
	}
	return res, nil
}