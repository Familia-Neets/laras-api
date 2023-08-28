package controllers

import (
	"Lara/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "No token provided"})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
		c.Abort()
		return
	}

	userId := claims["userId"].(float64)

	c.Set("user_id", uint(userId))
	c.Next()
}
