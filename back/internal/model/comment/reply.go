package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
)

type Reply struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Body      string             `json:"body" bson:"body"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	AuthorUn  string             `json:"author" bson:"author"`
	CommentID string             `json:"replyTo" bson:"replyTo"`
}

func NewReply(body, authorID, commentID string) (*Reply, error) {
	err := RegexValidate(&body, &authorID, &commentID)
	if err != nil {
		return nil, err
	}
	return &Reply{
		Body:      body,
		Timestamp: time.Now().Unix(),
		AuthorUn:  authorID,
		CommentID: commentID,
	}, nil
}

func (r *Reply) Update(newBody *string) error {
	if newBody == nil {
		return model.EmptyFieldsException{Message: model.EmptyKeyErrorMessage}
	}
	err := RegexValidate(newBody, nil, nil)
	if err != nil {
		return err
	}
	r.Body = *newBody
	r.Timestamp = time.Now().Unix()
	return nil
}
