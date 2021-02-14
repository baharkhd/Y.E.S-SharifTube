package course

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBDriver interface {
	Get(courseIDs primitive.ObjectID) (*Course, error)
	GetByFilter(keywords []string, start, amount int) ([]*Course, error)
	Insert(course *Course) (*Course, error)
	UpdateInfo(courseID primitive.ObjectID, title, summery, token string) error
	Delete(courseID primitive.ObjectID, members []string) error
	AddStd(username string, courseID primitive.ObjectID) error
	DelStd(username string, courseID primitive.ObjectID) error
	DelTa(username string, courseID primitive.ObjectID) error
	PromoteDemoteUser(course *Course) error
}
