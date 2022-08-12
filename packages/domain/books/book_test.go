package books_test

import (
	"readwise-backend/packages/core/repository"
	"readwise-backend/packages/domain/books"
	"testing"

	"github.com/stretchr/testify/require"
)

var strID string = "62ed605431d6cfc0b47a7f60"

func TestShouldSetBookDtoID(t *testing.T) {
	b := books.BookDto{}
	require.Equal(t, "", b.ID)

	bWithID := b.SetID(repository.GetPrimitiveObjectIDFromString(strID))
	require.Equal(t, strID, bWithID.ID)
}

func TestShouldMapBookDtoToModel(t *testing.T) {
	b := books.BookDto{ID: strID, Title: "Mock title", Description: "Mock description"}
	require.Implements(t, (*repository.RepositoryDto[books.BookModel, books.BookDto])(nil), b)
	mapped := b.ToModel()
	require.Implements(t, (*repository.RepositoryModel[books.BookDto])(nil), mapped)
	require.Equal(t, repository.GetPrimitiveObjectIDFromString(strID), mapped.ID)
}

func TestShouldMapBookModelToDto(t *testing.T) {
	b := books.BookModel{ID: repository.GetPrimitiveObjectIDFromString(strID), Title: "Mock title", Description: "Mock description"}
	require.Implements(t, (*repository.RepositoryModel[books.BookDto])(nil), b)
	mapped := b.ToEntity()
	require.Implements(t, (*repository.RepositoryDto[books.BookModel, books.BookDto])(nil), mapped)
	require.Equal(t, strID, mapped.ID)
}
