package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/course"
)

type AttachmentMongoDriver struct {
	collection *mongo.Collection
}

func (a AttachmentMongoDriver) Get(courseID *primitive.ObjectID, attachmentID primitive.ObjectID) (*attachment.Attachment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	var res course.Course
	var target bson.M
	if courseID == nil {
		target = bson.M{
			"inventory": bson.M{
				"$elemMatch": bson.M{
					"_id": attachmentID,
				},
			},
		}
	} else {
		target = bson.M{
			"_id": courseID,
			"inventory": bson.M{
				"$elemMatch": bson.M{
					"_id": attachmentID,
				},
			},
		}
	}
	projection := bson.M{
		"created_at":  1,
		"contents":    1,
		"pends":       1,
		"prof":        1,
		"students":    1,
		"summery":     1,
		"tas":         1,
		"title":       1,
		"token":       1,
		"inventory.$": 1,
	}
	if err := a.collection.FindOne(ctx, target, options.FindOne().SetProjection(projection)).Decode(&res); err != nil {
		return nil, model.AttachmentNotFoundException{Message: "attachment couldn't found."}
	}
	return res.Inventory[0], nil
}

func (a AttachmentMongoDriver) Insert(courseID primitive.ObjectID, attachment *attachment.Attachment) (*attachment.Attachment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	attachment.ID = primitive.NewObjectID()
	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$push": bson.M{
			"inventory": attachment,
		},
	}
	if _, err := a.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return attachment, nil
}

func (a AttachmentMongoDriver) UpdateInfo(courseID, attachmentID primitive.ObjectID, name, description string, timestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id":           courseID,
		"inventory._id": attachmentID,
	}
	change := bson.M{
		"$set": bson.M{
			"inventory.$.name":        name,
			"inventory.$.description": description,
			"inventory.$.timestamp":   timestamp,
		},
	}
	if _, err := a.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func (a AttachmentMongoDriver) Delete(courseID, attachmentID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$pull": bson.M{
			"inventory": bson.M{
				"_id": attachmentID,
			},
		},
	}
	if _, err := a.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func NewAttachmentMongoDriver(db, collection string) *AttachmentMongoDriver {
	return &AttachmentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
