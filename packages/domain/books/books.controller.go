package books

import (
	"net/http"
	"readwise-backend/packages/core/controller"
	"readwise-backend/packages/core/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func NewBookController() *BookController {
	BookRepository := NewBookRepository()
	controller := BookController{}
	controller.Repository = BookRepository
	return &controller
}

type BookController struct {
	controller.GinController[BookModel, BookDto]
}

func (c *BookController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	objectID := repository.GetPrimitiveObjectIDFromString(id)
	book := c.Repository.FindOne(ctx, BookModel{ID: objectID})
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) FindAll(ctx *gin.Context) {
	var body BookDto
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	filter := body.ToModel()
	books := c.Repository.FindAll(ctx, filter)

	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) CreateOne(ctx *gin.Context) {
	var book BookDto
	book.Created_at = time.Now()

	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	model := book.ToModel()
	id := c.Repository.CreateOne(ctx, model)
	book.ID = id
	ctx.JSON(http.StatusCreated, book)
}

func (c *BookController) CreateMultiple(ctx *gin.Context) {
	var booksBody []BookDto

	if err := ctx.BindJSON(&booksBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	for i := range booksBody {
		booksBody[i].Created_at = time.Now()
	}

	booksModels := repository.MapDtosToModels[BookModel](booksBody)

	ids := c.Repository.CreateMultiple(ctx, booksModels)
	for i := range booksBody {
		booksBody[i].ID = ids[i]
	}

	ctx.JSON(http.StatusCreated, booksBody)
}

func (c *BookController) UpdateOne(ctx *gin.Context) {
	var book_id = ctx.Param("id")
	var book BookDto

	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	id := c.Repository.UpdateOne(ctx, book_id, book.ToModel())
	book.ID = id
	ctx.JSON(http.StatusPartialContent, book)
}

func (c *BookController) UpdateBy(ctx *gin.Context) {
	var body struct {
		Filter   BookDto `json:"filter"`
		ToUpdate BookDto `json:"toUpdate"`
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	count := c.Repository.UpdateBy(ctx, body.Filter.ToModel(), body.ToUpdate.ToModel())
	ctx.JSON(http.StatusAccepted, gin.H{"modifiedCount": count})
}

func (c *BookController) DeleteBy(ctx *gin.Context) {
	var filter BookDto
	if err := ctx.BindJSON(&filter); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	deletedBooksCount := c.Repository.DeleteBy(ctx, filter.ToModel())
	ctx.JSON(http.StatusAccepted, gin.H{"deletedBooksCount": deletedBooksCount})
}

func (c *BookController) DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")
	primitiveID := repository.GetPrimitiveObjectIDFromString(id)
	deletedBooksCount := c.Repository.DeleteOne(ctx, BookModel{ID: primitiveID})
	ctx.JSON(http.StatusAccepted, gin.H{"deletedBookID": id, "deletedBooksCount": deletedBooksCount})
}
