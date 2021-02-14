package attachment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBDriver interface {
	Get(courseID *primitive.ObjectID, attachmentID primitive.ObjectID) (*Attachment, error)
	Insert(courseID primitive.ObjectID, attachment *Attachment) (*Attachment, error)
	UpdateInfo(courseID, attachmentID primitive.ObjectID, name, description string, timestamp int64) error
	Delete(courseID, attachmentID primitive.ObjectID) error
}
