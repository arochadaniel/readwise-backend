package books

import (
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine, bc *BookController) {
	var group = router.Group("/books")
	group.GET("", bc.FindAll)
	group.GET("/:id", bc.FindOne)
	group.POST("", bc.CreateOne)
	group.POST("/multiple", bc.CreateMultiple)
	group.PATCH("/:id", bc.UpdateOne)
	group.PATCH("", bc.UpdateBy)
	group.DELETE("/:id", bc.DeleteOne)
	group.DELETE("", bc.DeleteBy)
}
