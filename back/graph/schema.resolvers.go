package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"yes-sharifTube/graph/generated"
	"yes-sharifTube/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, target model.TargetUser) (model.CreateUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, toBe model.EditedUser) (model.UpdateUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (model.DeleteUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (model.LoginPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (model.LoginPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCourse(ctx context.Context, userID string, target model.TargetCourse) (model.CreateCoursePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCourseInfo(ctx context.Context, userID string, courseID string, toBe model.EditedCourse) (model.UpdateCourseInfoPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, userID string, courseID string) (model.DeleteCoursePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUserToCourse(ctx context.Context, userID string, courseID string, token string) (model.AddUserToCoursePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PromoteUserToTa(ctx context.Context, userID string, courseID string, targetUserID string) (model.PromoteToTAPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DemoteUserToStd(ctx context.Context, userID string, courseID string, targetUserID string) (model.DemoteToSTDPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadContent(ctx context.Context, userID string, courseID string, target model.TargetContent) (model.UploadContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditContent(ctx context.Context, userID string, courseID string, contentID string, target model.EditContent) (model.EditContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteContent(ctx context.Context, userID string, courseID string, contentID string) (model.DeleteContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadAttachment(ctx context.Context, userID string, courseID string, target model.TargetContent) (model.UploadAttachmentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditAttachment(ctx context.Context, userID string, courseID string, attachmentID string, target model.EditContent) (model.EditAttachmentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAttachment(ctx context.Context, userID string, courseID string, attachmentID string) (model.DeleteAttachmentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OfferContent(ctx context.Context, userID string, courseID string, target model.TargetPending) (model.OfferContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditOfferedContent(ctx context.Context, userID string, courseID string, pendingID string, target model.EditedPending) (model.EditOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOfferedContent(ctx context.Context, userID string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AcceptOfferedContent(ctx context.Context, userID string, courseID string, pendingID string, changed model.ChangedPending) (model.EditOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RejectOfferedContent(ctx context.Context, userID string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, userID string, contentID string, target model.TargetComment) (model.CreateCommentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateComment(ctx context.Context, userID string, contentID string, commentID string, target model.EditedComment) (model.EditCommentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, userID string, contentID string, commentID string) (model.DeleteCommentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReplyComment(ctx context.Context, userID string, contentID string, commentID string, target model.TargetComment) (model.ReplyCommentPayLoad, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, start int, amount int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserByFilter(ctx context.Context, filter model.UserFilter, start int, amount int) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Courses(ctx context.Context, start int, amount int) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CourseByFilter(ctx context.Context, filter model.CourseFilter, start int, amount int) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Content(ctx context.Context, id string) (*model.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contents(ctx context.Context, start int, amount int) ([]*model.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ContentByFilter(ctx context.Context, filter model.CourseFilter, start int, amount int) ([]*model.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CommentByFilter(ctx context.Context, filter model.CommentFilter, start int, amount int) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Attachment(ctx context.Context, id string) (*model.Attachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Attachments(ctx context.Context, courseID string, start int, amount int) ([]*model.Attachment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pending(ctx context.Context, id string) (*model.Pending, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PendingByFilter(ctx context.Context, filter model.PendingFilter, start int, amount int) ([]*model.Pending, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
