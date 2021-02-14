package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Body      string             `json:"body" bson:"body"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	AuthorUn  string             `json:"author" bson:"author"`
	Replies   []*Reply           `json:"replies" bson:"replies"`
	ContentID string             `json:"content" bson:"content"`
}

var DBD DBDriver

func New(body, authorID, contentID string) (*Comment, error) {
	err := RegexValidate(&body, &authorID, &contentID)
	if err != nil {
		return nil, err
	}
	return &Comment{
		Body:      body,
		Timestamp: time.Now().Unix(),
		AuthorUn:  authorID,
		Replies:   []*Reply{},
		ContentID: contentID,
	}, nil
}

func RegexValidate(body, authorUn, ownerID *string) error {
	if body != nil && modelUtil.IsSTREmpty(*body) {
		return model.RegexMismatchException{Message: "body field is empty"}
	}
	if authorUn != nil && modelUtil.IsSTREmpty(*authorUn) {
		return model.RegexMismatchException{Message: "author username field is empty"}
	}
	if ownerID != nil {
		_, err := primitive.ObjectIDFromHex(*ownerID)
		if err != nil {
			return model.RegexMismatchException{Message: "courseID/commentID field is invalid"}
		}
	}
	return nil
}

func (c *Comment) Update(newBody *string) error {
	if newBody == nil {
		return model.EmptyFieldsException{Message: model.EmptyKeyErrorMessage}
	}
	err := RegexValidate(newBody, nil, nil)
	if err != nil {
		return err
	}
	c.Body = *newBody
	c.Timestamp = time.Now().Unix()
	return nil
}

func (c *Comment) ConvertToReply(repID primitive.ObjectID) *Reply {
	return &Reply{
		ID:        c.ID,
		Body:      c.Body,
		Timestamp: c.Timestamp,
		AuthorUn:  c.AuthorUn,
		CommentID: repID.Hex(),
	}
}
