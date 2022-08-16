package books

import (
	r "readwise-backend/packages/core/repository"
	"readwise-backend/packages/domain/authors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookModel struct {
	ID                         primitive.ObjectID    `bson:"_id,omitempty"`
	Title                      string                `bson:"title,omitempty"`
	Description                string                `bson:"description,omitempty"`
	Created_at                 time.Time             `bson:"created_at,omitempty"`
	Author                     *authors.AuthorSubset `bson:"author,omitempty"`
	r.RepositoryModel[BookDto] `bson:"-"`
}

func (b BookModel) GetID() primitive.ObjectID {
	return b.ID
}

func (b BookModel) ToEntity() BookDto {
	return BookDto{
		ID:          b.ID.Hex(),
		Title:       b.Title,
		Description: b.Description,
		Created_at:  b.Created_at,
		Author:      b.Author,
	}
}

type BookDto struct {
	ID                                  string                `json:"id,omitempty"`
	Title                               string                `json:"title,omitempty"`
	Description                         string                `json:"description,omitempty"`
	Created_at                          time.Time             `json:"created_at,omitempty"`
	Author                              *authors.AuthorSubset `json:"author,omitempty"`
	r.RepositoryDto[BookModel, BookDto] `json:"-"`
}

func (b BookDto) ToModel() BookModel {
	return BookModel{
		ID:          r.GetPrimitiveObjectIDFromString(b.ID),
		Title:       b.Title,
		Description: b.Description,
		Created_at:  b.Created_at,
		Author:      b.Author,
	}
}

func (b BookDto) SetID(ID interface{}) BookDto {
	IDToSet := ID.(primitive.ObjectID).Hex()
	b.ID = IDToSet
	return b
}

func (b BookDto) Init() BookDto {
	b.Created_at = time.Now()
	return b
}
