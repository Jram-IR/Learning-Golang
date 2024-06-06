package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gtihub.com/Jram-IR/go-auth-jwt/config"
	"gtihub.com/Jram-IR/go-auth-jwt/controllers"
	"gtihub.com/Jram-IR/go-auth-jwt/middleware"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()

}
func main() {
	router := gin.Default()
	router.POST("/auth/signup", controllers.RegisterUser)
	router.POST("/auth/login", controllers.LoginUser)
	router.GET("/auth/validate", middleware.CheckAuth, controllers.ValidateUser)

	router.Run(os.Getenv("PORT"))

}
