package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/comment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/pending"
)

type CourseDBDriver interface {
	GetAll(courseIDs []primitive.ObjectID) ([]*course.Course, error)
	GetByFilter(keywords []string, start, amount int) ([]*course.Course, error)
	Insert(username string, title, token string, summery *string) (*course.Course, error)
	Update(username string, courseID primitive.ObjectID, title, token, summery *string) (*course.Course, error)
	Delete(username string, courseID primitive.ObjectID) (*course.Course, error)
	AddUser(username, token string, courseID primitive.ObjectID) (*course.Course, error)
	DeleteUser(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error)
	PromoteToTA(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error)
	DemoteToSTD(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error)
}

type PendingDBDriver interface {
	GetByFilter(courseID *primitive.ObjectID, status *pending.Status, uploaderUsername *string, start, amount int) ([]*pending.Pending, error)
	Insert(username string, courseID primitive.ObjectID, title, furl string, description *string) (*pending.Pending, error)
	Update(username string, courseID, pendingID primitive.ObjectID, title, description *string) (*pending.Pending, error)
	Delete(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error)
	Accept(username string, courseID, pendingID primitive.ObjectID, title, description *string) (*pending.Pending, error)
	Reject(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error)
}

type ContentDBDriver interface {
	Get(contentID primitive.ObjectID) (*content.Content, error)
	GetAll(courseID *primitive.ObjectID, tags []string, start, amount int) ([]*content.Content, error)
	Insert(username string, courseID primitive.ObjectID, title, vrul string, description *string, tags []string) (*content.Content, error)
	Update(username string, courseID, contentID primitive.ObjectID, title, description *string, tags []string) (*content.Content, error)
	Delete(username string, courseID, contentID primitive.ObjectID) (*content.Content, error)
}

type AttachmentDBDriver interface {
	Insert(username string, courseID primitive.ObjectID, name, arul string, description *string) (*attachment.Attachment, error)
	Update(username string, courseID, attachmentID primitive.ObjectID, name, description *string) (*attachment.Attachment, error)
	Delete(username string, courseID, attachmentID primitive.ObjectID) (*attachment.Attachment, error)
}

type CommentDBDriver interface {
	Insert(username string, contentID primitive.ObjectID, repliedAtID *primitive.ObjectID, body string) (*comment.Comment, *comment.Reply, error)
	Update(username string, contentID, commentID primitive.ObjectID, body *string) (*comment.Comment, *comment.Reply, error)
	Delete(username string, contentID, commentID primitive.ObjectID) (*comment.Comment, *comment.Reply, error)
}
