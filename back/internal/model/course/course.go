package course

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/pending"
	"yes-sharifTube/internal/model/user"
)

type Course struct {
	ID        primitive.ObjectID       `bson:"_id" json:"id,omitempty"`
	Title     string                   `json:"title" bson:"title"`
	Summery   string                   `json:"summery" bson:"summery"`
	CreatedAt int64                    `json:"created_at" bson:"created_at"`
	Token     string                   `json:"token" bson:"token"`
	ProfID    *user.User               `json:"prof" bson:"prof"`
	TaIDs     []*user.User             `json:"tas" bson:"tas"`
	StdIDs    []*user.User             `json:"students" bson:"students"`
	Contents  []*content.Content       `json:"contents" bson:"contents"`
	Pends     []*pending.Pending       `json:"pends" bson:"pends"`
	Inventory []*attachment.Attachment `json:"inventory" bson:"inventory"`
}

func New(title, summery, profUsername, token string) (*Course, error) {
	hashedToken, err := hashToken([]byte(token))
	if err != nil {
		return nil, model.InternalServerException{Message: "internal server error: couldn't hash password"}
	}

	return &Course{
		Title:     title,
		Summery:   summery,
		CreatedAt: time.Now().Unix(),
		ProfID: &user.User{
			Username: profUsername,
		},
		Token:     hashedToken,
		TaIDs:     []*user.User{},
		StdIDs:    []*user.User{},
		Contents:  []*content.Content{},
		Pends:     []*pending.Pending{},
		Inventory: []*attachment.Attachment{},
	}, nil
}

