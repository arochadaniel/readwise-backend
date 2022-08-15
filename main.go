package main

import (
	"context"
	"readwise-backend/packages/core/repository"
	"readwise-backend/packages/core/routing"
	"readwise-backend/packages/domain/authors"
	"readwise-backend/packages/domain/books"
)

func main() {
	router := routing.SetupUpAppRouter()
	db := repository.SetUpMongoDatabaseContainer()
	defer db.DB.Client().Disconnect(context.Background())

	books.SetupBooksPackage(router, db)
	authors.SetupAuthorsPackage(router, db)

	router.Run(":8080")
}
