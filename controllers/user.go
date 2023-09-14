package controllers

import (
	"Lara/helpers"
	"Lara/models"
	"Lara/models/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Register(c *gin.Context) {
	var user users.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//Verifica se o email ou username já está cadastrado
	var existingUser users.User
	err = models.Db.Where("email = ? OR username = ?", user.Email, user.Username).First(&existingUser).Error
	if err == nil {
		c.JSON(400, gin.H{
			"msg": "Email or username already registered",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)

	err = models.Db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Error creating user",
		})
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) Login(c *gin.Context) {
	var user users.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	var existingUser users.User
	err = models.Db.Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Email not registered",
		})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid credentials"})
		return
	}

	token, err := helpers.GenerateToken(existingUser.ID)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) ReadUser(c *gin.Context) {
	var user users.User
	user_id := c.Param("user_id")

	models.Db.Preload("Reviews").Find(&user, user_id)
	if user.ID == 0 {
		c.JSON(400, gin.H{
			"msg": "User not found",
		})
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) ReadUsers(c *gin.Context) {
	var users []users.User

	err := models.Db.Find(&users).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user users.User
	user_id := c.Param("user_id")

	err := models.Db.Find(&user, user_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "User not found",
		})
		return
	}

	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	models.Db.Save(&user)
	c.JSON(200, &user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	var user users.User
	user_id := c.Param("user_id")

	if err := models.Db.Find(&user, user_id).Error; err != nil {
		c.JSON(400, gin.H{
			"msg": "User not found",
		})
		return
	}

	models.Db.Delete(&user)
	c.JSON(200, &user)
}

func (uc *UserController) ReadUserReviews(c *gin.Context) {
	var user users.User
	user_id := c.Param("user_id")

	err := models.Db.Preload("Reviews").Find(&user, user_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "User not found",
		})
		return
	}

	c.JSON(200, user.Reviews)
}

func (uc *UserController) ReadUserStats(c *gin.Context) {
	var user users.User
	user_id := c.Param("user_id")

	err := models.Db.Preload("Reviews").Find(&user, user_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, user.GetStats())
}

func (uc *UserController) ReadUserLikes(c *gin.Context) {
	var user users.User
	user_id := c.Param("user_id")

	err := models.Db.Preload("Likes").Find(&user, user_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "User not found",
		})
		return
	}

	c.JSON(200, user.Likes)
}

func (uc *UserController) CreateContentInteraction(c *gin.Context) {
	content_type := c.Param("content_type")
	content_id := c.Param("content_id")

	user_id := c.MustGet("user_id").(uint)
	interaction_type := c.Param("interaction_type")

	content, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err = models.Db.First(content, content_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Content not found"})
		return
	}

	//Verifica se o usuário já interagiu com o conteúdo
	var existingInteraction users.UserReviewable
	if err = models.Db.Where("user_id = ? AND reviewable_id = ? AND reviewable_type = ?", user_id, content_id, content_type).First(&existingInteraction).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "User already interacted with this content"})
		return
	}

	switch interaction_type {
	case "planto":
		if err = models.Db.Model(&users.User{}).Association("PlanTo").Append(&users.UserReviewable{
			UserID:         user_id,
			ReviewableID:   content.GetID(),
			ReviewableType: content.GetType(),
			PlanTo:         true,
		}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
	case "current":
		if err = models.Db.Model(&users.User{}).Association("Current").Append(&users.UserReviewable{
			UserID:         user_id,
			ReviewableID:   content.GetID(),
			ReviewableType: content.GetType(),
			Current:        true,
		}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
	case "finished":
		if err = models.Db.Model(&users.User{}).Association("Finished").Append(&users.UserReviewable{
			UserID:         user_id,
			ReviewableID:   content.GetID(),
			ReviewableType: content.GetType(),
			Finished:       true,
		}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid interaction"})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Interaction added successfully"})
}

func (uc *UserController) ReadContentInteraction(c *gin.Context) {
	user_id := c.MustGet("user_id").(uint)
	interaction_type := c.Param("interaction_type")

	var user users.User
	err := models.Db.Preload("Reviews").Preload("PlanTo").Preload("Current").Preload("Finished").Find(&user, user_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	switch interaction_type {
	case "planto":
		c.JSON(http.StatusOK, user.PlanTo)
	case "current":
		c.JSON(http.StatusOK, user.Current)
	case "finished":
		c.JSON(http.StatusOK, user.Finished)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid interaction"})
	}
}
