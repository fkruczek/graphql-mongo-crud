package db

import (
	"context"

	"github.com/fk/gqlplayground/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (db *Database) CreateUser(newUser *model.NewUser) error {
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashedPassword
	_, err = db.Collection("users").InsertOne(context.Background(), newUser)
	return nil
}

func (db *Database) GetUserIdByUsername(username string) (string, error) {
	var user *model.User
	result := db.Collection("users").FindOne(context.Background(), bson.M{"username": username})
	err := result.Decode(&user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (db *Database) GetPasswordByUsername(username string) (string, error) {
	var user model.Login
	result := db.Collection("users").FindOne(context.Background(), bson.M{"username": username})
	err := result.Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
