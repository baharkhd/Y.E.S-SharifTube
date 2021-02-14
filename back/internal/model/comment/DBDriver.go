package comment

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBDriver interface {
	UpdateCommentsForContent(courseID, contentID primitive.ObjectID, comments []*Comment) error
	InsertComment(courseID, contentID primitive.ObjectID, comment *Comment) (*Comment,error)
	UpdateComment(courseID, contentID, commentID primitive.ObjectID, newBody string, timestamp int64) error
	DeleteComment(courseID, contentID primitive.ObjectID, comment *Comment) error
	InsertReply(courseID, contentID, repliedAtID primitive.ObjectID, reply *Reply) (*Reply, error)
	UpdateReply(courseID, contentID, replyID primitive.ObjectID, newBody string, timestamp int64) error
	DeleteReply(courseID, contentID, replyID primitive.ObjectID) error

}
