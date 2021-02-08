package controller

import (
	"yes-sharifTube/pkg/database"
	"yes-sharifTube/graph/model"
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

func (p *pendingController) AcceptPending(username, courseID, pendingID, newTitle, newDescription string) (*model.Pending, error) {
	panic("not implemented")
}

func (p *pendingController) RejectPending(username, courseID, pendingID string) (*model.Pending, error) {
	panic("not implemented")
}
