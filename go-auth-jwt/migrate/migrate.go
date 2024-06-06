package main

import (
	"gtihub.com/Jram-IR/go-auth-jwt/config"
	"gtihub.com/Jram-IR/go-auth-jwt/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()

}

func main() {
	config.DB.AutoMigrate(&models.User{})
}
