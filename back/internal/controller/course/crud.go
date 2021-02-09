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

func (c *courseController) CreateCourse(authorUsername, title, summery, token string) (*model.Course, error) {
	nc, err := course.New(title, summery, authorUsername, token)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Insert(authorUsername, nc)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *courseController) UpdateCourse(authorUsername, courseID, newTitle, newSummery, newToken string) (*model.Course, error) {
	nc, err := course.New(newTitle, newSummery, authorUsername, newToken)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	nc.ID, err = primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Update(authorUsername, nc)
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
