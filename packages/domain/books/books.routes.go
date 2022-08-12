package books

import (
	"github.com/gin-gonic/gin"
)

var br = NewBookRepository()
var bc = NewBookController(br)

func RegisterBookRoutes(router *gin.Engine) {
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
