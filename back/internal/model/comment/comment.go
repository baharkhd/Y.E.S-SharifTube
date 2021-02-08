package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Body      string             `json:"body" bson:"body"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	AuthorID  string             `json:"author" bson:"author"`
	Replies []*Reply             `json:"replies" bson:"replies"`
}

func New (body, authorID string) *Comment {
	return &Comment{
		Body:      body,
		Timestamp: time.Now().Unix(),
		AuthorID:  authorID,
		Replies:   []*Reply{},
	}
}