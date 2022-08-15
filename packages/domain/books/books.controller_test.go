package books_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"readwise-backend/packages/core/repository"
	"readwise-backend/packages/core/routing"
	"readwise-backend/packages/domain/books"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *repository.MongoDatabaseContainer

func TestMain(m *testing.M) {
	ctx := context.TODO()
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic("error instantiating mongo client")
	}

	if err := dbClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	db = &repository.MongoDatabaseContainer{}
	db.DatabaseContainer.DB = dbClient.Database("readwise")

	defer dbClient.Disconnect(ctx)

	result := m.Run()

	os.Exit(result)
}

func TestBooksCreateEndpoint(t *testing.T) {
	router := routing.SetupUpAppRouter()
	books.SetupBooksPackage(router, db)
	mockBook := books.BookDto{Title: "MockTitle", Description: "MockDescription", Created_at: time.Now()}
	mockBookJson, _ := json.Marshal(mockBook)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(mockBookJson))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var resBody books.BookDto
	json.Unmarshal(w.Body.Bytes(), &resBody)
	require.Equal(t, http.StatusCreated, w.Code)
	require.Equal(t, mockBook.Title, resBody.Title)
}

func TestBooksGetEndpoint(t *testing.T) {
	router := routing.SetupUpAppRouter()
	books.SetupBooksPackage(router, db)
	mockBook := books.BookDto{Title: "MockTitle", Description: "MockDescription", Created_at: time.Now()}
	mockBookJson, _ := json.Marshal(mockBook)

	postRecorder := httptest.NewRecorder()
	reqPost, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(mockBookJson))
	router.ServeHTTP(postRecorder, reqPost)

	var postResBody books.BookDto
	json.Unmarshal(postRecorder.Body.Bytes(), &postResBody)
	require.Equal(t, http.StatusCreated, postRecorder.Code)

	reqGet, _ := http.NewRequest("GET", fmt.Sprintf("/books/%s", postResBody.ID), nil)
	getRecorder := httptest.NewRecorder()
	router.ServeHTTP(getRecorder, reqGet)

	var getResBody books.BookDto
	json.Unmarshal(getRecorder.Body.Bytes(), &getResBody)
	require.Equal(t, getResBody.ID, postResBody.ID)
}

func TestBooksGetAllBooksEndpoint(t *testing.T) {
	router := routing.SetupUpAppRouter()
	books.SetupBooksPackage(router, db)
	mockBook := books.BookDto{}
	mockBookJson, _ := json.Marshal(mockBook)

	reqGet, _ := http.NewRequest("GET", "/books", bytes.NewBuffer(mockBookJson))
	getRecorder := httptest.NewRecorder()
	router.ServeHTTP(getRecorder, reqGet)

	var getResBody []books.BookDto
	json.Unmarshal(getRecorder.Body.Bytes(), &getResBody)
	require.GreaterOrEqual(t, len(getResBody), 1)
}

func TestBooksUpdateEndpoint(t *testing.T) {
	router := routing.SetupUpAppRouter()
	books.SetupBooksPackage(router, db)
	mockBook := books.BookDto{Title: "MockTitle", Description: "MockDescription", Created_at: time.Now()}
	mockBookJson, _ := json.Marshal(mockBook)

	postRecorder := httptest.NewRecorder()
	reqPost, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(mockBookJson))
	router.ServeHTTP(postRecorder, reqPost)

	var postResBody books.BookDto
	json.Unmarshal(postRecorder.Body.Bytes(), &postResBody)
	require.Equal(t, http.StatusCreated, postRecorder.Code)

	mockBookUpdate := books.BookDto{Title: "MockTitleUPDATED"}
	mockBookJsonUpdate, _ := json.Marshal(mockBookUpdate)
	patch, _ := http.NewRequest("PATCH", fmt.Sprintf("/books/%s", postResBody.ID), bytes.NewBuffer(mockBookJsonUpdate))
	patchRecorder := httptest.NewRecorder()
	router.ServeHTTP(patchRecorder, patch)

	var patchResBody books.BookDto
	json.Unmarshal(patchRecorder.Body.Bytes(), &patchResBody)
	require.Equal(t, patchResBody.ID, postResBody.ID)
	require.Equal(t, patchResBody.Title, mockBookUpdate.Title)
}

func TestBooksDeleteEndpoint(t *testing.T) {
	router := routing.SetupUpAppRouter()
	books.SetupBooksPackage(router, db)
	mockBook := books.BookDto{Title: "MockTitle", Description: "MockDescription", Created_at: time.Now()}
	mockBookJson, _ := json.Marshal(mockBook)

	postRecorder := httptest.NewRecorder()
	reqPost, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(mockBookJson))
	router.ServeHTTP(postRecorder, reqPost)

	var postResBody books.BookDto
	json.Unmarshal(postRecorder.Body.Bytes(), &postResBody)
	require.Equal(t, http.StatusCreated, postRecorder.Code)

	deleteRecorder := httptest.NewRecorder()
	reqDelete, _ := http.NewRequest("DELETE", fmt.Sprintf("/books/%s", postResBody.ID), nil)
	router.ServeHTTP(deleteRecorder, reqDelete)

	var deleteResBody books.BookDto
	json.Unmarshal(deleteRecorder.Body.Bytes(), &deleteResBody)
	require.Equal(t, http.StatusAccepted, deleteRecorder.Code)

	reqGet, _ := http.NewRequest("GET", fmt.Sprintf("/books/%s", postResBody.ID), nil)
	getRecorder := httptest.NewRecorder()
	router.ServeHTTP(getRecorder, reqGet)

	var getResBody books.BookDto
	json.Unmarshal(getRecorder.Body.Bytes(), &getResBody)
	require.Equal(t, getRecorder.Code, http.StatusNotFound)
}
