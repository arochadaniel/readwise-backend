package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetPrimitiveObjectIDFromString(id string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return primitive.NilObjectID
	}

	return objectID
}

func MapDtosToModels[Model RepositoryModel[Dto], Dto RepositoryDto[Model, Dto]](dtos []Dto) []Model {
	var models []Model
	for i := range dtos {
		models = append(models, dtos[i].ToModel())
	}

	return models
}

func MapModelsToDtos[Model RepositoryModel[Dto], Dto RepositoryDto[Model, Dto]](models []Model) []Dto {
	var dtos []Dto
	for i := range models {
		dtos = append(dtos, models[i].ToEntity())
	}

	return dtos
}
