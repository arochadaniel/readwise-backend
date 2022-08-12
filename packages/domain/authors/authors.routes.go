package authors

import "github.com/gin-gonic/gin"

func RegisterAuthorsRoutes(router *gin.Engine, ac *AuthorsController) {
	group := router.Group("/authors")
	group.POST("/", ac.CreateOne)
	group.PATCH("/:id", ac.UpdateOne)
}
