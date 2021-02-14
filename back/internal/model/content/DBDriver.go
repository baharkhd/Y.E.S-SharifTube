package content

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBDriver interface {
	Get(courseID *primitive.ObjectID, contentID primitive.ObjectID) (*Content, error)
	GetAll(courseID *primitive.ObjectID, tags []string, start, amount int) ([]*Content, error)
	Insert(courseID primitive.ObjectID, content *Content) (*Content, error)
	UpdateInfo(courseID, contentID primitive.ObjectID, title, description string, tags []string, timestamp int64)  error
	Delete(courseID, contentID primitive.ObjectID) error
}