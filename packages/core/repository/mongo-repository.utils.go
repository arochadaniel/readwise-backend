package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetUpMongoDatabaseContainer() *MongoDatabaseContainer {
	db := GetMongoClient(context.Background()).Database("readwise")
	dbContainer := &MongoDatabaseContainer{}
	dbContainer.DB = db
	return dbContainer
}

func GetMongoClient(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://database:27017"))

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
