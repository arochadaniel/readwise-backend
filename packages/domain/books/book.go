package books

import (
	r "readwise-backend/packages/core/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookModel struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty"`
	Title                      string             `bson:"title,omitempty"`
	Description                string             `bson:"description,omitempty"`
	Created_at                 time.Time          `bson:"created_at,omitempty"`
	r.RepositoryModel[BookDto] `bson:"-"`
}

func (b BookModel) ToEntity() BookDto {
	return BookDto{
		ID:          b.ID.Hex(),
		Title:       b.Title,
		Description: b.Description,
		Created_at:  b.Created_at,
	}
}

type BookDto struct {
	ID                         string    `json:"id"`
	Title                      string    `json:"title"`
	Description                string    `json:"description"`
	Created_at                 time.Time `json:"created_at"`
	r.RepositoryDto[BookModel] `json:"-"`
}

func (b BookDto) ToModel() BookModel {
	return BookModel{
		ID:          r.GetPrimitiveObjectIDFromString(b.ID),
		Title:       b.Title,
		Description: b.Description,
		Created_at:  b.Created_at,
	}
}
