package main

import (
	"go-crud/config"
	"go-crud/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	//load the env varialbles first when initializing
	config.LoadEnvVariables()
	// Connect to the postgres DB
	config.ConnectToDB()

}
func main() {
	router := gin.Default()
	//get all the posts
	router.GET("blog", controllers.GetAllPosts)

	//get the specific post by id
	router.GET("blog/:id", controllers.GetPostByID)

	//create the post
	router.POST("blog", controllers.CreatePost)

	//update the  post
	router.POST("blog/update", controllers.UpdatePost)

	//delete the post
	router.DELETE("blog/:id", controllers.DeletePost)

	router.Run(os.Getenv("PORT"))

}
