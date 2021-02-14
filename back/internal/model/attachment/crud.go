package attachment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
)

func Get(courseID *string, contentID string) (*Attachment, error) {
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
	attachment, err := DBD.Get(cpID, conID)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func Insert(courseID string, attachment *Attachment) (*Attachment, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	attachment, err = DBD.Insert(cID, attachment)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func Update(courseID string, attachment *Attachment) error{
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	return DBD.UpdateInfo(cID, attachment.ID, attachment.Name, attachment.Description, attachment.Timestamp)
}

func Delete(courseID string, attachment *Attachment) error{
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	return DBD.Delete(cID, attachment.ID)
}
