package controllers

import (
	"Lara/helpers"
	"Lara/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewController struct{}

func NewReviewController() *ReviewController {
	return &ReviewController{}
}

func (rc *ReviewController) CreateReview(c *gin.Context) {
	stringId := c.Param("content_id")
	content_type := c.Param("content_type")

	content_id, err := strconv.Atoi(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number"})
		return
	}

	// Criando review
	var review models.Review
	if err := c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	review.ReviewableID = uint(content_id)
	review.UserID = c.MustGet("user_id").(uint)

	// Definindo o tipo de review
	content, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	// Definindo o tipo de review (filme, série, livro, etc)
	review.ReviewableType = content.GetType()

	// Verificando se o conteúdo existe
	if err := models.Db.Find(content, review.ReviewableID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content not found"})
		return
	}

	// Verificando se o usuário existe
	var user models.User
	models.Db.Find(&user, review.UserID)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Verificando se o usuário já avaliou o conteúdo
	var existingReview models.Review
	models.Db.Where("user_id = ? AND reviewable_id = ? AND reviewable_type = ?", review.UserID, review.ReviewableID, review.ReviewableType).First(&existingReview)
	if existingReview.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Review already exists"})
		return
	}

	if err := models.Db.Create(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.Db.Model(content).Association("Reviews").Append(&review)
	models.Db.Model(&user).Association("Reviews").Append(&review)
	c.JSON(http.StatusOK, review)
}

func (rc *ReviewController) ReadReview(c *gin.Context) {
	id := c.Param("review_id")

	var review models.Review
	if err := models.Db.Preload("Likes").Find(&review, id).Error; err != nil {
		c.JSON(400, gin.H{"message": "Review not found"})
		return
	}

	c.JSON(200, review)
}

func (rc *ReviewController) UpdateReview(c *gin.Context) {
	user_id := c.MustGet("user_id").(uint)

	id := c.Param("review_id")

	var review models.Review
	if err := models.Db.Find(&review, id).Error; err != nil {
		c.JSON(400, gin.H{"message": "Review not found"})
		return
	}

	if review.UserID != user_id {
		c.JSON(400, gin.H{"message": "You can't update this review"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	if err := models.Db.Save(&review).Error; err != nil {
		c.JSON(400, gin.H{"message": "Failed to update review"})
		return
	}

	c.JSON(200, gin.H{"message": "Review updated successfully"})
}

func (rc *ReviewController) DeleteReview(c *gin.Context) {
	user_id := c.MustGet("user_id").(uint)
	review_id := c.Param("review_id")

	var review models.Review
	if err := models.Db.Where("id = ?", review_id).First(&review).Error; err != nil {
		c.JSON(400, gin.H{"message": "Review not found"})
		return
	}

	if review.UserID != user_id {
		c.JSON(400, gin.H{"message": "You can't delete this review"})
		return
	}

	if err := models.Db.Delete(&review).Error; err != nil {
		c.JSON(400, gin.H{"message": "Failed to delete review"})
		return
	}

	c.JSON(200, gin.H{"message": "Review deleted successfully"})
}

func (rc *ReviewController) LikeReview(c *gin.Context) {
	review_id := c.Param("review_id")
	user_id := c.MustGet("user_id")

	var review models.Review
	if err := models.Db.Find(&review, review_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Review not found"})
		return
	}

	var like models.Like
	if err := models.Db.Where("user_id = ? AND review_id = ?", user_id, review_id).First(&like).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You already liked this review"})
		return
	}

	like = models.Like{
		UserID:   user_id.(uint),
		ReviewID: review.ID,
	}

	if err := models.Db.Create(&like).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while liking review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review liked successfully"})
}

func (rc *ReviewController) UnlikeReview(c *gin.Context) {
	review_id := c.Param("review_id")
	user_id := c.MustGet("user_id")

	var review models.Review
	if err := models.Db.Find(&review, review_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Review not found"})
		return
	}

	var like models.Like
	if err := models.Db.Where("user_id = ? AND review_id = ?", user_id, review_id).First(&like).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You didn't like this review"})
		return
	}

	if err := models.Db.Delete(&like).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while unliking review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review unliked successfully"})
}
