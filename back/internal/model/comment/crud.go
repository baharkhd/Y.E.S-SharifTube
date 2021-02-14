package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
)

func InsertComment(courseID, contentID string, comment *Comment) (*Comment, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	conID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	comment, err = DBD.InsertComment(cID, conID, comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func InsertReply(courseID, contentID string, repliedID string, reply *Reply) (*Reply, error) {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	conID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	repID, err := primitive.ObjectIDFromHex(repliedID)
	if err != nil {
		return nil, model.InternalServerException{Message: err.Error()}
	}
	reply, err = DBD.InsertReply(cID, conID, repID, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

//todo better implementation
func Update(courseID, contentID string, comments []*Comment) error {
	cID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	courID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	return DBD.UpdateCommentsForContent(cID, courID, comments)
}