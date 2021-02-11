package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Body      string             `json:"body" bson:"body"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	AuthorUn  string             `json:"author" bson:"author"`
	Replies   []*Reply           `json:"replies" bson:"replies"`
	ContentID string             `json:"content" bson:"content"`
}

func New(body, authorID, contentID string) *Comment {
	return &Comment{
		Body:      body,
		Timestamp: time.Now().Unix(),
		AuthorUn:  authorID,
		Replies:   []*Reply{},
		ContentID: contentID,
	}
}

func (c Comment) Reshape() (*model.Comment, error) {
	// todo get author from database by its username
	var author *model.User

	res := &model.Comment{
		ID:        c.ID.Hex(),
		Author:    author,
		Body:      c.Body,
		Timestamp: int(c.Timestamp),
		Replies:   nil,
		ContentID: c.ContentID,
	}

	//reshape replies
	replies, err := ReshapeAllReplies(c.Replies)
	if err != nil {
		return nil, &model.InternalServerException{Message: "error while reshape replies of comment: /n" + err.Error()}
	}
	res.Replies = replies

	return res, nil
}

func ReshapeAll(courses []*Comment) ([]*model.Comment, error) {
	var cs []*model.Comment
	for _, c := range courses {
		tmp, err := c.Reshape()
		if err != nil {
			return nil, &model.InternalServerException{Message: "error while reshape comment array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}

func (c *Comment) Update(newBody string) {
	c.Body = newBody
	c.Timestamp = time.Now().Unix()
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
