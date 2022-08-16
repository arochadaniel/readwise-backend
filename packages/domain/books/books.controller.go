package books

import (
	"net/http"
	"readwise-backend/packages/core/controller"

	"github.com/gin-gonic/gin"
)

func NewBookController(r BookRepository) *BookController {
	c := BookController{}
	c.Repository = r
	return &c
}

type BookController struct {
	controller.GinController[BookModel, BookDto, BookRepository]
}

func (bc *BookController) InsertBookAndAuthor(ctx *gin.Context) {
	var dto BookDto

	if err := ctx.BindJSON(&dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	model := dto.Init().ToModel()
	response := bc.Repository.InsertBookAndAuthor(ctx, model)
	ctx.JSON(http.StatusCreated, response)
}
