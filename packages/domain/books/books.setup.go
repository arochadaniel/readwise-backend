package books

import (
	"readwise-backend/packages/core/repository"

	"github.com/gin-gonic/gin"
)

func SetupBooksPackage(router *gin.Engine, db *repository.MongoDatabaseContainer) {
	var br = NewBookRepository(db)
	var bc = NewBookController(br)
	RegisterBookRoutes(router, bc)
}
