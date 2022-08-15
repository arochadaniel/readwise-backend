package apptesting

import (
	"context"
	"readwise-backend/packages/core/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupTestingDB() (*repository.MongoDatabaseContainer, context.Context) {
	var db *repository.MongoDatabaseContainer
	ctx := context.TODO()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic("error instantiating mongo client")
	}

	if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	db = &repository.MongoDatabaseContainer{}
	db.DatabaseContainer.DB = dbClient.Database("readwise")

	return db, ctx
}
