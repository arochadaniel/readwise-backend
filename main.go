package main

import (
	"readwise-backend/packages/core/repository"
	"readwise-backend/packages/domain/authors"
	"readwise-backend/packages/domain/books"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := repository.SetUpMongoDatabaseContainer()

	var br = books.NewBookRepository(db)
	var bc = books.NewBookController(br)
	books.RegisterBookRoutes(router, bc)

	var ar = authors.NewAuthorsRepository(db)
	var ac = authors.NewAuthorsController(ar)
	authors.RegisterAuthorsRoutes(router, ac)

	router.Run(":8080")
}
