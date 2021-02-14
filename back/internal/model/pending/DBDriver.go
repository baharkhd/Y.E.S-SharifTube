package pending

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBDriver interface {
	Get(courseID *primitive.ObjectID, pendingID primitive.ObjectID) (*Pending, error)
	GetByFilter(courseID *primitive.ObjectID, status *Status, uploaderUsername *string, start, amount int) ([]*Pending, error)
	Insert(courseID primitive.ObjectID, pending *Pending) (*Pending, error)
	UpdateInfo(courseID, pendingID primitive.ObjectID, title, description string, timestamp int64)  error
	Delete(courseID, pendingID primitive.ObjectID) error
	UpdateStatus(courseID, pendingID primitive.ObjectID, newTitle, newDescription string, status Status, timestamp int64) error
}