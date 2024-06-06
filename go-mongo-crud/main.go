package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jram-IR/go-mongo-crud/config"
	"github.com/Jram-IR/go-mongo-crud/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}
func main() {
	fmt.Println("Testing mongo db connection")

	router := gin.Default()
	router.GET("/blogger", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello from the other side")
	})

	//creat a blog post
	router.POST("/blogger", controllers.AddBlog)

	//find post by title
	router.GET("/blogger/:title", controllers.FindPost)

	//update the post by finding it by title
	router.POST("/blogger/:author", controllers.UpdatePost)

	//delete the post specified by the title
	router.DELETE("/blogger/:title", controllers.DeletePost)

	// just for testing authorization using the decorator pattern
	router.GET("/blogger/admin", gin.HandlerFunc(onlyAdmin(handleAdminStuff)))

	router.Run(os.Getenv("PORT"))

}

// all are dummy code for authorization concept testing and decorator pattern learning
type usr struct {
	role string
}

func getUserFromDb() usr {
	return usr{
		role: "user",
	}
}

func handleAdminStuff(c *gin.Context) {
	//handle the admin stuff
	c.String(http.StatusOK, "This is where you can handle the adimin stuff since you are an admin!")

}

// define the function signature
type handler func(c *gin.Context)

// decorator pattern any handle can be verified for authorization
// onlyAdmin is the decorator function
func onlyAdmin(fn handler) handler {
	return func(c *gin.Context) {
		if u := getUserFromDb(); u.role != "admin" {
			c.String(http.StatusUnauthorized, "Unauthorized! Only Admins Allowed Here")
		} else {
			//execute the handler that was passed
			fn(c)
		}
	}
}
