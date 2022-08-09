package books

import r "readwise-backend/packages/core/repository"

func NewBookRepository() *r.MongoRepository[BookModel, BookDto] {
	return &r.MongoRepository[BookModel, BookDto]{CollectionName: "books"}
}
