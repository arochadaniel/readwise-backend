package books

import (
	"context"
	r "readwise-backend/packages/core/repository"
	"readwise-backend/packages/domain/authors"
	"time"
)

type BookRepository struct {
	*r.MongoRepository[BookModel, BookDto]
}

func NewBookRepository(db *r.MongoDatabaseContainer) *BookRepository {
	repo := BookRepository{}
	repo.DBContainer = db
	repo.MongoRepository = &r.MongoRepository[BookModel, BookDto]{CollectionName: "books"}
	return &repo
}

func (br *BookRepository) InsertBookAndAuthor(ctx context.Context, b BookModel) BookDto {
	var response BookDto
	authorsRepo := authors.NewAuthorsRepository(br.DBContainer)
	author := authors.AuthorModel{Name: b.Author.Name, Description: b.Author.Description, Created_at: time.Now()}
	authorsRepo.CreateOne(ctx, author)
	insertedBook := br.CreateOne(ctx, b)
	response = insertedBook

	return response
}
