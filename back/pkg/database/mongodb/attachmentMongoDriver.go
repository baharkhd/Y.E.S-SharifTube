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

func (a AttachmentMongoDriver) Insert(username string, courseID primitive.ObjectID, name, arul string, description *string) (*attachment.Attachment, error) {
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
	atch, err := attachment.New(primitive.NewObjectID(), name, arul, fc.ID.Hex(), description)
	if err != nil {
		return nil, err
	}
	change := bson.M{
		"$push": bson.M{
			"inventory": atch,
		},
	}
	if _, err = a.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return atch, nil
}

func (a AttachmentMongoDriver) Update(username string, courseID, attachmentID primitive.ObjectID, name, description *string) (*attachment.Attachment, error) {
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
	err := natch.Update(name, description)
	if err != nil {
		return nil, err
	}
	target = bson.M{
		"_id":           courseID,
		"inventory._id": natch.ID,
	}
	change := bson.M{
		"$set": bson.M{
			"inventory.$.name":        natch.Name,
			"inventory.$.description": natch.Description,
			"inventory.$.timestamp":   natch.Timestamp,
		},
	}
	if _, err = a.collection.UpdateOne(ctx, target, change); err != nil {
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
