package controller

import "yes-sharifTube/graph/model"

func (p *pendingController) GetPendings(courseID, uploaderUsername string, status model.Status, startIdx, amount int) ([]*model.Pending, error) {
	panic("not implemented")
}

func (p *pendingController) CreatePending(authorUsername, courseID, title, description, furl string) (model.Pending, error) {
	panic("not implemented")
}

func (p *pendingController) UpdatePending(authorUsername, courseID, pendingID, newTitle, newDescription string) (model.Pending, error) {
	panic("not implemented")
}

func (p *pendingController) DeletePending(authorUsername, courseID, pendingID string) (model.Pending, error) {
	panic("not implemented")
}
