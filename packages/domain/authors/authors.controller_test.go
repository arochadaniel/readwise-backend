package authors_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"readwise-backend/packages/core/apptesting"
	"readwise-backend/packages/core/repository"
	"readwise-backend/packages/core/routing"
	"readwise-backend/packages/domain/authors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var db *repository.MongoDatabaseContainer

func TestMain(m *testing.M) {
	dbContainer, ctx := apptesting.SetupTestingDB()
	db = dbContainer
	defer db.DB.Client().Disconnect(ctx)

	result := m.Run()

	os.Exit(result)
}

func TestAuthorsCreateEndpoint(t *testing.T) {
	router := routing.SetupUpAppRouter()
	authors.SetupAuthorsPackage(router, db)
	mockAuthor := authors.AuthorDto{Name: "MockName", Description: "MockDescription", Created_at: time.Now()}
	mockAuthorJson, _ := json.Marshal(mockAuthor)

	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(mockAuthorJson))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var resBody authors.AuthorDto
	json.Unmarshal(w.Body.Bytes(), &resBody)
	require.Equal(t, http.StatusCreated, w.Code)
}

func TestAuthorsUpdateEndpoint(t *testing.T) {
	router := routing.SetupUpAppRouter()
	authors.SetupAuthorsPackage(router, db)
	mockAuthor := authors.AuthorDto{Name: "MockName", Description: "MockDescription", Created_at: time.Now()}
	mockAuthorJson, _ := json.Marshal(mockAuthor)

	postRecorder := httptest.NewRecorder()
	reqPost, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(mockAuthorJson))
	router.ServeHTTP(postRecorder, reqPost)

	var postResBody authors.AuthorDto
	json.Unmarshal(postRecorder.Body.Bytes(), &postResBody)
	require.Equal(t, http.StatusCreated, postRecorder.Code)

	mockAuthorUpdate := authors.AuthorDto{Name: "MockNameUPDATED"}
	mockAuthorJsonUpdate, _ := json.Marshal(mockAuthorUpdate)
	patch, _ := http.NewRequest("PATCH", fmt.Sprintf("/authors/%s", postResBody.ID), bytes.NewBuffer(mockAuthorJsonUpdate))
	patchRecorder := httptest.NewRecorder()
	router.ServeHTTP(patchRecorder, patch)

	var patchResBody authors.AuthorDto
	json.Unmarshal(patchRecorder.Body.Bytes(), &patchResBody)
	require.Equal(t, patchResBody.ID, postResBody.ID)
	require.Equal(t, patchResBody.Name, mockAuthorUpdate.Name)
}
