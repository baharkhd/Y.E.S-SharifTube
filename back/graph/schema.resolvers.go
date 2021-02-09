package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"yes-sharifTube/graph/generated"
	"yes-sharifTube/graph/model"
	attachmentController "yes-sharifTube/internal/controller/attachment"
	commentController "yes-sharifTube/internal/controller/comment"
	contentController "yes-sharifTube/internal/controller/content"
	courseController "yes-sharifTube/internal/controller/course"
	pendingController "yes-sharifTube/internal/controller/pending"
)

func (r *mutationResolver) CreateUser(ctx context.Context, target model.TargetUser) (model.CreateUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userName string, toBe model.EditedUser) (model.UpdateUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userName string) (model.DeleteUserPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (model.LoginPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (model.LoginPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCourse(ctx context.Context, userName string, target model.TargetCourse) (model.CreateCoursePayload, error) {
	return courseController.GetCourseController().CreateCourse(userName, target.Title, *target.Summary, *target.Token)
}

func (r *mutationResolver) UpdateCourseInfo(ctx context.Context, userName string, courseID string, toBe model.EditedCourse) (model.UpdateCourseInfoPayload, error) {
	return courseController.GetCourseController().UpdateCourse(userName, courseID, *toBe.Title, *toBe.Summary, *toBe.Token)
}

func (r *mutationResolver) DeleteCourse(ctx context.Context, userName string, courseID string) (model.DeleteCoursePayload, error) {
	return courseController.GetCourseController().DeleteCourse(userName, courseID)
}

func (r *mutationResolver) AddUserToCourse(ctx context.Context, userName string, courseID string, token string) (model.AddUserToCoursePayload, error) {
	return courseController.GetCourseController().AddUserToCourse(userName, courseID, token)
}

func (r *mutationResolver) DeleteUserFromCourse(ctx context.Context, userName string, courseID string, targetUsername string) (model.DeleteUserFromCoursePayload, error) {
	return courseController.GetCourseController().DeleteUserFromCourse(userName, courseID, targetUsername)
}

func (r *mutationResolver) PromoteUserToTa(ctx context.Context, userName string, courseID string, targetUsername string) (model.PromoteToTAPayload, error) {
	return courseController.GetCourseController().PromoteUserToTA(userName, courseID, targetUsername)
}

func (r *mutationResolver) DemoteUserToStd(ctx context.Context, userName string, courseID string, targetUsername string) (model.DemoteToSTDPayload, error) {
	return courseController.GetCourseController().DemoteUserToSTD(userName, courseID, targetUsername)
}

func (r *mutationResolver) UploadContent(ctx context.Context, userName string, courseID string, target model.TargetContent) (model.UploadContentPayLoad, error) {
	return contentController.GetContentController().CreateContent(userName, courseID, target.Title, *target.Description, target.Vurl, target.Tags)
}

func (r *mutationResolver) EditContent(ctx context.Context, userName string, courseID string, contentID string, target model.EditContent) (model.EditContentPayLoad, error) {
	return contentController.GetContentController().UpdateContent(userName, courseID, contentID, *target.Title, *target.Description, target.Tags)
}

func (r *mutationResolver) DeleteContent(ctx context.Context, userName string, courseID string, contentID string) (model.DeleteContentPayLoad, error) {
	return contentController.GetContentController().DeleteContent(userName, courseID, contentID)
}

func (r *mutationResolver) UploadAttachment(ctx context.Context, userName string, courseID string, target model.TargetAttachment) (model.UploadAttachmentPayLoad, error) {
	return attachmentController.GetAttachmentController().CreateAttachment(userName, courseID, target.Name, *target.Description, target.Aurl)
}

func (r *mutationResolver) EditAttachment(ctx context.Context, userName string, courseID string, attachmentID string, target model.EditAttachment) (model.EditAttachmentPayLoad, error) {
	return attachmentController.GetAttachmentController().UpdateAttachment(userName, courseID, attachmentID, *target.Name, *target.Description)
}

func (r *mutationResolver) DeleteAttachment(ctx context.Context, userName string, courseID string, attachmentID string) (model.DeleteAttachmentPayLoad, error) {
	return attachmentController.GetAttachmentController().DeleteAttachment(userName, courseID, attachmentID)
}

func (r *mutationResolver) OfferContent(ctx context.Context, userName string, courseID string, target model.TargetPending) (model.OfferContentPayLoad, error) {
	return pendingController.GetPendingController().CreatePending(userName, courseID, target.Title, *target.Description, target.Furl)
}

func (r *mutationResolver) EditOfferedContent(ctx context.Context, userName string, courseID string, pendingID string, target model.EditedPending) (model.EditOfferedContentPayLoad, error) {
	return pendingController.GetPendingController().UpdatePending(userName, courseID, pendingID, *target.Title, *target.Description)
}

func (r *mutationResolver) DeleteOfferedContent(ctx context.Context, userName string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	return pendingController.GetPendingController().DeletePending(userName, courseID, pendingID)
}

func (r *mutationResolver) AcceptOfferedContent(ctx context.Context, userName string, courseID string, pendingID string, changed model.EditedPending) (model.EditOfferedContentPayLoad, error) {
	return pendingController.GetPendingController().AcceptPending(userName, courseID, pendingID, *changed.Title, *changed.Description)
}

func (r *mutationResolver) RejectOfferedContent(ctx context.Context, userName string, courseID string, pendingID string) (model.DeleteOfferedContentPayLoad, error) {
	return pendingController.GetPendingController().RejectPending(userName, courseID, pendingID)
}

func (r *mutationResolver) CreateComment(ctx context.Context, userName string, contentID string, repliedAtID *string, target model.TargetComment) (model.CreateCommentPayLoad, error) {
	con, rep, e := commentController.GetCommentController().CreateComment(userName, contentID, target.Body, repliedAtID)
	if con != nil {
		return con, e
	} else {
		return rep, e
	}
}

func (r *mutationResolver) UpdateComment(ctx context.Context, userName string, contentID string, commentID string, target model.EditedComment) (model.EditCommentPayLoad, error) {
	con, rep, e := commentController.GetCommentController().UpdateComment(userName, contentID, commentID, *target.Body)
	if con != nil {
		return con, e
	} else {
		return rep, e
	}
}

func (r *mutationResolver) DeleteComment(ctx context.Context, userName string, contentID string, commentID string) (model.DeleteCommentPayLoad, error) {
	con, rep, e := commentController.GetCommentController().DeleteComment(userName, contentID, commentID)
	if con != nil {
		return con, e
	} else {
		return rep, e
	}
}

func (r *queryResolver) User(ctx context.Context, username *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, start int, amount int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Courses(ctx context.Context, ids []string) ([]*model.Course, error) {
	return courseController.GetCourseController().GetCourses(ids)
}

func (r *queryResolver) CoursesByKeyWords(ctx context.Context, keyWords []string, start int, amount int) ([]*model.Course, error) {
	return courseController.GetCourseController().GetCoursesByKeyWords(keyWords, start, amount)
}

func (r *queryResolver) Content(ctx context.Context, id string) (*model.Content, error) {
	return contentController.GetContentController().GetContent(id)
}

func (r *queryResolver) Contents(ctx context.Context, tags []string, courseID *string, start int, amount int) ([]*model.Content, error) {
	return contentController.GetContentController().GetContents(tags, courseID, start, amount)
}

func (r *queryResolver) Pendings(ctx context.Context, filter model.PendingFilter, start int, amount int) ([]*model.Pending, error) {
	return pendingController.GetPendingController().GetPendings(filter.CourseID, filter.UploaderUsername, filter.Status, start, amount)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
