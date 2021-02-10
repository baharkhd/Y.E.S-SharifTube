package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/pkg/database/mongodb"
	)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string
	Email    string
	Username string
	Password string
	Courses  []string
}
var UserDBD *mongodb.UserMongoDriver

func New(name, email, username, password string) User {
	user := User{
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}
	return user
}


