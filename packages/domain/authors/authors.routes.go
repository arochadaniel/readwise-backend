package authors

import "github.com/gin-gonic/gin"

func RegisterAuthorsRoutes(router *gin.Engine, ac *AuthorsController) {
	group := router.Group("/authors")
	group.GET("", ac.FindAll)
	group.GET("/:id", ac.FindOne)
	group.POST("", ac.CreateOne)
	group.POST("/multiple", ac.CreateMultiple)
	group.PATCH("/:id", ac.UpdateOne)
	group.PATCH("", ac.UpdateBy)
	group.DELETE("/:id", ac.DeleteOne)
	group.DELETE("", ac.DeleteBy)
}
