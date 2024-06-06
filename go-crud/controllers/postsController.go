package controllers

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var post models.Post
	err := c.BindJSON(&post)
	if err != nil {
		c.String(http.StatusInternalServerError, "some thing went wrong :(")
		return
	}
	result := config.DB.Create(&post)
	if result.Error != nil {
		c.String(http.StatusBadRequest, "errror in creating the post")
		return
	}

	c.IndentedJSON(http.StatusOK, post)

}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	result := config.DB.Find(&posts)

	if result.Error != nil {
		c.String(http.StatusBadRequest, "error fetching the data")
		return

	}

	c.IndentedJSON(http.StatusOK, posts)

}

func GetPostByID(c *gin.Context) {
	var post models.Post
	postId := c.Param("id")

	result := config.DB.First(&post, postId)
	if result.Error != nil {
		c.String(http.StatusNotFound, "the post was not found ")
		return
	}
	c.IndentedJSON(http.StatusFound, post)

}

func UpdatePost(c *gin.Context) {
	var post models.Post
	err := c.BindJSON(&post)
	if err != nil {
		c.String(http.StatusBadRequest, "an error occurred")
		return
	}
	//update the post if already there else create it
	result := config.DB.Save(&post)
	if result.Error != nil {
		c.String(http.StatusBadRequest, "the post was not upadated")
		return
	}
	c.IndentedJSON(http.StatusOK, post)

}

func DeletePost(c *gin.Context) {
	postId := c.Param("id")
	var post models.Post
	result := config.DB.Delete(&post, postId)
	if result.Error != nil {
		c.String(http.StatusBadRequest, "the post could not be deleted")
		return
	}
	c.String(http.StatusGone, "the post was deleted")

}
