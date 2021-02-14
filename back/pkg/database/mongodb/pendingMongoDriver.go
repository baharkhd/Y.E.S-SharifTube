package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/pending"
)

type PendingMongoDriver struct {
	collection *mongo.Collection
}

func (p PendingMongoDriver) Get(courseID *primitive.ObjectID, pendingID primitive.ObjectID) (*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	var res course.Course
	var target bson.M
	if courseID == nil {
		target = bson.M{
			"pends": bson.M{
				"$elemMatch": bson.M{
					"_id": pendingID,
				},
			},
		}
	} else {
		target = bson.M{
			"_id": courseID,
			"pends": bson.M{
				"$elemMatch": bson.M{
					"_id": pendingID,
				},
			},
		}
	}
	projection := bson.M{
		"created_at": 1,
		"inventory":  1,
		"contents":   1,
		"prof":       1,
		"students":   1,
		"summery":    1,
		"tas":        1,
		"title":      1,
		"token":      1,
		"pends.$":    1,
	}
	if err := p.collection.FindOne(ctx, target, options.FindOne().SetProjection(projection)).Decode(&res); err != nil {
		return nil, model.PendingNotFoundException{Message: "pending couldn't found."}
	}
	return res.Pends[0], nil
}

func (p PendingMongoDriver) GetByFilter(courseID *primitive.ObjectID, status *pending.Status, uploaderUsername *string, start, amount int) ([]*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	filter := bson.A{}
	if courseID != nil {
		filter = append(filter,
			bson.M{
				"$eq": bson.A{
					"$$item.course",
					courseID.Hex(),
				},
			},
		)
	}
	if status != nil {
		filter = append(filter,
			bson.M{
				"$eq": bson.A{
					"$$item.status",
					status,
				},
			},
		)
	}
	if uploaderUsername != nil {
		filter = append(filter,
			bson.M{
				"$eq": bson.A{
					"$$item.uploaded_by",
					uploaderUsername,
				},
			},
		)
	}
	pipeline := []bson.M{
		{
			"$project": bson.M{
				"created_at": 1,
				"inventory":  1,
				"prof":       1,
				"students":   1,
				"summery":    1,
				"tas":        1,
				"title":      1,
				"token":      1,
				"contents":   1,
				"pends": bson.M{
					"$filter": bson.M{
						"input": "$pends",
						"as":    "item",
						"cond": bson.M{
							"$and": filter,
						},
					},
				},
			},
		},
	}

	courr, err := p.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	defer courr.Close(ctx)
	var pendings []*pending.Pending
	i := 0
	for courr.Next(context.Background()) {
		var ctmp course.Course
		_ = courr.Decode(&ctmp)
		for j, _ := range ctmp.Pends {
			pendings = append(pendings, ctmp.Pends[j])
		}
		i++
	}
	return pending.GetAll(pendings, start, amount), nil

}

func (p PendingMongoDriver) Insert(courseID primitive.ObjectID, pending *pending.Pending) (*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	pending.ID = primitive.NewObjectID()
	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$push": bson.M{
			"pends": pending,
		},
	}
	if _, err := p.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return pending, nil
}

func (p PendingMongoDriver) UpdateInfo(courseID, pendingID primitive.ObjectID, title, description string, timestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id":       courseID,
		"pends._id": pendingID,
	}
	change := bson.M{
		"$set": bson.M{
			"pends.$.title":       title,
			"pends.$.description": description,
			"pends.$.timestamp":   timestamp,
		},
	}
	if _, err := p.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func (p PendingMongoDriver) Delete(courseID, pendingID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$pull": bson.M{
			"pends": bson.M{
				"_id": pendingID,
			},
		},
	}
	if _, err := p.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func (p PendingMongoDriver) UpdateStatus(courseID, pendingID primitive.ObjectID, newTitle, newDescription string, status pending.Status, timestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id":       courseID,
		"pends._id": pendingID,
	}
	change := bson.M{
		"$set": bson.M{
			"pends.$.title":       newTitle,
			"pends.$.description": newDescription,
			"pends.$.status":      status,
			"pends.$.timestamp":   timestamp,
		},
	}
	if _, err := p.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func NewPendingMongoDriver(db, collection string) *PendingMongoDriver {
	return &PendingMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
