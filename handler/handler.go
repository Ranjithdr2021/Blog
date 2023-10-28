package handler

import (
	"blog/database"
	"blog/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostBlog(c *gin.Context) {
	db := database.NewDb().ConnectToDB()
	db.AutoMigrate(&models.Post{})
	post := models.Post{}
	json.NewDecoder(c.Request.Body).Decode(&post)
	pk := db.Create(post)
	if pk == nil {
		c.JSON(http.StatusInternalServerError, "unable to insert data")
	}
	c.JSON(http.StatusOK, "new blog post created")
}

func GetBlogPosts(c *gin.Context) {
	// Create a new record
	db := database.NewDb().ConnectToDB()
	db.AutoMigrate(&models.Post{})
	post := []models.Post{}
	db.Find(&post)
	if len(post) == 0 {
		c.JSON(http.StatusOK, "no blog posts available")
		return
	}
	c.JSON(http.StatusOK, post)
}

// GetBlogPost retrieves the post with the given id
func GetBlogPost(c *gin.Context) {
	id := c.Param("id")
	db := database.NewDb().ConnectToDB()
	db.AutoMigrate(&models.Post{})
	post := models.Post{}
	result := db.First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, "No blog found with given id")
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdateBlogPost updates the blog post with the given id
func UpdateBlogPost(c *gin.Context) {
	id := c.Param("id")
	db := database.NewDb().ConnectToDB()
	db.AutoMigrate(&models.Post{})
	post := models.Post{}
	json.NewDecoder(c.Request.Body).Decode(&post)
	result := db.Model(&models.Post{}).Where("id = ?", id).Updates(post)
	if result.Error != nil {
		fmt.Println("Error updating record:", result.Error)
	}
	c.JSON(http.StatusOK, "blog post updated successfully")
}

// DeleteBlogPost deletes the blog post with the given id
func DeleteBlogPost(c *gin.Context) {
	id := c.Param("id")
	db := database.NewDb().ConnectToDB()
	db.AutoMigrate(&models.Post{})
	post := []models.Post{}
	db.Delete(&post, id)
	c.JSON(http.StatusOK, "blog post deleted successfully")
}
