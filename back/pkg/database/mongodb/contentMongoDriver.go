package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
)

type ContentMongoDriver struct {
	collection *mongo.Collection
}

func (c *ContentMongoDriver) Get(courseID *primitive.ObjectID, contentID primitive.ObjectID) (*content.Content, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	var res course.Course
	var target bson.M
	if courseID == nil {
		target = bson.M{
			"contents": bson.M{
				"$elemMatch": bson.M{
					"_id": contentID,
				},
			},
		}
	} else {
		target = bson.M{
			"_id": courseID,
			"contents": bson.M{
				"$elemMatch": bson.M{
					"_id": contentID,
				},
			},
		}
	}
	projection := bson.M{
		"created_at": 1,
		"inventory":  1,
		"pends":      1,
		"prof":       1,
		"students":   1,
		"summery":    1,
		"tas":        1,
		"title":      1,
		"token":      1,
		"contents.$": 1,
	}
	if err := c.collection.FindOne(ctx, target, options.FindOne().SetProjection(projection)).Decode(&res); err != nil {
		return nil, model.ContentNotFoundException{Message: "content couldn't found."}
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
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
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

func (c *ContentMongoDriver) Insert(courseID primitive.ObjectID, content *content.Content) (*content.Content, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	content.ID = primitive.NewObjectID()
	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$push": bson.M{
			"contents": content,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return content, nil
}

func (c *ContentMongoDriver) UpdateInfo(courseID, contentID primitive.ObjectID, title, description string, tags []string, timestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id":          courseID,
		"contents._id": contentID,
	}
	change := bson.M{
		"$set": bson.M{
			"contents.$.title":       title,
			"contents.$.description": description,
			"contents.$.tags":        tags,
			"contents.$.timestamp":   timestamp,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func (c *ContentMongoDriver) Delete(courseID, contentID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$pull": bson.M{
			"contents": bson.M{
				"_id": contentID,
			},
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func NewContentMongoDriver(db, collection string) *ContentMongoDriver {
	return &ContentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
