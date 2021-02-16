package content

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
)

func Get(courseID *string, contentID string) (*Content, error) {

	// checking to be in cache first
	c, err := GetFromCache(contentID)
	if err == nil {
		return c, nil
	}

	// if not exists, get from database
	var cpID *primitive.ObjectID = nil
	if courseID != nil {
		cID, err := primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		cpID = &cID
	}
	conID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	content, err := DBD.Get(cpID, conID)
	if err != nil {
		return nil, err
	}

	// add the content to cache
	_ = content.Cache()

	return content, nil
}

func GetByFilter(courseID *string, tags []string, startIdx, amount int) ([]*Content, error) {
	var cpID *primitive.ObjectID = nil
	var cID primitive.ObjectID
	var err error
	if courseID != nil {
		cID, err = primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		cpID = &cID
	}
	contents, err := DBD.GetAll(cpID, tags, startIdx, amount)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func insert(courseID string, content *Content) (*Content, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	content, err = DBD.Insert(cID, content)
	if err != nil {
		return nil, err
	}

	// insert the content in cache
	_ = content.Cache()

	return content, nil
}

func Update(courseID string, content *Content) error {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	err = DBD.UpdateInfo(cID, content.ID, content.Title, content.Description, content.Tags, content.Timestamp)
	if err != nil {
		return err
	}

	// update the content in cache if exists
	content.UpdateCache()

	return nil
}

func Delete(courseID string, content *Content) error {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	err = DBD.Delete(cID, content.ID)
	if err != nil{
		return err
	}

	// delete from cache if exists
	DeleteFromCache(content.ID.Hex())

	return nil
}
