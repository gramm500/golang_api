package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InitMongoConnection(host string, db string) (*mongo.Database, context.Context) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(c, options.Client().ApplyURI(host))
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err = client.Disconnect(c); err != nil {
	//		panic(err)
	//	}
	//}()

	return client.Database(db), c
}
