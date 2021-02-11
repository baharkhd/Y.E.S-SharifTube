package controller

import (
	"yes-sharifTube/pkg/database"
)

type contentController struct {
	dbDriver database.ContentDBDriver
}

var contentc *contentController

func init() {
	contentc = &contentController{}
}

func GetContentController() *contentController {
	return contentc
}

func (c *contentController) SetDBDriver(dbDriver database.ContentDBDriver) {
	contentc.dbDriver = dbDriver
}

func (c *contentController) GetDBDriver() database.ContentDBDriver {
	return c.dbDriver
}
