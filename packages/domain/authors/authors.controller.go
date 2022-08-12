package authors

import (
	"net/http"
	"readwise-backend/packages/core/controller"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthorsController struct {
	controller.GinController[AuthorModel, AuthorDto, AuthorsRepository]
}

func NewAuthorsController(r *AuthorsRepository) *AuthorsController {
	c := AuthorsController{}
	c.Repository = r
	return &c
}

func (ac *AuthorsController) CreateOne(ctx *gin.Context) {
	var body AuthorDto
	body.Created_at = time.Now()

	if err := ctx.BindJSON(&body); err != nil {
		panic(err)
	}

	author := ac.Repository.CreateOne(ctx, body.ToModel())
	ctx.JSON(http.StatusCreated, author)
}

func (ac *AuthorsController) UpdateOne(ctx *gin.Context) {
	var id = ctx.Param("id")
	var body AuthorDto

	if err := ctx.BindJSON(&body); err != nil {
		panic(err)
	}

	updatedAuthor := ac.Repository.UpdateOne(ctx, id, body.ToModel())
	ctx.JSON(http.StatusAccepted, updatedAuthor)
}
