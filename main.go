package main

import (
	"readwise-backend/packages/core/routing"
	"readwise-backend/packages/domain/books"
)

func main() {
	books.RegisterBookRoutes()
	routing.Router.Run(":8080")
}
