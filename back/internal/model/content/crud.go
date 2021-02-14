package content

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
)

func Get(courseID *string, contentID string) (*Content, error) {
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

func Insert(courseID string, content *Content) (*Content, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	content, err = DBD.Insert(cID, content)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Update(courseID string, content *Content) error {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	return DBD.UpdateInfo(cID, content.ID, content.Title, content.Description, content.Tags, content.Timestamp)
}

func Delete(courseID string, content *Content) error {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	return DBD.Delete(cID, content.ID)
}
