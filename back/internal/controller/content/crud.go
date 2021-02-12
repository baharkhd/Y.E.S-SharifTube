package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/controller"
	"yes-sharifTube/internal/model/content"
)

func (c *contentController) GetContent(contentID string) (*model.Content, error) {
	cID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Get(cID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *contentController) GetContents(tags []string, courseID *string, startIdx, amount int) ([]*model.Content, error) {
	var cpID *primitive.ObjectID = nil
	var cID primitive.ObjectID
	var err error
	if courseID != nil {
		cID, err = primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return nil, &model.InternalServerException{Message: err.Error()}
		}
		cpID = &cID
	}
	contents, err := c.dbDriver.GetAll(cpID, tags, startIdx, amount)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return content.ReshapeAll(contents)
}

func (c *contentController) CreateContent(authorUsername, courseID, title string, description *string, vurl string, tags []string) (*model.Content, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Insert(authorUsername, cID, title, vurl, description, tags)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *contentController) UpdateContent(authorUsername, courseID, contentID string, newTitle, newDescription *string, newTags []string) (*model.Content, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cnID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Update(authorUsername, cID, cnID, newTitle, newDescription, newTags)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}

func (c *contentController) DeleteContent(authorUsername, courseID, contentID string) (*model.Content, error) {
	crID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	coID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, err := c.dbDriver.Delete(authorUsername, crID, coID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return cr.Reshape()
}
