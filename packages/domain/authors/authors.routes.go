package authors

import "github.com/gin-gonic/gin"

var ar = NewAuthorsRepository()
var ac = NewAuthorsController(ar)

func RegisterAuthorsRoutes(router *gin.Engine) {
	group := router.Group("/authors")
	group.POST("/", ac.CreateOne)
	group.PATCH("/:id", ac.UpdateOne)
}
