package controllers

import (
	"Lara/helpers"
	"Lara/models"
	"Lara/models/reviewable"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ContentController struct{}

func NewContentController() *ContentController {
	return &ContentController{}
}

func (cc *ContentController) CreateContent(c *gin.Context) {
	content_type := c.Param("content_type")

	content, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = c.BindJSON(&content)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = helpers.ValidateRequiredFields(content)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.Db.Create(content).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, content)
}

func (cc *ContentController) ReadContents(c *gin.Context) {
	content_type := c.Param("content_type")

	contentInstance, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	contentSlice := reflect.New(reflect.SliceOf(reflect.TypeOf(contentInstance).Elem())).Interface()

	err = models.Db.Find(contentSlice).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, contentSlice)
}

func (cc *ContentController) ReadContent(c *gin.Context) {
	content_type := c.Param("content_type")
	content_id := c.Param("content_id")

	contentInstance, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.Db.Preload("Reviews").First(contentInstance, content_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, contentInstance)
}

func (cc *ContentController) UpdateContent(c *gin.Context) {
	content_type := c.Param("content_type")
	content_id := c.Param("content_id")

	contentInstance, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.Db.First(contentInstance, content_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = c.BindJSON(&contentInstance)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = helpers.ValidateRequiredFields(contentInstance)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.Db.Save(contentInstance).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, contentInstance)
}

func (cc *ContentController) DeleteContent(c *gin.Context) {
	content_type := c.Param("content_type")
	content_id := c.Param("content_id")

	contentInstance, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.Db.First(contentInstance, content_id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.Db.Delete(contentInstance).Error
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "Content deleted successfully",
	})
}

func (cc *ContentController) ReadReviews(c *gin.Context) {
	content_type := c.Param("content_type")
	content_id := c.Param("content_id")

	content, err := helpers.GetContentInstance(content_type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reviews []reviewable.Review
	if err := models.Db.Where("reviewable_id = ? AND reviewable_type = ?", content_id, content.GetType()).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
