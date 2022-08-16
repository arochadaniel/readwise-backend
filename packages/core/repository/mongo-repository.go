package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatabaseContainer struct {
	DatabaseContainer[mongo.Database]
}

type MongoRepository[Model RepositoryModel[Dto], Dto RepositoryDto[Model, Dto]] struct {
	CollectionName string
	DBContainer    *MongoDatabaseContainer
	Repository[Model, Dto]
}

func (r *MongoRepository[Model, DTO]) FindAll(ctx context.Context, m Model) ([]DTO, error) {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
	cursor, err := collection.Find(ctx, m)

	if err != nil {
		log := fmt.Sprintf("no items found in collection %s, err: %d", r.CollectionName, err)
		panic(log)
	}

	result := []Model{}
	cursorErr := cursor.All(ctx, &result)

	if cursorErr != nil {
		panic(cursorErr)
	}

	response := MapModelsToDtos[Model, DTO](result)

	return response, nil
}

func (r *MongoRepository[Model, DTO]) FindOne(ctx context.Context, m Model) (DTO, error) {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
	result := collection.FindOne(ctx, map[string]primitive.ObjectID{"_id": m.GetID()})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			var emptyDTO DTO
			return emptyDTO, result.Err()
		}

		log := fmt.Sprintf("no items found in collection %s, err: %d", r.CollectionName, result.Err())
		panic(log)
	}

	var response Model
	err := result.Decode(&response)

	if err != nil {
		panic(err)
	}

	return response.ToEntity(), nil
}

func (r *MongoRepository[Model, DTO]) CreateOne(ctx context.Context, m Model) DTO {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
	result, err := collection.InsertOne(ctx, m)

	if err != nil {
		panic(err)
	}

	response := m.ToEntity().SetID(result.InsertedID)

	return response
}

func (r *MongoRepository[Model, DTO]) CreateMultiple(ctx context.Context, m []Model) []DTO {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
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
		toAppend := m[i].ToEntity().SetID(v)
		response = append(response, toAppend)
	}

	return response
}

func (r *MongoRepository[Model, DTO]) UpdateOne(ctx context.Context, id string, m Model) DTO {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
	primitiveID := GetPrimitiveObjectIDFromString(id)
	_, err := collection.UpdateByID(ctx, primitiveID, map[string]Model{"$set": m})

	if err != nil {
		panic(err)
	}

	response := m.ToEntity().SetID(primitiveID)

	return response
}

func (r *MongoRepository[Model, DTO]) UpdateBy(ctx context.Context, where Model, m Model) int64 {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
	result, err := collection.UpdateMany(ctx, where, map[string]Model{"$set": m})

	if err != nil {
		panic(err)
	}

	return result.ModifiedCount
}

func (r *MongoRepository[Model, DTO]) DeleteOne(ctx context.Context, filter Model) int64 {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		panic(err)
	}

	return result.DeletedCount
}

func (r *MongoRepository[Model, DTO]) DeleteBy(ctx context.Context, where Model) int64 {
	collection := r.DBContainer.DB.Collection(r.CollectionName)
	result, err := collection.DeleteMany(ctx, where)

	if err != nil {
		panic(err)
	}

	return result.DeletedCount
}
