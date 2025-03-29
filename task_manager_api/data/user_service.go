package data

import (
	"context"
	"errors"

	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection

func CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	count, _ := UserCollection.CountDocuments(ctx, bson.D{})
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPwd)
	user.ID = primitive.NewObjectID()
	return UserCollection.InsertOne(ctx, user)
}

func AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}
	return &user, nil
}

func PromoteUser(ctx context.Context, username string) error {
	result, err := UserCollection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil || result.MatchedCount == 0 {
		return errors.New("user not found or promotion failed")
	}
	return nil
}
