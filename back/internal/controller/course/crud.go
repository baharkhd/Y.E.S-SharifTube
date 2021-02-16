package controller

import (
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/user"
)

func GetCourses(username *string, courseIDs []string) ([]*course.Course, error) {
	// get courses from database
	courses, err := course.GetAll(courseIDs)
	if err != nil {
		return nil, err
	}
	return course.FilterPendsOfCourses(username, courses), nil
}

func GetCoursesByKeyWords(username *string, keywords []string, startIdx, amount int) ([]*course.Course, error) {
	// get courses from database
	courses, err := course.GetByFilter(keywords, startIdx, amount)
	if err != nil {
		return nil, err
	}
	return course.FilterPendsOfCourses(username, courses), nil
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
	return cr.FilterPendsOfCourse(&authorUsername), nil
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
	// check if user can update course
	if err = cr.IsUserAllowedToUpdateCourse(authorUsername); err != nil{
		return nil, err
	}
	// update the course in database
	if err = course.Update(cr); err != nil {
		return nil, err
	}
	return cr.FilterPendsOfCourse(&authorUsername), nil
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
	// check if user can delete course
	if err = cr.IsUserAllowedToDeleteCourse(authorUsername); err != nil{
		return nil, err
	}
	// delete the course from database
	if err = course.Delete(cr); err != nil {
		return nil, err
	}
	return cr.FilterPendsOfCourse(&authorUsername), nil
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
	// check if user can add to the course
	if err = cr.IsUserAllowedToAddUserInCourse(username, token); err != nil{
		return nil, err
	}
	// add the user to course in database
	cr, err = course.AddUser(username, cr)
	if err != nil {
		return nil, err
	}
	return cr.FilterPendsOfCourse(&username), nil
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
	// check if user can delete another from course
	if err = cr.IsUserAllowedToDeleteUserInCourse(username, targetUsername); err != nil{
		return nil, err
	}
	// delete the user from course in database
	cr, err = course.DeleteUser(targetUsername, cr)
	if err != nil {
		return nil, err
	}
	return cr.FilterPendsOfCourse(&username), nil
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
	// check if user can promote another in course
	if err = cr.IsUserAllowedToPromoteUserInCourse(username, targetUsername); err != nil{
		return nil, err
	}
	// promote user to ta in database
	cr, err = course.PromoteUser(targetUsername, cr)
	if err != nil {
		return nil, err
	}
	return cr.FilterPendsOfCourse(&username), nil
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
	// check if user can demote to student course
	if err = cr.IsUserAllowedToDemoteUserInCourse(username, targetUsername); err != nil{
		return nil, err
	}
	// demote at to student in database
	cr, err = course.DemoteUser(targetUsername, cr)
	if err != nil {
		return nil, err
	}
	return cr.FilterPendsOfCourse(&username), nil
}
