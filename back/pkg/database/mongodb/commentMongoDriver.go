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

func (c CommentMongoDriver) Insert(username string, contentID primitive.ObjectID, repliedAtID *primitive.ObjectID, cmd *comment.Comment) (*comment.Comment, *comment.Reply, error) {
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
	var rep *comment.Reply = nil
	cmd.ID = primitive.NewObjectID()
	if repliedAtID == nil {
		cmd.ContentID = cnt.ID.Hex()
		cnt.Comments = append(cnt.Comments, cmd)
	} else {
		ctmd, _ := cnt.GetComment(*repliedAtID)
		if ctmd == nil {
			return nil, nil, database.ThrowCommentNotFoundException(repliedAtID.Hex())
		}
		rep = cmd.ConvertToReply(*repliedAtID)
		ctmd.Replies = append(ctmd.Replies, rep)
	}
	if err := c.UpdateCommentsForContent(ctx, fc.ID, cnt.ID, cnt.Comments); err != nil {
		return nil, nil, err
	}
	if rep != nil {
		return nil, rep, nil
	}
	return cmd, rep, nil
}

func (c CommentMongoDriver) Update(username string, contentID primitive.ObjectID, cmt *comment.Comment) (*comment.Comment, *comment.Reply, error) {
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
	ccmt, rcmt := cnt.GetComment(cmt.ID)
	if ccmt == nil {
		return nil, nil, database.ThrowCommentNotFoundException(cmt.ID.Hex())
	}
	if rcmt == nil {
		if username != ccmt.AuthorUn {
			return nil, nil, database.ThrowUserNotAllowedException(username)
		}
		ccmt.Update(cmt.Body)
	} else {
		if username != rcmt.AuthorUn {
			return nil, nil, database.ThrowUserNotAllowedException(username)
		}
		rcmt.Update(cmt.Body)
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
		if err.Error() == "NotAllowed" {
			return nil, nil, database.ThrowUserNotAllowedException(username)
		} else
		if err.Error() == "NotFound" {
			return nil, nil, database.ThrowCommentNotFoundException(commentID.Hex())
		} else {
			return nil, nil, err
		}
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
