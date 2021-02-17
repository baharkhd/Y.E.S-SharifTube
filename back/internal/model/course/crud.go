package course

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
)

func Get(courseID string) (*Course, error) {

	// checking to be in cache first
	c, err := GetFromCache(courseID)
	if err == nil {
		return c, nil
	}

	// if not exists, get from database
	objID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	course, err := DBD.Get(objID)
	if err != nil {
		return nil, err
	}

	// add the course to cache
	_ = course.Cache()

	return course, nil
}

func GetAll(courseIDs []string) ([]*Course, error) {
	var courses []*Course
	for _, cID := range courseIDs {

		// checking to be in cache first
		course, err := GetFromCache(cID)
		if err != nil {

			// if not exists, get from database
			objID, err := primitive.ObjectIDFromHex(cID)
			if err != nil {
				return nil, model.InternalServerException{Message: err.Error()}
			}
			course, err = DBD.Get(objID)
			if err != nil {
				return nil, err
			}

			// add the course to cache
			_ = course.Cache()

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

	// create the bucket
	//OSD.NewBucket(OSD.GetRoot(),course.ID.Hex())

	// add the course to cache
	_ = course.Cache()

	return course, nil
}

func Update(course *Course) error {
	err := DBD.UpdateInfo(course.ID, course.Title, course.Summery, course.Token)
	if err != nil {
		return err
	}

	// add the course to cache
	DeleteFromCache(course.ID.Hex())
	_ = course.Cache()

	return nil
}

func Delete(course *Course) error {
	err := DBD.Delete(course.ID, append(course.StdUns, append(course.TaUns, course.ProfUn)...))
	if err != nil {
		return err
	}

	// delete from cache if exists
	DeleteFromCache(course.ID.Hex())

	return nil
}

func AddUser(username string, course *Course) (*Course, error) {
	course.StdUns = append(course.StdUns, username)
	if err := DBD.AddStd(username, course.ID); err != nil {
		return nil, err
	}

	// update the course in cache if exists
	course.UpdateCache()

	return course, nil
}

func DeleteUser(targetUsername string, course *Course) (*Course, error) {
	if modelUtil.ContainsInStringArray(course.TaUns, targetUsername) {
		course.TaUns = modelUtil.RemoveFromStringArray(course.TaUns, targetUsername)
		if err := DBD.DelTa(targetUsername, course.ID); err != nil {
			return nil, err
		}
	} else {
		course.StdUns = modelUtil.RemoveFromStringArray(course.StdUns, targetUsername)
		if err := DBD.DelStd(targetUsername, course.ID); err != nil {
			return nil, err
		}
	}

	// update the course in cache if exists
	course.UpdateCache()

	return course, nil
}

func PromoteUser(targetUsername string, course *Course) (*Course, error) {
	course.StdUns = modelUtil.RemoveFromStringArray(course.StdUns, targetUsername)
	course.TaUns = append(course.TaUns, targetUsername)
	if err := DBD.PromoteDemoteUser(course); err != nil {
		return nil, err
	}

	// update the course in cache if exists
	course.UpdateCache()

	return course, nil
}

func DemoteUser(targetUsername string, course *Course) (*Course, error) {
	course.TaUns = modelUtil.RemoveFromStringArray(course.TaUns, targetUsername)
	course.StdUns = append(course.StdUns, targetUsername)
	if err := DBD.PromoteDemoteUser(course); err != nil {
		return nil, err
	}

	// update the course in cache if exists
	course.UpdateCache()

	return course, nil
}
