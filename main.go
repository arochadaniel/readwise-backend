package main

import (
	"readwise-backend/packages/domain/authors"
	"readwise-backend/packages/domain/books"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	books.RegisterBookRoutes(router)
	authors.RegisterAuthorsRoutes(router)
	router.Run(":8080")
}
