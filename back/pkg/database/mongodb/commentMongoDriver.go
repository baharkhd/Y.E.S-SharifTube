package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"yes-sharifTube/internal/model/comment"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/pkg/database"
)

type CommentMongoDriver struct {
	collection *mongo.Collection
}

func (c CommentMongoDriver) Insert(username string, contentID primitive.ObjectID, repliedAtID *primitive.ObjectID, body string) (*comment.Comment, *comment.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"contents": bson.M{
			"$elemMatch": bson.M{
				"_id": contentID,
			},
		},
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
	if err := c.collection.FindOne(ctx, target, options.FindOne().SetProjection(projection)).Decode(&fc); err != nil {
		return nil, nil, database.ThrowContentNotFoundException(contentID.Hex())
	}
	cnt := fc.Contents[0]
	if !fc.IsUserParticipateInCourse(username) {
		return nil, nil, database.ThrowUserNotAllowedException(username)
	}
	var cmd *comment.Comment = nil
	var rep *comment.Reply = nil
	var err error
	if repliedAtID == nil {
		cmd, err = comment.New(primitive.NewObjectID(), body, username, cnt.ID.Hex())
		if err != nil {
			return nil, nil, err
		}
		cnt.Comments = append(cnt.Comments, cmd)
	} else {
		ctmd, _ := cnt.GetComment(*repliedAtID)
		if ctmd == nil {
			return nil, nil, database.ThrowCommentNotFoundException(repliedAtID.Hex())
		}
		rep, err = comment.NewReply(primitive.NewObjectID(), body, username, repliedAtID.Hex())
		if err != nil {
			return nil, nil, err
		}
		ctmd.Replies = append(ctmd.Replies, rep)
	}
	if err = c.UpdateCommentsForContent(ctx, fc.ID, cnt.ID, cnt.Comments); err != nil {
		return nil, nil, err
	}
	if rep != nil {
		return nil, rep, nil
	}
	return cmd, rep, nil
}

func (c CommentMongoDriver) Update(username string, contentID, commentID primitive.ObjectID, body *string) (*comment.Comment, *comment.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"contents": bson.M{
			"$elemMatch": bson.M{
				"_id": contentID,
			},
		},
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
	if err := c.collection.FindOne(ctx, target, options.FindOne().SetProjection(projection)).Decode(&fc); err != nil {
		return nil, nil, database.ThrowContentNotFoundException(contentID.Hex())
	}
	cnt := fc.Contents[0]
	ccmt, rcmt := cnt.GetComment(commentID)
	if ccmt == nil {
		return nil, nil, database.ThrowCommentNotFoundException(commentID.Hex())
	}
	if rcmt == nil {
		if username != ccmt.AuthorUn {
			return nil, nil, database.ThrowUserNotAllowedException(username)
		}
		err := ccmt.Update(body)
		if err != nil {
			return nil, nil, err
		}
	} else {
		if username != rcmt.AuthorUn {
			return nil, nil, database.ThrowUserNotAllowedException(username)
		}
		err := rcmt.Update(body)
		if err != nil {
			return nil, nil, err
		}
	}
	if err := c.UpdateCommentsForContent(ctx, fc.ID, cnt.ID, cnt.Comments); err != nil {
		return nil, nil, err
	}
	if rcmt != nil {
		return nil, rcmt, nil
	}
	return ccmt, rcmt, nil
}

func (c CommentMongoDriver) Delete(username string, contentID, commentID primitive.ObjectID) (*comment.Comment, *comment.Reply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"contents": bson.M{
			"$elemMatch": bson.M{
				"_id": contentID,
			},
		},
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
	if err := c.collection.FindOne(ctx, target, options.FindOne().SetProjection(projection)).Decode(&fc); err != nil {
		return nil, nil, database.ThrowContentNotFoundException(contentID.Hex())
	}
	cnt := fc.Contents[0]
	ccmt, rcmt, err := fc.RemoveComment(username, commentID, cnt)
	if err != nil {
		return nil, nil, err
	}
	if err = c.UpdateCommentsForContent(ctx, fc.ID, cnt.ID, cnt.Comments); err != nil {
		return nil, nil, err
	}
	if rcmt != nil {
		return nil, rcmt, nil
	}
	return ccmt, rcmt, nil
}

func (c CommentMongoDriver) UpdateCommentsForContent(ctx context.Context, courseID, contentID primitive.ObjectID, comments []*comment.Comment) error {
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
		return database.ThrowInternalDBException(err.Error())
	}
	return nil
}

func NewCommentMongoDriver(db, collection string) *CommentMongoDriver {
	return &CommentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
