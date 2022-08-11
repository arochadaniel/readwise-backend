package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient = GetMongoClient(context.Background())
var MongoDatabase = MongoClient.Database("readwise")

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

func GetMongoCollection(collectionName string) *mongo.Collection {
	collection := MongoDatabase.Collection(collectionName)
	return collection
}

type MongoRepository[Model RepositoryModel[Dto], Dto RepositoryDto[Model, Dto]] struct {
	CollectionName string
	Repository[Model, Dto]
}

func (r *MongoRepository[Model, DTO]) FindAll(ctx context.Context, d Model) []DTO {
	collection := MongoDatabase.Collection(r.CollectionName)
	cursor, err := collection.Find(ctx, d)

	if err != nil {
		log := fmt.Sprintf("no items found in collection %s", r.CollectionName)
		panic(log)
	}

	result := []Model{}
	cursorErr := cursor.All(ctx, &result)

	if cursorErr != nil {
		panic(cursorErr)
	}

	response := []DTO{}

	for i := range result {
		response = append(response, result[i].ToEntity())
	}

	return response
}

func (r *MongoRepository[Model, DTO]) FindOne(ctx context.Context, m Model) DTO {
	collection := MongoDatabase.Collection(r.CollectionName)
	result := collection.FindOne(ctx, m)

	if result.Err() != nil {
		log := fmt.Sprintf("no items found in collection %s", r.CollectionName)
		panic(log)
	}

	var response Model
	err := result.Decode(&response)

	if err != nil {
		panic(err)
	}

	return response.ToEntity()
}

func (r *MongoRepository[Model, DTO]) CreateOne(ctx context.Context, m Model) DTO {
	collection := MongoDatabase.Collection(r.CollectionName)
	result, err := collection.InsertOne(ctx, m)

	if err != nil {
		panic(err)
	}

	response := m.ToEntity().SetID(result.InsertedID)

	return response
}

func (r *MongoRepository[Model, DTO]) CreateMultiple(ctx context.Context, m []Model) []DTO {
	collection := MongoDatabase.Collection(r.CollectionName)
	var values []interface{}

	for i := range m {
		values = append(values, m[i])
	}

	result, err := collection.InsertMany(ctx, values)

	if err != nil {
		panic(err)
	}

	response := []DTO{}

	for i, v := range result.InsertedIDs {
		response = append(response, m[i].ToEntity().SetID(v))
	}

	return response
}

func (r *MongoRepository[Model, DTO]) UpdateOne(ctx context.Context, id string, m Model) DTO {
	collection := MongoDatabase.Collection(r.CollectionName)
	primitiveID := GetPrimitiveObjectIDFromString(id)
	_, err := collection.UpdateByID(ctx, primitiveID, map[string]Model{"$set": m})

	if err != nil {
		panic(err)
	}

	return m.ToEntity().SetID(primitiveID)
}

func (r *MongoRepository[Model, DTO]) UpdateBy(ctx context.Context, where Model, m Model) int64 {
	collection := MongoDatabase.Collection(r.CollectionName)
	result, err := collection.UpdateMany(ctx, where, map[string]Model{"$set": m})

	if err != nil {
		panic(err)
	}

	return result.ModifiedCount
}

func (r *MongoRepository[Model, DTO]) DeleteOne(ctx context.Context, filter Model) int64 {
	collection := MongoDatabase.Collection(r.CollectionName)
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		panic(err)
	}

	return result.DeletedCount
}

func (r *MongoRepository[Model, DTO]) DeleteBy(ctx context.Context, where Model) int64 {
	collection := MongoDatabase.Collection(r.CollectionName)
	result, err := collection.DeleteMany(ctx, where)

	if err != nil {
		panic(err)
	}

	return result.DeletedCount
}
