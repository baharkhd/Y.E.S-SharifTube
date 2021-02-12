package controller

import (
	"yes-sharifTube/pkg/database"
)

type attachmentController struct {
	dbDriver database.AttachmentDBDriver
}

var ac *attachmentController

func init() {
	ac = &attachmentController{}
}

func GetAttachmentController() *attachmentController {
	return ac
}

func (a *attachmentController) SetDBDriver(dbDriver database.AttachmentDBDriver) {
	ac.dbDriver = dbDriver
}
