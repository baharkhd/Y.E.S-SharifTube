package controller

import (
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/user"
)

func GetCourses(courseIDs []string) ([]*course.Course, error) {
	// get courses from database
	courses, err := course.GetAll(courseIDs)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func GetCoursesByKeyWords(keywords []string, startIdx, amount int) ([]*course.Course, error) {
	// get courses from database
	courses, err := course.GetByFilter(keywords, startIdx, amount)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func CreateCourse(authorUsername, title string, summery *string, token string) (*course.Course, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// create a user
	cr, err := course.New(title, authorUsername, token, summery)
	if err != nil {
		return nil, err
	}
	// insert the user into database
	cr, err = course.Insert(cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func UpdateCourse(authorUsername, courseID string, newTitle, newSummery, newToken *string) (*course.Course, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// update the course
	err = cr.Update(newTitle, newSummery, newToken)
	if err != nil {
		return nil, err
	}
	// update the course in database
	if err = course.Update(authorUsername, cr); err != nil {
		return nil, err
	}
	return cr, nil
}

func DeleteCourse(authorUsername, courseID string) (*course.Course, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// delete the course from database
	if err = course.Delete(authorUsername, cr); err != nil {
		return nil, err
	}
	return cr, nil
}

func AddUserToCourse(username, courseID, token string) (*course.Course, error) {
	// check if user exists in database
	if _, err := user.Get(username); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// add the user to course in database
	cr, err = course.AddUser(username, token, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func DeleteUserFromCourse(username, courseID, targetUsername string) (*course.Course, error) {
	// check if user exists in database
	if _, err := user.Get(username); err != nil {
		return nil, err
	}
	if _, err := user.Get(targetUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// delete the user from course in database
	cr, err = course.DeleteUser(username, targetUsername, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func PromoteUserToTA(username, courseID, targetUsername string) (*course.Course, error) {
	// check if user exists in database
	if _, err := user.Get(username); err != nil {
		return nil, err
	}
	if _, err := user.Get(targetUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// promote user to ta in database
	cr, err = course.PromoteUser(username, targetUsername, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func DemoteUserToSTD(username, courseID, targetUsername string) (*course.Course, error) {
	// check if user exists in database
	if _, err := user.Get(username); err != nil {
		return nil, err
	}
	if _, err := user.Get(targetUsername); err != nil {
		return nil, err
	}
	// get the course from database
	cr, err := course.Get(courseID)
	if err != nil {
		return nil, err
	}
	// demote at to student in database
	cr, err = course.DemoteUser(username, targetUsername, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}
