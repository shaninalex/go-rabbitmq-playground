package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetCreateAccount struct {
	Sub       string `json:"id,omitempty"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"lastname" binding:"required"`
	LastName  string `json:"firstname" binding:"required"`
}

func (a *GetCreateAccount) Create(coll *mongo.Collection) (*mongo.InsertOneResult, error) {
	fmt.Println(a)
	result := coll.FindOne(context.TODO(), bson.M{"username": a.Username, "email": a.Email})
	if result.Err() != nil {
		return coll.InsertOne(context.TODO(), &a)
	}
	return nil, errors.New("account already exists")
}

func (a *GetCreateAccount) Get(coll *mongo.Collection) error {
	filter := bson.D{{"sub", a.Sub}}
	err := coll.FindOne(context.TODO(), filter).Decode(&a)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("account not found")
		}
		return err
	}
	return nil
}

type UpdateAccount struct {
	Username  *string
	Email     *string
	FirstName *string
	LastName  *string
}

func Update(coll *mongo.Collection, sub string, payload UpdateAccount) error {
	filter := bson.M{"sub": sub}

	var account GetCreateAccount
	if err := coll.FindOne(context.TODO(), filter).Decode(&account); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("account not found")
		}
		return err
	}

	updateAccountStruct(&account, &payload)

	myBSONM, err := createBSONM(account)
	if err != nil {
		return err
	}

	result := coll.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": myBSONM})
	if result.Err() != nil {
		return result.Err()
	}

	log.Printf("Updated account: %s with: %v", sub, payload)
	return nil
}

func createBSONM(myStruct interface{}) (bson.M, error) {
	// Marshal the MyStruct object into a BSON document
	myStructBytes, err := bson.Marshal(myStruct)
	if err != nil {
		return nil, err
	}

	// Unmarshal the BSON document into a bson.M map
	var myBSONM bson.M
	err = bson.Unmarshal(myStructBytes, &myBSONM)
	if err != nil {
		return nil, err
	}

	// Modify the bson.M map to use the JSON names of the fields
	s := reflect.ValueOf(myStruct)
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		tag := typeOfT.Field(i).Tag.Get("json")

		// Remove the original key and replace it with the JSON name
		if key, ok := myBSONM[typeOfT.Field(i).Name]; ok {
			delete(myBSONM, typeOfT.Field(i).Name)
			myBSONM[tag] = key
		}
	}

	return myBSONM, nil
}

func updateAccountStruct(main *GetCreateAccount, update *UpdateAccount) {
	if update.Username != nil {
		main.Username = *update.Username
	}
	if update.Email != nil {
		main.Email = *update.Email
	}
	if update.FirstName != nil {
		main.FirstName = *update.FirstName
	}
	if update.LastName != nil {
		main.LastName = *update.LastName
	}
}
