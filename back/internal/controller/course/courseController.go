package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/controller"
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

func (c *courseController) AddUserToCourse(username, courseID, token string) (*model.Course, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.AddUser(username, token, cID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *courseController) DeleteUserFromCourse(username, courseID, targetUsername string) (*model.Course, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.DeleteUser(username, targetUsername, cID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *courseController) PromoteUserToTA(username, courseID, targetUsername string) (*model.Course, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.PromoteToTA(username, targetUsername, cID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *courseController) DemoteUserToSTD(username, courseID, targetUsername string) (*model.Course, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.DemoteToSTD(username, targetUsername, cID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}
