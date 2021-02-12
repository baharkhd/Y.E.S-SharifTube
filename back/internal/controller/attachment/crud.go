package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/controller"
)

func (a *attachmentController) CreateAttachment(authorUsername, courseID, name string, description *string, aurl string) (*model.Attachment, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	ar, err := a.dbDriver.Insert(authorUsername, cID, name, aurl, description)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return ar.Reshape(), nil
}

func (a *attachmentController) UpdateAttachment(authorUsername, courseID, attachmentID string, newName, newDescription *string) (*model.Attachment, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	aID, err := primitive.ObjectIDFromHex(attachmentID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	ar, err := a.dbDriver.Update(authorUsername, cID, aID, newName, newDescription)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return ar.Reshape(), nil
}

func (a *attachmentController) DeleteAttachment(authorUsername, courseID, attachmentID string) (*model.Attachment, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	aID, err := primitive.ObjectIDFromHex(attachmentID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	ar, err := a.dbDriver.Delete(authorUsername, cID, aID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return ar.Reshape(), nil
}
