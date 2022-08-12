package authors_test

import (
	"readwise-backend/packages/core/repository"
	"readwise-backend/packages/domain/authors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldSetAuthorDtoID(t *testing.T) {
	a := authors.AuthorDto{Name: "mockName", Description: "mockDescription"}
	require.Equal(t, "", a.ID)

	strID := "62ed605431d6cfc0b47a7f60"
	aWithID := a.SetID(repository.GetPrimitiveObjectIDFromString(strID))
	require.Equal(t, strID, aWithID.ID)
}

func TestAuthorDtoShouldMapToAuthorModel(t *testing.T) {
	strID := "62ed605431d6cfc0b47a7f60"
	a := authors.AuthorDto{ID: strID, Name: "mockName", Description: "mockDescription"}
	m := authors.AuthorModel{ID: repository.GetPrimitiveObjectIDFromString(strID), Name: "mockName", Description: "mockDescription"}

	mapped := a.ToModel()
	require.Equal(t, m, mapped)
}

func TestAuthorModelShouldMapToAuthorDto(t *testing.T) {
	strID := "62ed605431d6cfc0b47a7f60"
	a := authors.AuthorDto{ID: strID, Name: "mockName", Description: "mockDescription"}
	m := authors.AuthorModel{ID: repository.GetPrimitiveObjectIDFromString(strID), Name: "mockName", Description: "mockDescription"}

	mapped := m.ToEntity()
	require.Equal(t, a, mapped)
}
