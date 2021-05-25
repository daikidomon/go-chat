package main

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI("mongodb://root:password@mongodb:27017"))
	println(err)
}

type Book struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func NewBook(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}

func main() {
	book := NewBook("Pride and Prejudice", 345)

	// Make sure to pass the model by reference.
	err := mgm.Coll(book).Create(book)

	println(err)
	print("Test")
}
