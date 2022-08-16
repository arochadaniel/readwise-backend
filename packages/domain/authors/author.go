package authors

import (
	"readwise-backend/packages/core/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorSubset struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

type AuthorModel struct {
	ID                                    primitive.ObjectID `bson:"_id,omitempty"`
	Name                                  string             `bson:"name,omitempty"`
	Description                           string             `bson:"description,omitempty"`
	Created_at                            time.Time          `bson:"created_at,omitempty"`
	repository.RepositoryModel[AuthorDto] `bson:"-"`
}

func (a AuthorModel) GetID() primitive.ObjectID {
	return a.ID
}

func (a AuthorModel) ToEntity() AuthorDto {
	return AuthorDto{
		ID:          a.ID.Hex(),
		Name:        a.Name,
		Description: a.Description,
		Created_at:  a.Created_at,
	}
}

type AuthorDto struct {
	ID                                               string    `json:"id"`
	Name                                             string    `json:"name"`
	Description                                      string    `json:"description"`
	Created_at                                       time.Time `json:"created_at"`
	repository.RepositoryDto[AuthorModel, AuthorDto] `json:"-"`
}

func (a AuthorDto) ToModel() AuthorModel {
	return AuthorModel{
		ID:          repository.GetPrimitiveObjectIDFromString(a.ID),
		Name:        a.Name,
		Description: a.Description,
		Created_at:  a.Created_at,
	}
}

func (a AuthorDto) SetID(ID interface{}) AuthorDto {
	IDToSet := ID.(primitive.ObjectID).Hex()
	a.ID = IDToSet
	return a
}

func (a AuthorDto) Init() AuthorDto {
	a.Created_at = time.Now()
	return a
}
