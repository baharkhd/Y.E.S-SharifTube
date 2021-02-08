package controller

import "yes-sharifTube/graph/model"

func (c *courseController) GetCourses(courseIDs []*string) (*model.Course, error) {
	panic("not implemented")
}

func (c *courseController) GetCoursesByKeyWords(keywords []*string, startIdx, amount int) ([]*model.Course, error) {
	panic("not implemented")
}

func (c *courseController) CreateCourse(authorUsername, title, summery, token string) (*model.Course, error) {
	panic("not implemented")
}

func (c *courseController) UpdateCourse(authorUsername, courseID, newTitle, newSummery, newToken string) (*model.Course, error) {
	panic("not implemented")
}

func (c *courseController) DeleteCourse(authorUsername, courseID string) (*model.Course, error) {
	panic("not implemented")
}