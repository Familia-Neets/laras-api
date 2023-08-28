package api

import (
	"Lara/controllers"

	"github.com/gin-gonic/gin"
)

func content(r *gin.Engine) {
	contentController := controllers.NewContentController()

	r.POST("/:content_type", contentController.CreateContent)
	r.GET("/:content_type", contentController.ReadContents)
	r.GET("/:content_type/:content_id", contentController.ReadContent)
	r.PUT("/:content_type/:content_id", contentController.UpdateContent)
	r.DELETE("/:content_type/:content_id", contentController.DeleteContent)

	r.GET("/:content_type/:content_id/reviews", contentController.ReadReviews)
}
