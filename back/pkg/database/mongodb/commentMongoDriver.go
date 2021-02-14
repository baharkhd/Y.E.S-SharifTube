package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/comment"
)

type CommentMongoDriver struct {
	collection *mongo.Collection
}

func (c CommentMongoDriver) UpdateCommentsForContent(courseID, contentID primitive.ObjectID, comments []*comment.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id":          courseID,
		"contents._id": contentID,
	}
	change := bson.M{
		"$set": bson.M{
			"contents.$.comments": comments,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func (c CommentMongoDriver) InsertComment(courseID, contentID primitive.ObjectID, comment *comment.Comment) (*comment.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	comment.ID = primitive.NewObjectID()
	target := bson.M{
		"_id":          courseID,
		"contents._id": contentID,
	}
	change := bson.M{
		"$push": bson.M{
			"contents.$.comments": comment,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return comment, nil
}

func (c CommentMongoDriver) UpdateComment(courseID, contentID, commentID primitive.ObjectID, newBody string, timestamp int64) error {
	//todo implementation
	return nil
}

func (c CommentMongoDriver) DeleteComment(courseID, contentID primitive.ObjectID, comment *comment.Comment) error {
	//todo implementation
	return nil
}

func (c CommentMongoDriver) InsertReply(courseID, contentID, repliedAtID primitive.ObjectID, reply *comment.Reply) (*comment.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	reply.ID = primitive.NewObjectID()
	target := bson.D{
		{"_id", courseID},
		{"contents._id", contentID},
		//{"contents.comments._id", repliedAtID},
	}
	change := bson.M{
		"$push": bson.M{
			"contents.$[con].comments.$[elem].replies": reply,
		},
	}
	//todo error handling?
	_ = c.collection.FindOneAndUpdate(ctx, target, change, options.FindOneAndUpdate().SetArrayFilters(
		options.ArrayFilters{
			Filters: []interface{}{
				bson.M{"con._id": contentID},
				bson.M{"elem._id": repliedAtID,
				},
			},
		},
	))
	return reply, nil
}

func (c CommentMongoDriver) UpdateReply(courseID, contentID, replyID primitive.ObjectID, newBody string, timestamp int64) error {
	//todo implementation
	return nil
}

func (c CommentMongoDriver) DeleteReply(courseID, contentID, replyID primitive.ObjectID) error {
	//todo implementation
	return nil
}

func NewCommentMongoDriver(db, collection string) *CommentMongoDriver {
	return &CommentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
