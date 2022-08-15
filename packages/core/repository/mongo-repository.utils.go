package repository

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetUpMongoDatabaseContainer() *MongoDatabaseContainer {
	db := GetMongoClient(context.Background()).Database("readwise")
	dbContainer := MongoDatabaseContainer{}
	dbContainer.DatabaseContainer.DB = db
	return &dbContainer
}

func GetMongoClient(ctx context.Context) *mongo.Client {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		panic("INVALID/EMPTY MONGO DB URI")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic("error instantiating mongo client")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}

func GetMongoCollection(collectionName string, dbContainer *MongoDatabaseContainer) *mongo.Collection {
	collection := dbContainer.DB.Collection(collectionName)
	return collection
}
