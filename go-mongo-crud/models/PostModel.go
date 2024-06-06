package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Body   string             `bson:"body,omitempty"`
	Author string             `bson:"author,omitempty"`
}
