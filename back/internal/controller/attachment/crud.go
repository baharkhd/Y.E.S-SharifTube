package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal"
	"yes-sharifTube/internal/model/attachment"
)

func (a *attachmentController) CreateAttachment(authorUsername, courseID, name, description, aurl string) (*model.Attachment, error) {
	an, err := attachment.New(name, description, aurl, courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	ar, err := a.dbDriver.Insert(authorUsername, cID, an)
	err = internal.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return ar.Reshape(), nil
}

func (a *attachmentController) UpdateAttachment(authorUsername, courseID, attachmentID, newName, newDescription string) (*model.Attachment, error) {
	an, err := attachment.New(newName, newDescription, "", courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	an.ID, err = primitive.ObjectIDFromHex(attachmentID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	ar, err := a.dbDriver.Update(authorUsername, cID, an)
	err = internal.CastDBExceptionToGQLException(err)
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
	err = internal.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return ar.Reshape(), nil
}
