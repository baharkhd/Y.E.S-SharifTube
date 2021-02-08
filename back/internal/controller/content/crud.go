package controller

import "yes-sharifTube/graph/model"

func (c *contentController) GetContent(contentID string) (*model.Content, error) {
	panic("not implemented")
}

func (c *contentController) GetContents(tags []*string, startIdx, amount int) ([]*model.Content, error) {
	panic("not implemented")
}

func (c *contentController) CreateContent(authorUsername, courseID, title, description, vurl string, tags []string) (*model.Content, error) {
	panic("not implemented")
}

func (c *contentController) UpdateContent(authorUsername, courseID, contentID, newTitle, newDescription string, newTags []string) (*model.Content, error) {
	panic("not implemented")
}

func (c *contentController) DeleteContent(authorUsername, courseID, contentID string) (*model.Content, error) {
	panic("not implemented")
}
