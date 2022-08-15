package books

import (
	"net/http"
	"readwise-backend/packages/core/controller"
	"readwise-backend/packages/core/repository"
	"time"

	"github.com/gin-gonic/gin"
)

func NewBookController(r *BookRepository) *BookController {
	controller := BookController{}
	controller.Repository = r
	return &controller
}

type BookController struct {
	controller.GinController[BookModel, BookDto, BookRepository]
}

func (c *BookController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	objectID := repository.GetPrimitiveObjectIDFromString(id)
	book, err := c.Repository.FindOne(ctx, BookModel{ID: objectID})

	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) FindAll(ctx *gin.Context) {
	var body BookDto
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	filter := body.ToModel()
	books, err := c.Repository.FindAll(ctx, filter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) CreateOne(ctx *gin.Context) {
	var book BookDto
	book.Created_at = time.Now()

	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	model := book.ToModel()
	response := c.Repository.InsertBookAndAuthor(ctx, model)
	ctx.JSON(http.StatusCreated, response)
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
	response := c.Repository.CreateMultiple(ctx, booksModels)
	ctx.JSON(http.StatusCreated, response)
}

func (c *BookController) UpdateOne(ctx *gin.Context) {
	var book_id = ctx.Param("id")
	var book BookDto

	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	response := c.Repository.UpdateOne(ctx, book_id, book.ToModel())
	ctx.JSON(http.StatusPartialContent, response)
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
