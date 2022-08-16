package authors

import (
	"readwise-backend/packages/core/controller"
)

type AuthorsController struct {
	controller.GinController[AuthorModel, AuthorDto, AuthorsRepository]
}

func NewAuthorsController(r AuthorsRepository) *AuthorsController {
	c := AuthorsController{}
	c.Repository = r
	return &c
}
