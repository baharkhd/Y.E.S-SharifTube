package course

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/pending"
)

type Course struct {
	ID        primitive.ObjectID       `bson:"_id" json:"id,omitempty"`
	Title     string                   `json:"title" bson:"title"`
	Summery   string                   `json:"summery" bson:"summery"`
	CreatedAt int64                    `json:"created_at" bson:"created_at"`
	Token     string                   `json:"token" bson:"token"`
	ProfUn    string                   `json:"prof" bson:"prof"`
	TaUns     []string                 `json:"tas" bson:"tas"`
	StdUns    []string                 `json:"students" bson:"students"`
	Contents  []*content.Content       `json:"contents" bson:"contents"`
	Pends     []*pending.Pending       `json:"pends" bson:"pends"`
	Inventory []*attachment.Attachment `json:"inventory" bson:"inventory"`
}

func New(title, summery, profUsername, token string) (*Course, error) {
	hashedToken, err := modelUtil.HashToken([]byte(token))
	if err != nil {
		return nil, model.InternalServerException{Message: "internal server error: couldn't hash token"}
	}

	return &Course{
		Title:     title,
		Summery:   summery,
		CreatedAt: time.Now().Unix(),
		ProfUn:    profUsername,
		Token:     hashedToken,
		TaUns:     []string{},
		StdUns:    []string{},
		Contents:  []*content.Content{},
		Pends:     []*pending.Pending{},
		Inventory: []*attachment.Attachment{},
	}, nil
}

func (c Course) Reshape() (*model.Course, error) {
	//todo get Users from database by usernames
	var prof *model.User
	var tas []*model.User
	var students []*model.User

	res := &model.Course{
		ID:        c.ID.Hex(),
		Title:     c.Title,
		Summary:   &c.Summery,
		CreatedAt: int(c.CreatedAt),
		Prof:      prof,
		Tas:       tas,
		Pends:     nil,
		Students:  students,
		Contents:  nil,
		Inventory: nil,
	}

	//reshape pendings
	pends, err := pending.ReshapeAll(c.Pends)
	if err != nil {
		return nil, &model.InternalServerException{Message: "error while reshape pending array of course: /n" + err.Error()}
	}
	res.Pends = pends

	//reshape contents
	contents, err := content.ReshapeAll(c.Contents)
	if err != nil {
		return nil, &model.InternalServerException{Message: "error while reshape contents of course: /n" + err.Error()}
	}
	res.Contents = contents

	//reshape inventory
	res.Inventory = attachment.ReshapeAll(c.Inventory)

	return res, nil
}

func ReshapeAll(courses []*Course) ([]*model.Course, error) {
	var cs []*model.Course
	for _, c := range courses {
		tmp, err := c.Reshape()
		if err != nil {
			return nil, &model.InternalServerException{Message: "error while reshape course array: " + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}
