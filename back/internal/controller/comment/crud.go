package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/controller"
)

func (c *commentController) CreateComment(authorUsername, contentID, body string, repliedID *string) (*model.Comment, *model.Reply, error) {
	cID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, nil, &model.InternalServerException{Message: err.Error()}
	}
	var rpID *primitive.ObjectID = nil
	var rID primitive.ObjectID
	if repliedID != nil {
		rID, err = primitive.ObjectIDFromHex(*repliedID)
		if err != nil {
			return nil, nil, &model.InternalServerException{Message: err.Error()}
		}
		rpID = &rID
	}
	cr, rr, err := c.dbDriver.Insert(authorUsername, cID, rpID, body)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, nil, err
	}
	if cr != nil {
		fc, err := cr.Reshape()
		return fc, nil, err
	} else {
		fc, err := rr.Reshape()
		return nil, fc, err
	}
}

func (c *commentController) UpdateComment(authorUsername, contentID, commentID string, newBody *string) (*model.Comment, *model.Reply, error) {
	cID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, nil, &model.InternalServerException{Message: err.Error()}
	}
	cmID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, rr, err := c.dbDriver.Update(authorUsername, cID, cmID, newBody)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, nil, err
	}
	if cr != nil {
		fc, err := cr.Reshape()
		return fc, nil, err
	} else {
		fc, err := rr.Reshape()
		return nil, fc, err
	}
}

func (c *commentController) DeleteComment(authorUsername, contentID, commentID string) (*model.Comment, *model.Reply, error) {
	coID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, nil, &model.InternalServerException{Message: err.Error()}
	}
	cnID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, nil, &model.InternalServerException{Message: err.Error()}
	}
	cr, rr, err := c.dbDriver.Delete(authorUsername, coID, cnID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, nil, err
	}
	if cr != nil {
		fc, err := cr.Reshape()
		return fc, nil, err
	} else {
		fc, err := rr.Reshape()
		return nil, fc, err
	}
}