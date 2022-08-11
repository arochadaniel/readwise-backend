package controller

import (
	"context"
	r "readwise-backend/packages/core/repository"
)

type Controller[Model r.RepositoryModel[DTO], DTO r.RepositoryDto[Model, DTO]] interface {
	FindAll(ctx *context.Context)
	FindOne(ctx *context.Context)
	CreateOne(ctx *context.Context)
	CreateMultiple(ctx *context.Context)
	UpdateOne(ctx *context.Context)
	UpdateBy(ctx *context.Context)
	DeleteOne(ctx *context.Context)
	DeleteBy(ctx *context.Context)
}
