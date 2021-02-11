package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/pkg/database"
)

type AttachmentMongoDriver struct {
	collection *mongo.Collection
}

func (a AttachmentMongoDriver) Insert(username string, courseID primitive.ObjectID, atch *attachment.Attachment) (*attachment.Attachment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := a.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	atch.ID = primitive.NewObjectID()
	atch.CourseID = fc.ID.Hex()
	change := bson.M{
		"$push": bson.M{
			"inventory": atch,
		},
	}
	if _, err := a.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return atch, nil
}

func (a AttachmentMongoDriver) Update(username string, courseID primitive.ObjectID, atch *attachment.Attachment) (*attachment.Attachment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := a.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	natch := fc.GetAttachment(atch.ID)
	if natch == nil {
		return nil, database.ThrowAttachmentNotFoundException(atch.ID.Hex())
	}
	natch.Update(atch.Name, atch.Description)
	target = bson.M{
		"_id":           courseID,
		"inventory._id": atch.ID,
	}
	change := bson.M{
		"$set": bson.M{
			"inventory.$.name":        atch.Name,
			"inventory.$.description": atch.Description,
			"inventory.$.timestamp":   atch.Timestamp,
		},
	}
	if _, err := a.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return natch, nil
}

func (a AttachmentMongoDriver) Delete(username string, courseID, attachmentID primitive.ObjectID) (*attachment.Attachment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := a.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	natch := fc.GetAttachment(attachmentID)
	if natch == nil {
		return nil, database.ThrowAttachmentNotFoundException(attachmentID.Hex())
	}
	change := bson.M{
		"$pull": bson.M{
			"inventory": bson.M{
				"_id": attachmentID,
			},
		},
	}
	if _, err := a.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return natch, nil
}

func NewAttachmentMongoDriver(db, collection string) *AttachmentMongoDriver {
	return &AttachmentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
