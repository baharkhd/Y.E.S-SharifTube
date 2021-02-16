package controller

import (
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/user"
)

func CreateAttachment(authorUsername, courseID, name string, description *string, aurl string) (*attachment.Attachment, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// check if user can insert attachment
	err = cr.IsUserAllowedToInsertAttachment(authorUsername)
	if err != nil {
		return nil, err
	}
	// create an attachment
	an, err := attachment.New(name, aurl, courseID, description)
	if err != nil {
		return nil, err
	}
	// insert the attachment into database
	an, err = attachment.Insert(courseID, an)
	if err != nil {
		return nil, err
	}
	// maintain consistency in cache
	cr.AddAttachment(an)
	cr.UpdateCache()
	return an, nil
}

func UpdateAttachment(authorUsername, courseID, attachmentID string, newName, newDescription *string) (*attachment.Attachment, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get the attachment from database
	an, err := attachment.Get(&courseID, attachmentID)
	if err != nil {
		return nil, err
	}
	// check if user can update attachment
	err = cr.IsUserAllowedToUpdateAttachment(authorUsername)
	if err != nil {
		return nil, err
	}
	// update the attachment
	err = an.Update(newName, newDescription)
	if err != nil {
		return nil, err
	}
	// update the attachment in database
	err = attachment.Update(courseID, an)
	if err != nil {
		return nil, err
	}
	// maintain consistency in cache
	cr.UpdateAttachment(an)
	cr.UpdateCache()
	return an, nil
}

func DeleteAttachment(authorUsername, courseID, attachmentID string) (*attachment.Attachment, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get the attachment from database
	an, err := attachment.Get(&courseID, attachmentID)
	if err != nil {
		return nil, err
	}
	// check if user can delete attachment
	err = cr.IsUserAllowedToDeleteAttachment(authorUsername)
	if err != nil {
		return nil, err
	}
	// delete the attachment from database
	err = attachment.Delete(courseID, an)
	if err != nil {
		return nil, err
	}
	// maintain consistency in cache
	cr.DeleteAttachment(an.ID)
	cr.UpdateCache()
	return an, nil
}
