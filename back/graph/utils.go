package graph

import (
	"context"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/middleware/auth"
	"yes-sharifTube/internal/middleware/ggcontext"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/comment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/pending"
	"yes-sharifTube/internal/model/user"
)

/* some useful functions to convert model objects from our models to graphql models
 */
func reformatUsers(all []*user.User) []*model.User {
	var result []*model.User
	for _, targetUser := range all {
		result = append(result, reformatUser(targetUser))
	}
	return result
}

func reformatUser(targetUser *user.User) *model.User {
	if targetUser == nil {
		return nil
	}
	return &model.User{
		ID:        targetUser.ID.Hex(),
		Name:      &targetUser.Name,
		Email:     &targetUser.Email,
		Username:  targetUser.Username,
		CourseIDs: targetUser.Courses,
	}
}

func extractUsernameFromContext(ctx context.Context) string {
	ginContext, _ := ggcontext.GinContextFromContext(ctx)
	return auth.ForContext(ginContext)
}

func getUsername(ctx context.Context, name *string) string {
	var username string
	if name == nil || *name == "" {
		username = extractUsernameFromContext(ctx)
	} else {
		username = *name
	}
	return username
}

func deref(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func fetchUsername(ctx context.Context) (*string, error) {
	username := extractUsernameFromContext(ctx)
	if username == "" {
		return nil, model.UserNotAllowedException{Message: "please login first!"}
	}
	return &username, nil
}

func reformatCourse(c *course.Course) *model.Course {
	res := &model.Course{
		ID:        c.ID.Hex(),
		Title:     c.Title,
		Summary:   &c.Summery,
		CreatedAt: int(c.CreatedAt),
		Prof:      reformatUser(user.GetS(c.ProfUn)),
		Tas:       reformatUsers(user.GetA(c.TaUns)),
		Pends:     nil,
		Students:  reformatUsers(user.GetA(c.StdUns)),
		Contents:  nil,
		Inventory: nil,
	}

	//reshape pendings
	res.Pends = reformatPendings(c.Pends)

	//reshape contents
	res.Contents = reformatContents(c.Contents)

	//reshape inventory
	res.Inventory = reformatAttachments(c.Inventory)

	return res
}

func reformatCourses(courses []*course.Course) []*model.Course {
	var cs []*model.Course
	for _, c := range courses {
		cs = append(cs, reformatCourse(c))
	}
	return cs
}

func reformatPending(p *pending.Pending) *model.Pending {
	return &model.Pending{
		ID:          p.ID.Hex(),
		Title:       p.Title,
		Description: &p.Description,
		Status:      p.Status.Reshape(),
		Timestamp:   int(p.Timestamp),
		UploadedBy:  reformatUser(user.GetS(p.UploadedByUn)),
		Furl:        p.Furl,
		CourseID:    p.CourseID,
	}
}

func reformatPendings(pendings []*pending.Pending) []*model.Pending {
	var ps []*model.Pending
	for _, p := range pendings {
		ps = append(ps, reformatPending(p))
	}
	return ps
}

func reformatContent(c *content.Content) *model.Content {
	res := &model.Content{
		ID:          c.ID.Hex(),
		Title:       c.Title,
		Description: &c.Description,
		Timestamp:   int(c.Timestamp),
		UploadedBy:  reformatUser(user.GetS(c.UploadedByUn)),
		ApprovedBy:  reformatUser(user.GetS(c.ApprovedByUn)),
		Vurl:        c.Vurl,
		Tags:        c.Tags,
		Comments:    nil,
		CourseID:    c.CourseID,
	}

	//reshape comments
	res.Comments = reformatComments(c.Comments)

	return res
}

func reformatContents(contents []*content.Content) []*model.Content {
	var cs []*model.Content
	for _, c := range contents {
		cs = append(cs, reformatContent(c))
	}
	return cs
}

func reformatAttachment(a *attachment.Attachment) *model.Attachment {
	return &model.Attachment{
		ID:          a.ID.Hex(),
		Name:        a.Name,
		Aurl:        a.Aurl,
		Description: &a.Description,
		Timestamp:   int(a.Timestamp),
		CourseID:    a.CourseID,
	}
}

func reformatAttachments(attachments []*attachment.Attachment) []*model.Attachment {
	var cs []*model.Attachment
	for _, c := range attachments {
		cs = append(cs, reformatAttachment(c))
	}
	return cs
}

func reformatComment(c *comment.Comment) *model.Comment {
	res := &model.Comment{
		ID:        c.ID.Hex(),
		Author:    reformatUser(user.GetS(c.AuthorUn)),
		Body:      c.Body,
		Timestamp: int(c.Timestamp),
		Replies:   nil,
		ContentID: c.ContentID,
	}

	//reshape replies
	res.Replies = ReshapeAllReplies(c.Replies)

	return res
}

func reformatComments(courses []*comment.Comment) []*model.Comment {
	var cs []*model.Comment
	for _, c := range courses {
		cs = append(cs, reformatComment(c))
	}
	return cs
}

func reformatReply(r *comment.Reply) *model.Reply {
	return &model.Reply{
		ID:        r.ID.Hex(),
		Author:    reformatUser(user.GetS(r.AuthorUn)),
		Body:      r.Body,
		Timestamp: int(r.Timestamp),
		CommentID: r.CommentID,
	}
}

func ReshapeAllReplies(replies []*comment.Reply) []*model.Reply {
	var cs []*model.Reply
	for _, c := range replies {
		cs = append(cs, reformatReply(c))
	}
	return cs
}
