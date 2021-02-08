package content

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/internal/model/comment"
)

type Content struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Timestamp    int64              `json:"timestamp" bson:"timestamp"`
	UploadedByID string             `json:"uploaded_by" bson:"uploaded_by"`
	ApprovedByID string             `json:"approved_by" bson:"approved_by"`
	Vurl         string             `json:"vurl" bson:"vurl"` //todo better implementation
	Tags         []string           `json:"tags" bson:"tags"`
	Comments     []*comment.Comment `json:"comments" bson:"comments"`
}

func New(title, description, uploadedBy, approvedBy, vurl string, tags []string) (*Content, error) {
	//todo regex checking for url field

	return &Content{
		Title:        title,
		Description:  description,
		Timestamp:    time.Now().Unix(),
		UploadedByID: uploadedBy,
		ApprovedByID: approvedBy,
		Vurl:         vurl,
		Tags:         tags,
		Comments:     []*comment.Comment{},
	}, nil
}
