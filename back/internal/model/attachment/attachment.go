package attachment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Attachment struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Timestamp   int64              `json:"timestamp" bson:"timestamp"`
	Aurl        string             `json:"aurl" bson:"aurl"`  //todo better implementation
}

func New(name, description, aurl string) (*Attachment, error) {
	//todo regex checking for url field

	return &Attachment{
		Name:        name,
		Description: description,
		Timestamp:   time.Now().Unix(),
		Aurl:        aurl,
	}, nil
}
