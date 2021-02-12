package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username" bson:"username"`
}
