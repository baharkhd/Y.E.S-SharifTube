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

func NewReply(ID primitive.ObjectID, body, authorID, commentID string) (*Reply, error) {
	err := RegexValidate(&body, &authorID, &commentID)
	if err != nil {
		return nil, err
	}
	return &Reply{
		ID:        ID,
		Body:      body,
		Timestamp: time.Now().Unix(),
		AuthorUn:  authorID,
		CommentID: commentID,
	}, nil
}

func (r Reply) Reshape() (*model.Reply, error) {
	// todo get author from database by its username
	var author *model.User

	return &model.Reply{
		ID:        r.ID.Hex(),
		Author:    author,
		Body:      r.Body,
		Timestamp: int(r.Timestamp),
		CommentID: r.CommentID,
	}, nil
}

func ReshapeAllReplies(replies []*Reply) ([]*model.Reply, error) {
	var cs []*model.Reply
	for _, c := range replies {
		tmp, err := c.Reshape()
		if err != nil {
			return nil, model.InternalServerException{Message: "error while reshape reply array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
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
