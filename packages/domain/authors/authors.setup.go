package authors

import (
	"readwise-backend/packages/core/repository"

	"github.com/gin-gonic/gin"
)

func SetupAuthorsPackage(router *gin.Engine, db *repository.MongoDatabaseContainer) {
	var ar = NewAuthorsRepository(db)
	var ac = NewAuthorsController(ar)
	RegisterAuthorsRoutes(router, ac)
}
