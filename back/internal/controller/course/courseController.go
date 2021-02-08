package controller

import (
	"yes-sharifTube/graph/model"
	"yes-sharifTube/pkg/database"
)

type courseController struct {
	dbDriver database.CourseDBDriver
}

var coursec *courseController

func init() {
	coursec = &courseController{}
}

func GetCourseController() *courseController {
	return coursec
}

func (c *courseController) SetDBDriver(dbDriver database.CourseDBDriver) {
	coursec.dbDriver = dbDriver
}

func (c *courseController) AddUserToCourse(username, courseId, token string) (*model.Course, error) {
	panic("not implemented")
}

func (c *courseController) DeleteUserFromCourse(username, courseId, targetUsername string) (*model.Course, error) {
	panic("not implemented")
}

func (c *courseController) PromoteUserToTA(username, courseID, targetUsername string) (*model.Course, error) {
	panic("not implemented")
}

func (c *courseController) DemoteUserToSTD(username, courseID, targetUsername string) (*model.Course, error) {
	panic("not implemented")
}
