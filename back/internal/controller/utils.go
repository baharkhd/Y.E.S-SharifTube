package controller

import (
	"yes-sharifTube/graph/model"
	"yes-sharifTube/pkg/database"
)


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
