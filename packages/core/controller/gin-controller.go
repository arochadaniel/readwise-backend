package controller

import (
	r "readwise-backend/packages/core/repository"
)

type GinController[Model r.RepositoryModel[DTO], DTO r.RepositoryDto[Model]] struct {
	Repository *r.MongoRepository[Model, DTO]
	Controller[Model, DTO]
}
