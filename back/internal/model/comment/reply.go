package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reply struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Body      string             `json:"body" bson:"body"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	AuthorID  string             `json:"author" bson:"author"`
}

func NewReply(body, authorID string) *Reply {
	return &Reply{
		Body:      body,
		Timestamp: time.Now().Unix(),
		AuthorID:  authorID,
	}
}
