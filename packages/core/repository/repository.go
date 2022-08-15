package repository

import (
	"context"
)

type DatabaseContainer[CDB any] struct {
	DB *CDB
}
type RepositoryDto[Model any, Dto any] interface {
	SetID(ID interface{}) Dto
	ToModel() Model
}

type RepositoryModel[Dto any] interface {
	ToEntity() Dto
}

type Repository[Model RepositoryModel[Dto], Dto RepositoryDto[Model, Dto]] interface {
	FindAll(ctx context.Context, m Model) ([]Dto, error)
	FindOne(ctx context.Context, m Model) (Dto, error)
	CreateOne(ctx context.Context, m Model) Dto
	CreateMultiple(ctx context.Context, m []Model) []Dto
	UpdateOne(ctx context.Context, ID string, m Model) Dto
	UpdateBy(ctx context.Context, where Model, m Model) int64
	DeleteOne(ctx context.Context, where Model) int64
	DeleteBy(ctx context.Context, where Model) int64
}
