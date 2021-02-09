package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/controller"
	"yes-sharifTube/internal/model/pending"
)

func (p *pendingController) GetPendings(courseID, uploaderUsername *string, status *model.Status, startIdx, amount int) ([]*model.Pending, error) {
	var cpID *primitive.ObjectID = nil
	var cID primitive.ObjectID
	var spt *pending.Status = nil
	var st pending.Status
	var err error
	if courseID != nil {
		cID, err = primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return nil, &model.InternalServerException{Message: err.Error()}
		}
		cpID = &cID
	}
	if status != nil {
		st = pending.NewStatus(*status)
		spt = &st
	}
	pr, err := p.dbDriver.GetByFilter(cpID, spt, uploaderUsername, startIdx, amount)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return pending.ReshapeAll(pr)
}

func (p *pendingController) CreatePending(authorUsername, courseID, title, description, furl string) (*model.Pending, error) {
	np, err := pending.New(title, description, authorUsername, furl, courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pr, err := p.dbDriver.Insert(authorUsername, cID, np)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return pr.Reshape()
}

func (p *pendingController) UpdatePending(authorUsername, courseID, pendingID, newTitle, newDescription string) (*model.Pending, error) {
	np, err := pending.New(newTitle, newDescription, authorUsername, "", courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	np.ID, err = primitive.ObjectIDFromHex(pendingID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pr, err := p.dbDriver.Update(authorUsername, cID, np)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return pr.Reshape()
}

func (p *pendingController) DeletePending(authorUsername, courseID, pendingID string) (*model.Pending, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pID, err := primitive.ObjectIDFromHex(pendingID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pr, err := p.dbDriver.Delete(authorUsername, cID, pID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return pr.Reshape()
}
