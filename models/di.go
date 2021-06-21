package models

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang_api/db"
)

type DI struct {
	Mongo  *mongo.Database
	Config Config
}

func NewDI(config Config) *DI {
	di := &DI{Config: config}

	di.Mongo = db.InitMongoConnection(config.MongoHost, config.MongoDB)

	return di
}
