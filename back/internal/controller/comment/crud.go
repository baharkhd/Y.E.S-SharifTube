package controller

import "yes-sharifTube/graph/model"

func (c *commentController) CreateComment(authorUsername, contentID, repliedID, body string) (*model.Comment, error) {
	panic("not implemented")
}

func (c *commentController) UpdateComment(authorUsername, contentID, commentID, newBody string) (*model.Comment, error) {
	panic("not implemented")
}

func (c *commentController) DeleteComment(authorUsername, contentID, commentID string) (*model.Comment, error) {
	panic("not implemented")
}

