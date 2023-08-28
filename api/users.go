package api

import (
	"Lara/controllers"

	"github.com/gin-gonic/gin"
)

func users(r *gin.Engine) {
	userController := controllers.NewUserController()

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.GET("/users", userController.ReadUsers)
	r.GET("/users/:user_id", userController.ReadUser)
	r.PUT("/users/:user_id", userController.UpdateUser)
	r.DELETE("/users/:user_id", userController.DeleteUser)

	r.GET("/users/:user_id/stats", userController.ReadUserStats)
	r.GET("/users/:user_id/reviews", userController.ReadUserReviews)
	r.GET("/users/:user_id/likes", userController.ReadUserLikes)

	r.POST("/:content_type/:content_id/:interaction_type", userController.CreateContentInteraction)
	r.GET("/users/:user_id/:interaction_type", userController.ReadContentInteraction)
}
