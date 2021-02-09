package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"yes-sharifTube/internal/model/comment"
)

type CommentMongoDriver struct {
	collection *mongo.Collection
}

func (c CommentMongoDriver) Insert(username string, contentID primitive.ObjectID, repliedAtID *primitive.ObjectID, comment *comment.Comment) (*comment.Comment, *comment.Reply, error) {
	panic("not implemented")
}
func (c CommentMongoDriver) Update(username string, contentID primitive.ObjectID, comment *comment.Comment) (*comment.Comment, *comment.Reply, error) {
	panic("not implemented")
}
func (c CommentMongoDriver) Delete(username string, contentID, commentID primitive.ObjectID) (*comment.Comment, *comment.Reply, error) {
	panic("not implemented")
}

func NewCommentMongoDriver(db, collection string) *CommentMongoDriver {
	return &CommentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}