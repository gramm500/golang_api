package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Role  string             `json:"role,omitempty" bson:"role,omitempty"`
}

var client *mongo.Client

func main() {
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		panic(err)
	}
	client, err := mongo.Connect(ctx, "")
	fmt.Println(client)
}
