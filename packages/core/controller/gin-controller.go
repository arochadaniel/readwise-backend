package controller

import (
	"net/http"
	"readwise-backend/packages/core/repository"

	"github.com/gin-gonic/gin"
)

type GinController[Model repository.RepositoryModel[DTO], DTO repository.RepositoryDto[Model, DTO], R repository.Repository[Model, DTO]] struct {
	Repository R
	Controller[Model, DTO]
}

func (c *GinController[Model, DTO, R]) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	objectID := repository.GetPrimitiveObjectIDFromString(id)
	var dto DTO
	dto = dto.SetID(objectID)
	var model = dto.ToModel()
	book, err := c.Repository.FindOne(ctx, model)

	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *GinController[Model, DTO, R]) FindAll(ctx *gin.Context) {
	var body DTO
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

func (c *GinController[Model, DTO, R]) CreateOne(ctx *gin.Context) {
	var dto DTO

	if err := ctx.BindJSON(&dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	model := dto.Init().ToModel()
	response := c.Repository.CreateOne(ctx, model)
	ctx.JSON(http.StatusCreated, response)
}

func (c *GinController[Model, DTO, R]) CreateMultiple(ctx *gin.Context) {
	var body []DTO

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	for i := range body {
		body[i] = body[i].Init()
	}

	models := repository.MapDtosToModels[Model](body)
	response := c.Repository.CreateMultiple(ctx, models)
	ctx.JSON(http.StatusCreated, response)
}

func (c *GinController[Model, DTO, R]) UpdateOne(ctx *gin.Context) {
	var id = ctx.Param("id")
	var dto DTO

	if err := ctx.BindJSON(&dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	response := c.Repository.UpdateOne(ctx, id, dto.ToModel())
	ctx.JSON(http.StatusPartialContent, response)
}

func (c *GinController[Model, DTO, R]) UpdateBy(ctx *gin.Context) {
	var body struct {
		Filter   DTO `json:"filter"`
		ToUpdate DTO `json:"toUpdate"`
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	count := c.Repository.UpdateBy(ctx, body.Filter.ToModel(), body.ToUpdate.ToModel())
	ctx.JSON(http.StatusAccepted, gin.H{"modifiedCount": count})
}

func (c *GinController[Model, DTO, R]) DeleteBy(ctx *gin.Context) {
	var filter DTO
	if err := ctx.BindJSON(&filter); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	deletedBooksCount := c.Repository.DeleteBy(ctx, filter.ToModel())
	ctx.JSON(http.StatusAccepted, gin.H{"deletedBooksCount": deletedBooksCount})
}

func (c *GinController[Model, DTO, R]) DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")
	primitiveID := repository.GetPrimitiveObjectIDFromString(id)
	var dto DTO
	dto = dto.SetID(primitiveID)

	deletedBooksCount := c.Repository.DeleteOne(ctx, dto.ToModel())
	ctx.JSON(http.StatusAccepted, gin.H{"deletedBookID": id, "deletedBooksCount": deletedBooksCount})
}
