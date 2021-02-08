package pending

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Pending struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Status       Status             `json:"status" bson:"status"`
	Timestamp    int64              `json:"timestamp" bson:"timestamp"`
	UploadedByID string             `json:"uploaded_by" bson:"uploaded_by"`
	Furl         string             `json:"furl" bson:"furl"` //todo better implementation
}

func New(title, description, uploadedByID, furl string) (*Pending, error) {
	//todo regex checking for url field

	return &Pending{
		Title:        title,
		Description:  description,
		Status:       PENDING,
		Timestamp:    time.Now().Unix(),
		UploadedByID: uploadedByID,
		Furl:         furl,
	}, nil
}
