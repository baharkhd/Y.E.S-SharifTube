package controller

import "yes-sharifTube/graph/model"

func (a *attachmentController) CreateAttachment(authorUsername, courseID, name, description, aurl string) (*model.Attachment, error) {
	panic("not implemented")
}

func (a *attachmentController) UpdateAttachment(authorUsername, courseID, attachmentID, newName, newDescription string) (*model.Attachment, error) {
	panic("not implemented")
}

func (a *attachmentController) DeleteAttachment(authorUsername, courseID, attachmentID string) (*model.Attachment, error) {
	panic("not implemented")
}
