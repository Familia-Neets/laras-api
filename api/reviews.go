package api

import (
	"Lara/controllers"

	"github.com/gin-gonic/gin"
)

func reviews(r *gin.Engine) {
	reviewController := controllers.NewReviewController()

	r.POST("/:content_type/:content_id/reviews", controllers.AuthMiddleware, reviewController.CreateReview)
	r.GET("/reviews/:review_id", reviewController.ReadReview)
	r.PUT("/reviews/:review_id", controllers.AuthMiddleware, reviewController.UpdateReview)
	r.DELETE("/reviews/:review_id", controllers.AuthMiddleware, reviewController.DeleteReview)

	r.POST("/reviews/:review_id/like", controllers.AuthMiddleware, reviewController.LikeReview)
	r.DELETE("/reviews/:review_id/like", controllers.AuthMiddleware, reviewController.UnlikeReview)
}
