package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"yes-sharifTube/graph/generated"
	"yes-sharifTube/graph/model"
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

func (r *mutationResolver) CreateCourse(ctx context.Context, username string, target model.TargetCourse) (model.CreateCoursePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCourseInfo(ctx context.Context, username string, courseID string, toBe model.EditedCourse) (model.UpdateCourseInfoPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, username string, courseID string) (model.DeleteCoursePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUserToCourse(ctx context.Context, username string, courseID string, token string) (model.AddUserToCoursePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PromoteUserToTa(ctx context.Context, username string, courseID string, targetUserID string) (model.PromoteToTAPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DemoteUserToStd(ctx context.Context, username string, courseID string, targetUserID string) (model.DemoteToSTDPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadContent(ctx context.Context, username string, courseID string, target model.TargetContent) (model.UploadContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditContent(ctx context.Context, username string, courseID string, contentID string, target model.EditContent) (model.EditContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteContent(ctx context.Context, username string, courseID string, contentID string) (model.DeleteContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadAttachment(ctx context.Context, username string, courseID string, target model.TargetContent) (model.UploadAttachmentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditAttachment(ctx context.Context, username string, courseID string, attachmentID string, target model.EditContent) (model.EditAttachmentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAttachment(ctx context.Context, username string, courseID string, attachmentID string) (model.DeleteAttachmentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OfferContent(ctx context.Context, username string, courseID string, target model.TargetPending) (model.OfferContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditOfferedContent(ctx context.Context, username string, courseID string, pendingID string, target model.EditedPending) (model.EditOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOfferedContent(ctx context.Context, username string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AcceptOfferedContent(ctx context.Context, username string, courseID string, pendingID string, changed model.EditedPending) (model.EditOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RejectOfferedContent(ctx context.Context, username string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, username string, contentID string, repliedAtID *string, target model.TargetComment) (model.CreateCommentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, username string, contentID string, commentID string, target model.EditedComment) (model.EditCommentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, username string, contentID string, commentID string) (model.DeleteCommentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, username *string) (*model.User, error) {
	target, err := user.Get(getUserName(ctx, username))
	return reformatUser(target), err
}

func (r *queryResolver) Users(ctx context.Context, start int, amount int) ([]*model.User, error) {
	all, err := user.GetAll(int64(start), int64(amount))
	return reformatUsers(all), err
}

func (r *queryResolver) Courses(ctx context.Context, id []string) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoursesByKeyWord(ctx context.Context, keyWord *string, start int, amount int) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Content(ctx context.Context, id string) (*model.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contents(ctx context.Context, tags []string, courseID *string, start int, amount int) ([]*model.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pendings(ctx context.Context, filter model.PendingFilter, start int, amount int) ([]*model.Pending, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
