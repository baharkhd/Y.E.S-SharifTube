package controller

import (
	"github.com/99designs/gqlgen/graphql"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/user"
)

func GetContent(contentID string) (*content.Content, error) {
	cn, err := content.Get(nil, contentID)
	if err != nil {
		return nil, err
	}
	return cn, nil
}

func GetContents(tags []string, courseID *string, startIdx, amount int) ([]*content.Content, error) {
	contents, err := content.GetByFilter(courseID, tags, startIdx, amount)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func CreateContent(authorUsername, courseID, title string, description *string, upload graphql.Upload, tags []string) (*content.Content, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}

	// create new course
	cn, err := cr.AddNewContent(authorUsername, title, description, upload, tags)
	if err!=nil{
		return nil, err
	}

	return cn, nil
}

func UpdateContent(authorUsername, courseID, contentID string, newTitle, newDescription *string, newTags []string) (*content.Content, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get the content & course from database
	cn, err := content.Get(&courseID, contentID)
	if err != nil {
		return nil, err
	}
	//check if user can update content
	err = cr.IsUserAllowedToUpdateContent(authorUsername,cn)
	if err != nil {
		return nil, err
	}
	// update the content
	err = cn.Update(newTitle, newDescription, newTags)
	if err != nil {
		return nil, err
	}
	// update the content in database
	if err = content.Update(courseID, cn); err != nil {
		return nil, err
	}
	return cn, nil
}

func DeleteContent(authorUsername, courseID, contentID string) (*content.Content, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get the content & course from database
	cn, err := content.Get(&courseID, contentID)
	if err != nil {
		return nil, err
	}
	//check if user can delete content
	err = cr.IsUserAllowedToDeleteContent(authorUsername,cn)
	if err != nil {
		return nil, err
	}
	// delete the content from database
	if err = content.Delete(courseID, cn); err != nil {
		return nil, err
	}
	return cn, nil
}
