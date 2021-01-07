package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoDBConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	return client, nil
}

func GetMongoDBCollection(DBname string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDBConnection()

	if err != nil {
		log.Fatal("DB is not connected")
		return nil, err
	}

	collection := client.Database(DBname).Collection(CollectionName)
	log.Fatal("DB is connected")
	return collection, nil
}
