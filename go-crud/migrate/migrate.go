package main

import (
	"fmt"
	"go-crud/config"
	"go-crud/models"
	"log"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	err := config.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Unable to Auto Migrate")
	}
	fmt.Println("Auto Migrated!")
}
