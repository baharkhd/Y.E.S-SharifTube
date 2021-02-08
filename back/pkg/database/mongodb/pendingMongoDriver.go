package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"yes-sharifTube/internal/model/pending"
)

type PendingMongoDriver struct {
	collection *mongo.Collection
}

func (p PendingMongoDriver) GetByFilter(courseID primitive.ObjectID, status pending.Status, uploaderUsername string, start, amount int) ([]*pending.Pending, error) {
	panic("not implemented")
}
func (p PendingMongoDriver) Insert(username string, courseID primitive.ObjectID, pending *pending.Pending) (*pending.Pending, error) {
	panic("not implemented")
}
func (p PendingMongoDriver) Update(username string, courseID primitive.ObjectID, pending *pending.Pending) (*pending.Pending, error) {
	panic("not implemented")
}
func (p PendingMongoDriver) Delete(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error) {
	panic("not implemented")
}
func (p PendingMongoDriver) Accept(username string, courseID primitive.ObjectID, pending *pending.Pending) (*pending.Pending, error) {
	panic("not implemented")
}
func (p PendingMongoDriver) Reject(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error) {
	panic("not implemented")
}

func NewPendingMongoDriver(db, collection string) *PendingMongoDriver {
	return &PendingMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}