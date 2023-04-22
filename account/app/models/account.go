package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	Id        string `json:"id,omitempty", ommitempty`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"last_name"`
	LastName  string `json:"first_name"`
}

func (a *Account) Create(coll *mongo.Collection) (*mongo.InsertOneResult, error) {
	return coll.InsertOne(context.TODO(), &a)
}

func (a *Account) Get(coll *mongo.Collection) error {
	filter := bson.D{{"id", a.Id}}
	var result Account
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	a = &result
	return err
}

func (a *Account) Update(coll *mongo.Collection) error {
	filter := bson.D{{"id", a.Id}}
	update := bson.D{{"$set", bson.D{{"avg_rating", 4.4}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	log.Printf("Updated account: %v", result)
	return err
}
