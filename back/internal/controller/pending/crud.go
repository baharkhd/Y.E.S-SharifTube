package controller

import (
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/pending"
	"yes-sharifTube/internal/model/user"
)

func GetPendings(username *string, courseID, uploaderUsername *string, status *model.Status, startIdx, amount int) ([]*pending.Pending, error) {
	prs, err := pending.GetByFilter(courseID, uploaderUsername, status, startIdx, amount)
	if err != nil {
		return nil, err
	}
	var fpr []*pending.Pending
	for _, pr := range prs {
		c, _ := course.Get(pr.CourseID)
		fpr = append(fpr, c.FilterPending(username, pr))
	}
	return prs, nil
}

func CreatePending(authorUsername, courseID, title string, description *string, furl string) (*pending.Pending, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// create a pending
	pn, err := pending.New(title, authorUsername, furl, courseID, description)
	if err != nil {
		return nil, err
	}
	// check if user can offer
	err = cr.IsUserAllowedToInsertPending(authorUsername)
	if err != nil {
		return nil, err
	}
	// insert the pending into database
	pn, err = pending.Insert(courseID, pn)
	if err != nil {
		return nil, err
	}
	return pn, nil
}

func UpdatePending(authorUsername, courseID, pendingID string, newTitle, newDescription *string) (*pending.Pending, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get pending from database
	pn, err := pending.Get(&courseID, pendingID)
	if err != nil {
		return nil, err
	}
	// update the pending
	err = pn.Update(newTitle, newDescription)
	if err != nil {
		return nil, err
	}
	// check if user can update offer
	err = cr.IsUserAllowedToUpdatePending(authorUsername, pn)
	if err != nil {
		return nil, err
	}
	// update the pending in database
	err = pending.Update(courseID, pn)
	if err != nil {
		return nil, err
	}
	return pn, nil
}

func DeletePending(authorUsername, courseID, pendingID string) (*pending.Pending, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get pending from database
	pn, err := pending.Get(&courseID, pendingID)
	if err != nil {
		return nil, err
	}
	// check if user can delete offer
	err = cr.IsUserAllowedToDeletePending(authorUsername, pn)
	if err != nil {
		return nil, err
	}
	// delete the pending from database
	err = pending.Delete(courseID, pn)
	if err != nil {
		return nil, err
	}
	return pn, nil
}

func AcceptPending(username, courseID, pendingID string, newTitle, newDescription *string) (*pending.Pending, error) {
	// check if user exists in database
	if _, err := user.Get(username); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get pending from database
	pn, err := pending.Get(&courseID, pendingID)
	if err != nil {
		return nil, err
	}
	// check if user can accept offer
	err = cr.IsUserAllowedToAcceptPending(username, pn)
	if err != nil {
		return nil, err
	}
	// update the pending
	err = pn.Update(newTitle, newDescription)
	if err != nil {
		return nil, err
	}
	// update the pending in database
	pn, err = pending.Accept(courseID, pn)
	if err != nil {
		return nil, err
	}
	// accept that pending into content
	nc, err := content.New(pn.Title, pn.UploadedByUn, pn.Furl, pn.CourseID, &pn.Description, &username, nil)
	if err != nil {
		return nil, err
	}
	_, err = content.Insert(courseID, nc)
	if err != nil {
		return nil, err
	}
	return pn, nil
}

func RejectPending(username, courseID, pendingID string) (*pending.Pending, error) {
	// check if user exists in database
	if _, err := user.Get(username); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// get pending from database
	pn, err := pending.Get(&courseID, pendingID)
	if err != nil {
		return nil, err
	}
	// check if user can reject offer
	err = cr.IsUserAllowedToRejectPending(username, pn)
	if err != nil {
		return nil, err
	}
	// reject the pending in database
	pn, err = pending.Reject(courseID, pn)
	if err != nil {
		return nil, err
	}
	return pn, nil
}
