package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/controller"
	modelUtil "yes-sharifTube/internal/model"
	"yes-sharifTube/internal/model/course"
)

func (c *courseController) GetCourses(courseIDs []string) ([]*model.Course, error) {
	objIDs, err := modelUtil.ConvertStringsToObjectIDs(courseIDs)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	courses, err := c.dbDriver.GetAll(objIDs)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return course.ReshapeAll(courses)
}

func (c *courseController) GetCoursesByKeyWords(keywords []string, startIdx, amount int) ([]*model.Course, error) {
	courses, err := c.dbDriver.GetByFilter(keywords, startIdx, amount)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return course.ReshapeAll(courses)
}

func (c *courseController) CreateCourse(authorUsername, title string, summery *string, token string) (*model.Course, error) {
	cr, err := c.dbDriver.Insert(authorUsername, title, token, summery)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *courseController) UpdateCourse(authorUsername, courseID string, newTitle, newSummery, newToken *string) (*model.Course, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Update(authorUsername, cID, newTitle, newToken, newSummery)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *courseController) DeleteCourse(authorUsername, courseID string) (*model.Course, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Delete(authorUsername, cID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}
