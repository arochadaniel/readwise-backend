package repository

import (
	"context"
)

type RepositoryDto[Model any] interface {
	ToModel() Model
}

type RepositoryModel[Dto any] interface {
	ToEntity() Dto
}

type Repository[Model RepositoryModel[Dto], Dto RepositoryDto[Model]] interface {
	FindAll(ctx context.Context, m Model) []Dto
	FindOne(ctx context.Context, ID int) Dto
	CreateOne(ctx context.Context, m Model) string
	CreateMultiple(ctx context.Context, m []Model) []string
	UpdateOne(ctx context.Context, ID string, m Model) string
	UpdateBy(ctx context.Context, where Model, m Model) int64
	DeleteOne(ctx context.Context, where Model) int64
	DeleteBy(ctx context.Context, where Model) int64
}
