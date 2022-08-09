package books

import "readwise-backend/packages/core/routing"

var bookController = NewBookController()

func RegisterBookRoutes() {
	var group = routing.Router.Group("/books")
	group.GET("", bookController.FindAll)
	group.GET("/:id", bookController.FindOne)
	group.POST("", bookController.CreateOne)
	group.POST("/multiple", bookController.CreateMultiple)
	group.PATCH("/:id", bookController.UpdateOne)
	group.PATCH("", bookController.UpdateBy)
	group.DELETE("/:id", bookController.DeleteOne)
	group.DELETE("", bookController.DeleteBy)
}
