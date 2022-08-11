package controller

import (
	r "readwise-backend/packages/core/repository"
)

type GinController[Model r.RepositoryModel[DTO], DTO r.RepositoryDto[Model, DTO], R r.Repository[Model, DTO]] struct {
	Repository *R
	Controller[Model, DTO]
}
