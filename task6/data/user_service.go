package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCol *mongo.Collection

func SetUserCollection(client *mongo.Client) {
	userCol = client.Database("task6").Collection("users")
}

func RegisterUser(user models.User) error {
	// FIX 1: Capture 'cancel' here
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if username exists
	var existingUser models.User
	err := userCol.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return errors.New("username already exists")
	}
    
	// Check if database is empty to assign Admin role
	count, _ := userCol.CountDocuments(ctx, bson.M{})
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err = userCol.InsertOne(ctx, user)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	// FIX 2: Properly handle ctx and cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := userCol.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

func PromoteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := userCol.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
