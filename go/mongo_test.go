package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// TestMongo ...
func TestMongo(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	filter := bson.M{"name": "pi"}
	var result bson.M
	collection := client.Database("testing").Collection("numbers")
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		t.Error(err)
	}
	type PI struct {
		name string
		pi   float64
	}
	if result["name"].(string) != "pi" {
		t.Fail()
	}

	fmt.Println(result["name"].(string) == "pi")

}
