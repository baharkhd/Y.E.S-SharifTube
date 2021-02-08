package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"yes-sharifTube/internal/model/attachment"
)

type AttachmentMongoDriver struct {
	collection *mongo.Collection
}

func (a AttachmentMongoDriver) Insert(username string, courseID primitive.ObjectID, attachment *attachment.Attachment) (*attachment.Attachment, error) {
	panic("not implemented")
}
func (a AttachmentMongoDriver) Update(username string, courseID primitive.ObjectID, attachment *attachment.Attachment) (*attachment.Attachment, error) {
	panic("not implemented")
}
func (a AttachmentMongoDriver) Delete(username string, courseID, attachmentID primitive.ObjectID) (*attachment.Attachment, error) {
	panic("not implemented")
}

func NewAttachmentMongoDriver(db, collection string) *AttachmentMongoDriver {
	return &AttachmentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}