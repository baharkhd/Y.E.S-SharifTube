package course

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
)

func Get(courseID string) (*Course, error) {
	objID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	course, err := DBD.Get(objID)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func GetAll(courseIDs []string) ([]*Course, error) {
	var courses []*Course
	for _, cID := range courseIDs {
		objID, err := primitive.ObjectIDFromHex(cID)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		course, err := DBD.Get(objID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func GetByFilter(keywords []string, start, amount int) ([]*Course, error) {
	courses, err := DBD.GetByFilter(keywords, start, amount)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func Insert(course *Course) (*Course, error) {
	course, err := DBD.Insert(course)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func Update(username string, course *Course) error {
	if !course.IsUserProfessor(username) {
		return model.UserNotAllowedException{Message: "you can't change this course. because you are not professor"}
	}
	return DBD.UpdateInfo(course.ID, course.Title, course.Summery, course.Token)
}

func Delete(username string, course *Course) error {
	if !course.IsUserProfessor(username) {
		return model.UserNotAllowedException{Message: "you are not professor"}
	}
	return DBD.Delete(course.ID, append(course.StdUns, append(course.TaUns, course.ProfUn)...))
}

func AddUser(username, token string, course *Course) (*Course, error) {
	if course.IsUserParticipateInCourse(username) {
		return nil, model.DuplicateUsernameException{Message: "you been been added before"}
	}
	if !course.CheckCourseToken(token) {
		return nil, model.IncorrectTokenException{Message: "wrong course token"}
	}
	course.StdUns = append(course.StdUns, username)
	if err := DBD.AddStd(username, course.ID); err != nil {
		return nil, err
	}
	return course, nil
}

func DeleteUser(username, targetUsername string, course *Course) (*Course, error) {
	if !course.IsUserParticipateInCourse(targetUsername) {
		return nil, model.UserNotFoundException{Message: "you weren't participate in course"}
	}
	if !course.IsUserAllowedToDeleteUser(username, targetUsername) {
		return nil, model.UserNotAllowedException{Message: "you can't remove this user"}
	}
	if modelUtil.ContainsInStringArray(course.TaUns, targetUsername) {
		course.TaUns = modelUtil.RemoveFromStringArray(course.TaUns, targetUsername)
		if err := DBD.DelTa(username, course.ID); err != nil {
			return nil, err
		}
	} else {
		course.StdUns = modelUtil.RemoveFromStringArray(course.StdUns, targetUsername)
		if err := DBD.DelStd(username, course.ID); err != nil {
			return nil, err
		}
	}
	return course, nil
}

func PromoteUser(username, targetUsername string, course *Course) (*Course, error) {
	if !course.IsUserProfOrTA(username) {
		return nil, model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	if !course.IsUserStudent(targetUsername) {
		return nil, model.UserIsNotSTDException{Message: "you are not student"}
	}
	course.StdUns = modelUtil.RemoveFromStringArray(course.StdUns, targetUsername)
	course.TaUns = append(course.TaUns, targetUsername)
	if err := DBD.PromoteDemoteUser(course); err != nil {
		return nil, err
	}
	return course, nil
}

func DemoteUser(username, targetUsername string, course *Course) (*Course, error) {
	if !course.IsUserProfOrTA(username) {
		return nil, model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	if !course.IsUserTA(targetUsername) {
		return nil, model.UserIsNotSTDException{Message: "you are not ta"}
	}
	course.TaUns = modelUtil.RemoveFromStringArray(course.TaUns, targetUsername)
	course.StdUns = append(course.StdUns, targetUsername)
	if err := DBD.PromoteDemoteUser(course); err != nil {
		return nil, err
	}
	return course, nil
}
