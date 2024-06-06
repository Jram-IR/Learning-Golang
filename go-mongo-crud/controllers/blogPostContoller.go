package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Jram-IR/go-mongo-crud/config"
	"github.com/Jram-IR/go-mongo-crud/helper"
	"github.com/Jram-IR/go-mongo-crud/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// get the collection reference

func AddBlog(c *gin.Context) {
	var post models.Post

	if err := c.BindJSON(&post); err != nil {
		helper.HandleError(c, err, http.StatusBadRequest, "Error add post !")
		return
	}

	_, err := config.BlogCol.InsertOne(context.Background(), post) // returns result and error
	if err != nil {
		helper.HandleError(c, err, http.StatusBadRequest, "Error when inserting in the DB")
		return
	}

	c.IndentedJSON(http.StatusOK, post)

}

func FindPost(c *gin.Context) {
	var post models.Post
	title := c.Param("title")
	// should be in bson and T of title should be small as mentioned in the struct declaration
	filter := bson.D{{Key: "title", Value: title}}
	err := config.BlogCol.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		helper.HandleError(c, err, http.StatusNotFound, "Post not found")
		return
	}
	c.IndentedJSON(http.StatusOK, post)

}

func UpdatePost(c *gin.Context) {
	var updatedPost models.Post
	author := c.Param("author")
	// should be in bson and T of title should be small as mentioned in the struct declaration
	filter := bson.D{{Key: "author", Value: author}}

	if err := c.BindJSON(&updatedPost); err != nil {
		helper.HandleError(c, err, 500, "unable to get body of json")
		return
	}

	//marshall the struct into bytes
	updateBsonBytes, err := bson.Marshal(updatedPost)
	if err != nil {
		helper.HandleError(c, err, 500, "error on marshalling bson")
		return
	}

	//unmarshall the bytes into the bson.D data
	var convertedBson bson.D
	e := bson.Unmarshal(updateBsonBytes, &convertedBson)
	if e != nil {
		log.Fatal(e)
	}

	// Define the update document with the modifications to apply
	update := bson.D{
		{Key: "$set", Value: convertedBson},
	}

	//update the document using the update One function
	res, err := config.BlogCol.UpdateOne(context.Background(), filter, update)
	if err != nil {
		helper.HandleError(c, err, 500, "cant update the post in db")
		return
	}

	m1 := res.MatchedCount
	m2 := res.ModifiedCount

	//respond
	c.IndentedJSON(http.StatusOK, gin.H{
		"matched":  m1,
		"modified": m2,
	})

}

func DeletePost(c *gin.Context) {

	title := c.Param("title")
	// should be in bson and T of title should be small as mentioned in the struct declaration
	filter := bson.D{{Key: "title", Value: title}}
	delRes, err := config.BlogCol.DeleteOne(context.Background(), filter)
	if err != nil {
		helper.HandleError(c, err, http.StatusNotFound, "Post not found")
		return
	}
	if delRes.DeletedCount == 0 {
		helper.HandleError(c, err, http.StatusNotFound, "Couldn't delete post")
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"deleted count": delRes.DeletedCount,
	})

}
