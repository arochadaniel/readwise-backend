package authors

import "readwise-backend/packages/core/repository"

type AuthorsRepository struct {
	*repository.MongoRepository[AuthorModel, AuthorDto]
}

func NewAuthorsRepository(db *repository.MongoDatabaseContainer) *AuthorsRepository {
	repo := AuthorsRepository{}
	repo.MongoRepository = &repository.MongoRepository[AuthorModel, AuthorDto]{CollectionName: "authors", DBContainer: db}
	return &repo
}
