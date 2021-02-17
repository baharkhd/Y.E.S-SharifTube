package pending

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
)

func Get(courseID *string, pendingID string) (*Pending, error) {
	var cpID *primitive.ObjectID = nil
	if courseID != nil {
		cID, err := primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		cpID = &cID
	}
	pID, err := primitive.ObjectIDFromHex(pendingID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	pending, err := DBD.Get(cpID, pID)
	if err != nil {
		return nil, err
	}
	return pending, nil
}

func GetByFilter(courseID, uploaderUsername *string, status *model.Status, startIdx, amount int) ([]*Pending, error) {
	var cpID *primitive.ObjectID = nil
	var cID primitive.ObjectID
	var spt *Status = nil
	var st Status
	var err error
	if courseID != nil {
		cID, err = primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		cpID = &cID
	}
	if status != nil {
		st = NewStatus(*status)
		spt = &st
	}
	pends, err := DBD.GetByFilter(cpID, spt, uploaderUsername, startIdx, amount)
	if err != nil {
		return nil, err
	}
	return pends, err
}

func Insert(courseID string, pending *Pending) (*Pending, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	pending, err = DBD.Insert(cID, pending)
	if err != nil {
		return nil, err
	}
	return pending, nil
}

func Update(courseID string, pending *Pending) error {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	return DBD.UpdateInfo(cID, pending.ID, pending.Title, pending.Description, pending.Timestamp)
}

func Delete(courseID string, pending *Pending) error {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	return DBD.Delete(cID, pending.ID)
}

func Accept(courseID string, pending *Pending) (*Pending, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	pending.Status = ACCEPTED
	err = DBD.UpdateStatus(cID, pending.ID, pending.Title, pending.Description, pending.Message, ACCEPTED, pending.Timestamp)
	if err != nil {
		return nil, err
	}
	return pending, nil
}

func Reject(courseID string, pending *Pending) (*Pending, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	pending.Status = REJECTED
	err = DBD.UpdateStatus(cID, pending.ID, pending.Title, pending.Description, pending.Message, REJECTED, pending.Timestamp)
	if err != nil {
		return nil, err
	}
	return pending, nil
}
