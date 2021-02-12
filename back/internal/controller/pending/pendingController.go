package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/controller"
	"yes-sharifTube/pkg/database"
)

type pendingController struct {
	dbDriver database.PendingDBDriver
}

var pendingc *pendingController

func init() {
	pendingc = &pendingController{}
}

func GetPendingController() *pendingController {
	return pendingc
}

func (p *pendingController) SetDBDriver(dbDriver database.PendingDBDriver) {
	pendingc.dbDriver = dbDriver
}

func (p *pendingController) AcceptPending(username, courseID, pendingID string, newTitle, newDescription *string) (*model.Pending, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pID, err := primitive.ObjectIDFromHex(pendingID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pr, err := p.dbDriver.Accept(username, cID, pID, newTitle, newDescription)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return pr.Reshape()
}

func (p *pendingController) RejectPending(username, courseID, pendingID string) (*model.Pending, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pID, err := primitive.ObjectIDFromHex(pendingID)
	if err != nil {
		return nil, &model.InternalServerException{Message: err.Error()}
	}
	pr, err := p.dbDriver.Reject(username, cID, pID)
	err = controller.CastDBExceptionToGQLException(err)
	if err != nil {
		return nil, err
	}
	return pr.Reshape()
}
