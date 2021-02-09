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
	AuthorID  string             `json:"author" bson:"author"`
	CommentID string             `json:"replyTo" bson:"replyTo"`
}

func NewReply(body, authorID, commentID string) *Reply {
	return &Reply{
		Body:      body,
		Timestamp: time.Now().Unix(),
		AuthorID:  authorID,
		CommentID: commentID,
	}
}

func (c Reply) Reshape() (*model.Reply, error) {
	// todo get author from database by its username
	var author *model.User

	return &model.Reply{
		ID:        c.ID.Hex(),
		Author:    author,
		Body:      c.Body,
		Timestamp: int(c.Timestamp),
		CommentID: c.CommentID,
	}, nil
}

func ReshapeAllReplies(replies []*Reply) ([]*model.Reply, error) {
	var cs []*model.Reply
	for _, c := range replies {
		tmp, err := c.Reshape()
		if err != nil {
			return nil, &model.InternalServerException{Message: "error while reshape reply array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}