package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang_api/db"
	"golang_api/repositories"
)

type DI struct {
	Mongo           *mongo.Database
	MongoRepository *repositories.MongoDbRepository
	Config          Config
}

func NewDI(config Config) *DI {
	di := &DI{Config: config}

	database, c := db.InitMongoConnection(config.MongoHost, config.MongoDB)
	di.Mongo = database

	di.MongoRepository = repositories.NewMongoDbRepository(di.Mongo, c)

	return di
}
