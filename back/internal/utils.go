package internal

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"yes-sharifTube/graph/model"
	database "yes-sharifTube/pkg/database"
)

func HashToken(pwd []byte) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost);
		err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

func CheckTokenHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ConvertStringsToObjectIDs(arr []string) ([]primitive.ObjectID, error) {
	var coIDs []primitive.ObjectID
	for _, cID := range arr {
		objID, err := primitive.ObjectIDFromHex(cID)
		if err != nil {
			return nil, err
		}
		coIDs = append(coIDs, objID)
	}
	return coIDs, nil
}

// casting database errors to model.graphQL exceptions
func CastDBExceptionToGQLException(err error) error {
	if err != nil {
		switch err.(type) {
		case *database.InternalDBError:
			return &model.InternalServerException{Message: err.Error()}
		case *database.AllFieldsEmpty:
			return &model.AllFieldsEmptyException{Message: err.Error()}
		case *database.DuplicateUsername:
			return &model.DuplicateUsernameException{Message: err.Error()}
		case *database.UserNotFound:
			return &model.UserNotFoundException{Message: err.Error()}
		case *database.UserNotAllowed:
			return &model.UserNotAllowedException{Message: err.Error()}
		case *database.CourseNotFound:
			return &model.CourseNotFoundException{Message: err.Error()}
		case *database.IncorrectToken:
			return &model.IncorrectTokenException{Message: err.Error()}
		case *database.UserIsNotTA:
			return &model.UserIsNotTAException{Message: err.Error()}
		case *database.UserIsNotSTD:
			return &model.UserIsNotSTDException{Message: err.Error()}
		case *database.ContentNotFound:
			return &model.ContentNotFoundException{Message: err.Error()}
		case *database.AttachmentNotFound:
			return &model.AttachmentNotFoundException{Message: err.Error()}
		case *database.PendingNotFound:
			return &model.PendingNotFoundException{Message: err.Error()}
		case *database.CommentNotFound:
			return &model.CommentNotFoundException{Message: err.Error()}
		case *database.OfferedContentRejected:
			return &model.OfferedContentRejectedException{Message: err.Error()}
		}
	}
	return nil
}
