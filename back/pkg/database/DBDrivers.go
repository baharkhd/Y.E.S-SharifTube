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
	Insert(username string, course *course.Course) (*course.Course, error)
	Update(username string, course *course.Course) (*course.Course, error)
	Delete(username string, courseID primitive.ObjectID) (*course.Course, error)
	AddUser(username, token string, courseID primitive.ObjectID) (*course.Course, error)
	DeleteUser(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error)
	PromoteToTA(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error)
	DemoteToSTD(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error)
}

type PendingDBDriver interface {
	GetByFilter(courseID primitive.ObjectID, status pending.Status, uploaderUsername string, start, amount int) ([]*pending.Pending, error)
	Insert(username string, courseID primitive.ObjectID, pending *pending.Pending) (*pending.Pending, error)
	Update(username string, courseID primitive.ObjectID, pending *pending.Pending) (*pending.Pending, error)
	Delete(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error)
	Accept(username string, courseID primitive.ObjectID, pending *pending.Pending) (*pending.Pending, error)
	Reject(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error)
}

type ContentDBDriver interface {
	Get(contentID primitive.ObjectID) (*content.Content, error)
	GetAll(courseID primitive.ObjectID, tags []string, start, amount int) ([]*content.Content, error)
	Insert(username string, courseID primitive.ObjectID, content *content.Content) (*content.Content, error)
	Update(username string, courseID primitive.ObjectID, content *content.Content) (*content.Content, error)
	Delete(username string, courseID, contentID primitive.ObjectID) (*content.Content, error)
}

type AttachmentDBDriver interface {
	Insert(username string, courseID primitive.ObjectID, attachment *attachment.Attachment) (*attachment.Attachment, error)
	Update(username string, courseID primitive.ObjectID, attachment *attachment.Attachment) (*attachment.Attachment, error)
	Delete(username string, courseID, attachmentID primitive.ObjectID) (*attachment.Attachment, error)
}

type CommentDBDriver interface {
	Insert(username string, courseID, repliedAtID primitive.ObjectID, comment *comment.Comment) (*comment.Comment, error)
	Update(username string, courseID primitive.ObjectID, comment *comment.Comment) (*comment.Comment, error)
	Delete(username string, courseID, commentID primitive.ObjectID) (*comment.Comment, error)
}
