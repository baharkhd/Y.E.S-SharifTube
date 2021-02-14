package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"yes-sharifTube/graph/generated"
	"yes-sharifTube/graph/model"
	attachmentController "yes-sharifTube/internal/controller/attachment"
	commentController "yes-sharifTube/internal/controller/comment"
	contentController "yes-sharifTube/internal/controller/content"
	courseController "yes-sharifTube/internal/controller/course"
	pendingController "yes-sharifTube/internal/controller/pending"
	"yes-sharifTube/internal/model/user"
	"yes-sharifTube/pkg/jwt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, target model.TargetUser) (model.CreateUserPayload, error) {
	println("user: " + extractUsernameFromContext(ctx))
	newUser, err := user.New(deref(target.Name), deref(target.Email), target.Username, target.Password)
	if err != nil {
		switch err.(type) {
		case model.DuplicateUsernameException:
			return err.(model.DuplicateUsernameException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatUser(newUser), err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, toBe model.EditedUser) (model.UpdateUserPayload, error) {
	username := extractUsernameFromContext(ctx)
	if username == "" {
		return model.UserNotAllowedException{Message: "please login first!"}, nil
	}
	update, err := user.Update(username, toBe)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatUser(update), err
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (model.DeleteUserPayload, error) {
	username := extractUsernameFromContext(ctx)
	err := user.Delete(username)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return model.OperationSuccessfull{}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (model.LoginPayload, error) {
	token, err := user.Login(input.Username, input.Password)
	if err != nil {
		switch err.(type) {
		case model.UserPassMissMatchException:
			return err.(model.UserPassMissMatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return model.Token{Token: token}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (model.LoginPayload, error) {
	username := extractUsernameFromContext(ctx)
	if username == "" {
		return model.InternalServerException{Message: "user not found!"}, nil
	}
	// generate new token
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return model.InternalServerException{}, nil
	}
	return model.Token{Token: token}, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, username *string, target model.TargetCourse) (model.CreateCoursePayload, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := courseController.CreateCourse(*username, target.Title, target.Summary, target.Token)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatCourse(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) UpdateCourseInfo(ctx context.Context, username *string, courseID string, toBe model.EditedCourse) (model.UpdateCourseInfoPayload, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := courseController.UpdateCourse(*username, courseID, toBe.Title, toBe.Summary, toBe.Token)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.EmptyFieldsException:
			return err.(model.EmptyFieldsException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatCourse(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, username *string, courseID string) (model.DeleteCoursePayload, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := courseController.DeleteCourse(*username, courseID)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatCourse(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) AddUserToCourse(ctx context.Context, username *string, courseID string, token string) (model.AddUserToCoursePayload, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := courseController.AddUserToCourse(*username, courseID, token)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.DuplicateUsernameException:
			return err.(model.DuplicateUsernameException), nil
		case model.IncorrectTokenException:
			return err.(model.IncorrectTokenException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatCourse(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) DeleteUserFromCourse(ctx context.Context, username *string, courseID string, targetUsername string) (model.DeleteUserFromCoursePayload, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := courseController.DeleteUserFromCourse(*username, courseID, targetUsername)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatCourse(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) PromoteUserToTa(ctx context.Context, username *string, courseID string, targetUsername string) (model.PromoteToTAPayload, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := courseController.PromoteUserToTA(*username, courseID, targetUsername)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.UserIsNotSTDException:
			return err.(model.UserIsNotSTDException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatCourse(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) DemoteUserToStd(ctx context.Context, username *string, courseID string, targetUsername string) (model.DemoteToSTDPayload, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := courseController.DemoteUserToSTD(*username, courseID, targetUsername)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.UserIsNotTAException:
			return err.(model.UserIsNotTAException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatCourse(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) UploadContent(ctx context.Context, username *string, courseID string, target model.TargetContent) (model.UploadContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := contentController.CreateContent(*username, courseID, target.Title, target.Description, target.Vurl, target.Tags)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatContent(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) EditContent(ctx context.Context, username *string, courseID string, contentID string, target model.EditContent) (model.EditContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := contentController.UpdateContent(*username, courseID, contentID, target.Title, target.Description, target.Tags)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.ContentNotFoundException:
			return err.(model.ContentNotFoundException), nil
		case model.EmptyFieldsException:
			return err.(model.EmptyFieldsException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatContent(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) DeleteContent(ctx context.Context, username *string, courseID string, contentID string) (model.DeleteContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := contentController.DeleteContent(*username, courseID, contentID)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.ContentNotFoundException:
			return err.(model.ContentNotFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatContent(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) UploadAttachment(ctx context.Context, username *string, courseID string, target model.TargetAttachment) (model.UploadAttachmentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := attachmentController.CreateAttachment(*username, courseID, target.Name, target.Description, target.Aurl)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatAttachment(res), nil
}

func (r *mutationResolver) EditAttachment(ctx context.Context, username *string, courseID string, attachmentID string, target model.EditAttachment) (model.EditAttachmentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := attachmentController.UpdateAttachment(*username, courseID, attachmentID, target.Name, target.Description)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.AttachmentNotFoundException:
			return err.(model.AttachmentNotFoundException), nil
		case model.EmptyFieldsException:
			return err.(model.EmptyFieldsException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatAttachment(res), err
}

func (r *mutationResolver) DeleteAttachment(ctx context.Context, username *string, courseID string, attachmentID string) (model.DeleteAttachmentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := attachmentController.DeleteAttachment(*username, courseID, attachmentID)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.AttachmentNotFoundException:
			return err.(model.AttachmentNotFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	return reformatAttachment(res), err
}

func (r *mutationResolver) OfferContent(ctx context.Context, username *string, courseID string, target model.TargetPending) (model.OfferContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := pendingController.CreatePending(*username, courseID, target.Title, target.Description, target.Furl)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatPending(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) EditOfferedContent(ctx context.Context, username *string, courseID string, pendingID string, target model.EditedPending) (model.EditOfferedContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := pendingController.UpdatePending(*username, courseID, pendingID, target.Title, target.Description)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.PendingNotFoundException:
			return err.(model.PendingNotFoundException), nil
		case model.EmptyFieldsException:
			return err.(model.EmptyFieldsException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		case model.OfferedContentNotPendingException:
			return err.(model.OfferedContentNotPendingException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatPending(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) DeleteOfferedContent(ctx context.Context, username *string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := pendingController.DeletePending(*username, courseID, pendingID)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.PendingNotFoundException:
			return err.(model.PendingNotFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatPending(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) AcceptOfferedContent(ctx context.Context, username *string, courseID string, pendingID string, changed model.EditedPending) (model.EditOfferedContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := pendingController.AcceptPending(*username, courseID, pendingID, changed.Title, changed.Description)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.PendingNotFoundException:
			return err.(model.PendingNotFoundException), nil
		case model.OfferedContentNotPendingException:
			return err.(model.OfferedContentNotPendingException), nil
		case model.EmptyFieldsException:
			return err.(model.EmptyFieldsException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatPending(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) RejectOfferedContent(ctx context.Context, username *string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	res, err := pendingController.RejectPending(*username, courseID, pendingID)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.CourseNotFoundException:
			return err.(model.CourseNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.PendingNotFoundException:
			return err.(model.PendingNotFoundException), nil
		case model.OfferedContentNotPendingException:
			return err.(model.OfferedContentNotPendingException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	f, err := reformatPending(res)
	if err != nil {
		return err.(model.InternalServerException), nil
	}
	return f, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, username *string, contentID string, repliedAtID *string, target model.TargetComment) (model.CreateCommentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	con, rep, err := commentController.CreateComment(*username, contentID, target.Body, repliedAtID)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.ContentNotFoundException:
			return err.(model.ContentNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.CommentNotFoundException:
			return err.(model.CommentNotFoundException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	if con != nil {
		f, err := reformatComment(con)
		if err != nil {
			return err.(model.InternalServerException), nil
		}
		return f, nil
	} else {
		f, err := reformatReply(rep)
		if err != nil {
			return err.(model.InternalServerException), nil
		}
		return f, nil
	}
}

func (r *mutationResolver) UpdateComment(ctx context.Context, username *string, contentID string, commentID string, target model.EditedComment) (model.EditCommentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	con, rep, err := commentController.UpdateComment(*username, contentID, commentID, target.Body)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.ContentNotFoundException:
			return err.(model.ContentNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.CommentNotFoundException:
			return err.(model.CommentNotFoundException), nil
		case model.EmptyFieldsException:
			return err.(model.EmptyFieldsException), nil
		case model.RegexMismatchException:
			return err.(model.RegexMismatchException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	if con != nil {
		f, err := reformatComment(con)
		if err != nil {
			return err.(model.InternalServerException), nil
		}
		return f, nil
	} else {
		f, err := reformatReply(rep)
		if err != nil {
			return err.(model.InternalServerException), nil
		}
		return f, nil
	}
}

func (r *mutationResolver) DeleteComment(ctx context.Context, username *string, contentID string, commentID string) (model.DeleteCommentPayLoad, error) {
	if username == nil {
		var err error
		username, err = fetchUsername(ctx)
		if err != nil {
			return err.(model.UserNotFoundException), nil
		}
	}
	con, rep, err := commentController.DeleteComment(*username, contentID, commentID)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundException:
			return err.(model.UserNotFoundException), nil
		case model.ContentNotFoundException:
			return err.(model.ContentNotFoundException), nil
		case model.UserNotAllowedException:
			return err.(model.UserNotAllowedException), nil
		case model.CommentNotFoundException:
			return err.(model.CommentNotFoundException), nil
		default:
			return err.(model.InternalServerException), nil
		}
	}
	if con != nil {
		f, err := reformatComment(con)
		if err != nil {
			return err.(model.InternalServerException), nil
		}
		return f, nil
	} else {
		f, err := reformatReply(rep)
		if err != nil {
			return err.(model.InternalServerException), nil
		}
		return f, nil
	}
}

func (r *queryResolver) User(ctx context.Context, username *string) (*model.User, error) {
	target, err := user.Get(getUsername(ctx, username))
	return reformatUser(target), err
}

func (r *queryResolver) Users(ctx context.Context, start int, amount int) ([]*model.User, error) {
	all, err := user.GetAll(int64(start), int64(amount))
	return reformatUsers(all), err
}

func (r *queryResolver) Courses(ctx context.Context, ids []string) ([]*model.Course, error) {
	res, err := courseController.GetCourses(ids)
	if err != nil {
		return nil, err
	}
	return reformatCourses(res)
}

func (r *queryResolver) CoursesByKeyWords(ctx context.Context, keyWords []string, start int, amount int) ([]*model.Course, error) {
	res, err := courseController.GetCoursesByKeyWords(keyWords, start, amount)
	if err != nil {
		return nil, err
	}
	return reformatCourses(res)
}

func (r *queryResolver) Content(ctx context.Context, id string) (*model.Content, error) {
	res, err := contentController.GetContent(id)
	if err != nil {
		return nil, err
	}
	return reformatContent(res)
}

func (r *queryResolver) Contents(ctx context.Context, tags []string, courseID *string, start int, amount int) ([]*model.Content, error) {
	res, err := contentController.GetContents(tags, courseID, start, amount)
	if err != nil {
		return nil, err
	}
	return reformatContents(res)
}

func (r *queryResolver) Pendings(ctx context.Context, filter model.PendingFilter, start int, amount int) ([]*model.Pending, error) {
	res, err := pendingController.GetPendings(filter.CourseID, filter.UploaderUsername, filter.Status, start, amount)
	if err != nil {
		return nil, err
	}
	return reformatPendings(res)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
