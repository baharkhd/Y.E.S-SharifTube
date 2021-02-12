package controller

import (
	"yes-sharifTube/pkg/database"
)

type commentController struct {
	dbDriver database.CommentDBDriver
}

var commentc *commentController

func init() {
	commentc = &commentController{}
}

func GetCommentController() *commentController {
	return commentc
}

func (c *commentController) SetDBDriver(dbDriver database.CommentDBDriver) {
	commentc.dbDriver = dbDriver
}
