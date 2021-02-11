package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/pkg/database"
)

type ContentMongoDriver struct {
	collection *mongo.Collection
}

func (c *ContentMongoDriver) Get(contentID primitive.ObjectID) (*content.Content, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	var res course.Course
	target := bson.M{
		"contents": bson.M{
			"$elemMatch": bson.M{
				"_id": contentID,
			},
		},
	}
	projection := bson.M{
		"contents.$": 1,
	}
	if err := c.collection.FindOne(ctx, target, options.FindOne().SetProjection(projection)).Decode(&res); err != nil {
		return nil, database.ThrowContentNotFoundException(contentID.Hex())
	}
	return res.Contents[0], nil
}

//todo search by given tags Intersect with content tags
func (c *ContentMongoDriver) GetAll(courseID *primitive.ObjectID, tags []string, start, amount int) ([]*content.Content, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	var pipeline []bson.M
	if courseID != nil {
		pipeline = append(pipeline,
			bson.M{
				"$match": bson.M{
					"_id": courseID,
				},
			},
		)
	}
	pipeline = append(pipeline,
		bson.M{
			"$project": bson.M{
				"created_at": 1,
				"inventory":  1,
				"pends":      1,
				"prof":       1,
				"students":   1,
				"summery":    1,
				"tas":        1,
				"title":      1,
				"token":      1,
				"contents": bson.M{
					"$filter": bson.M{
						"input": "$contents",
						"as":    "item",
						"cond": bson.M{
							"$setIsSubset": bson.A{tags, "$$item.tags"},
						},
					},
				},
			},
		},
	)

	courr, err := c.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	defer courr.Close(ctx)
	var contents []*content.Content
	i := 0
	for courr.Next(context.Background()) {
		var ctmp course.Course
		_ = courr.Decode(&ctmp)
		for j, _ := range ctmp.Contents {
			contents = append(contents, ctmp.Contents[j])
		}
		i++
	}
	return content.GetAll(contents, start, amount), nil
}

func (c *ContentMongoDriver) Insert(username string, courseID primitive.ObjectID, cnt *content.Content) (*content.Content, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	cnt.ID = primitive.NewObjectID()
	cnt.CourseID = fc.ID.Hex()
	change := bson.M{
		"$push": bson.M{
			"contents": cnt,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return cnt, nil
}

func (c *ContentMongoDriver) Update(username string, courseID primitive.ObjectID, cnt *content.Content) (*content.Content, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	ncnt := fc.GetContent(cnt.ID)
	if ncnt == nil {
		return nil, database.ThrowContentNotFoundException(cnt.ID.Hex())
	}
	ncnt.Update(cnt.Title, cnt.Description, cnt.Tags)
	target = bson.M{
		"_id":          courseID,
		"contents._id": cnt.ID,
	}
	change := bson.M{
		"$set": bson.M{
			"contents.$.title":       cnt.Title,
			"contents.$.description": cnt.Description,
			"contents.$.tags":        cnt.Tags,
			"contents.$.timestamp":   cnt.Timestamp,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return ncnt, nil
}

func (c *ContentMongoDriver) Delete(username string, courseID, contentID primitive.ObjectID) (*content.Content, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	ncnt := fc.GetContent(contentID)
	if ncnt == nil {
		return nil, database.ThrowContentNotFoundException(contentID.Hex())
	}
	change := bson.M{
		"$pull": bson.M{
			"contents": bson.M{
				"_id": contentID,
			},
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return ncnt, nil
}

func NewContentMongoDriver(db, collection string) *ContentMongoDriver {
	return &ContentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
