package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"yes-sharifTube/internal/model/course"
)

type CourseMongoDriver struct {
	collection *mongo.Collection
}

func (c CourseMongoDriver) GetAll(courseIDs []primitive.ObjectID) ([]*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) GetByFilter(keywords []string, start, amount int) ([]*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) Insert(username string, course *course.Course) (*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) Update(username string, course *course.Course) (*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) Delete(username string, courseID primitive.ObjectID) (*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) AddUser(username, token string, courseID primitive.ObjectID) (*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) DeleteUser(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) PromoteToTA(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error) {
	panic("not implemented")
}
func (c CourseMongoDriver) DemoteToSTD(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error) {
	panic("not implemented")
}

func NewCourseMongoDriver(db, collection string) *CourseMongoDriver {
	return &CourseMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
