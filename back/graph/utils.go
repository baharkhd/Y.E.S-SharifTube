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

func reformatCourse(c *course.Course) (*model.Course, error) {
	prof, _ := user.Get(c.ProfUn)
	tas, _ := user.GetS(c.TaUns)
	students, _ := user.GetS(c.StdUns)

	res := &model.Course{
		ID:        c.ID.Hex(),
		Title:     c.Title,
		Summary:   &c.Summery,
		CreatedAt: int(c.CreatedAt),
		Prof:      reformatUser(prof),
		Tas:       reformatUsers(tas),
		Pends:     nil,
		Students:  reformatUsers(students),
		Contents:  nil,
		Inventory: nil,
	}

	//reshape pendings
	pends, err := reformatPendings(c.Pends)
	if err != nil {
		return nil, model.InternalServerException{Message: "error while reshape pending array of course: /n" + err.Error()}
	}
	res.Pends = pends

	//reshape contents
	contents, err := reformatContents(c.Contents)
	if err != nil {
		return nil, model.InternalServerException{Message: "error while reshape contents of course: /n" + err.Error()}
	}
	res.Contents = contents

	//reshape inventory
	res.Inventory = reformatAttachments(c.Inventory)

	return res, nil
}

func reformatCourses(courses []*course.Course) ([]*model.Course, error) {
	var cs []*model.Course
	for _, c := range courses {
		tmp, err := reformatCourse(c)
		if err != nil {
			return nil, model.InternalServerException{Message: "error while reshape course array: " + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}

func reformatPending(p *pending.Pending) (*model.Pending, error) {
	uploader, _ := user.Get(p.UploadedByUn)

	return &model.Pending{
		ID:          p.ID.Hex(),
		Title:       p.Title,
		Description: &p.Description,
		Status:      p.Status.Reshape(),
		Timestamp:   int(p.Timestamp),
		UploadedBy:  reformatUser(uploader),
		Furl:        p.Furl,
		CourseID:    p.CourseID,
	}, nil
}

func reformatPendings(pendings []*pending.Pending) ([]*model.Pending, error) {
	var ps []*model.Pending
	for _, p := range pendings {
		tmp, err := reformatPending(p)
		if err != nil {
			return nil, model.InternalServerException{Message: "error while reshape pending array: /n" + err.Error()}
		}
		ps = append(ps, tmp)
	}
	return ps, nil
}

func reformatContent(c *content.Content) (*model.Content, error) {
	uploader, _ := user.Get(c.UploadedByUn)
	approver, _ := user.Get(c.ApprovedByUn)

	res := &model.Content{
		ID:          c.ID.Hex(),
		Title:       c.Title,
		Description: &c.Description,
		Timestamp:   int(c.Timestamp),
		UploadedBy:  reformatUser(uploader),
		ApprovedBy:  reformatUser(approver),
		Vurl:        c.Vurl,
		Tags:        c.Tags,
		Comments:    nil,
		CourseID:    c.CourseID,
	}

	//reshape comments
	comments, err := reformatComments(c.Comments)
	if err != nil {
		return nil, model.InternalServerException{Message: "error while reshape comments of content: /n" + err.Error()}
	}
	res.Comments = comments

	return res, nil
}

func reformatContents(contents []*content.Content) ([]*model.Content, error) {
	var cs []*model.Content
	for _, c := range contents {
		tmp, err := reformatContent(c)
		if err != nil {
			return nil, model.InternalServerException{Message: "error while reshape content array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
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

func reformatComment(c *comment.Comment) (*model.Comment, error) {
	author, _ := user.Get(c.AuthorUn)

	res := &model.Comment{
		ID:        c.ID.Hex(),
		Author:    reformatUser(author),
		Body:      c.Body,
		Timestamp: int(c.Timestamp),
		Replies:   nil,
		ContentID: c.ContentID,
	}

	//reshape replies
	replies, err := ReshapeAllReplies(c.Replies)
	if err != nil {
		return nil, model.InternalServerException{Message: "error while reshape replies of comment: /n" + err.Error()}
	}
	res.Replies = replies

	return res, nil
}

func reformatComments(courses []*comment.Comment) ([]*model.Comment, error) {
	var cs []*model.Comment
	for _, c := range courses {
		tmp, err := reformatComment(c)
		if err != nil {
			return nil, model.InternalServerException{Message: "error while reshape comment array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}

func reformatReply(r *comment.Reply) (*model.Reply, error) {
	author, _ := user.Get(r.AuthorUn)

	return &model.Reply{
		ID:        r.ID.Hex(),
		Author:    reformatUser(author),
		Body:      r.Body,
		Timestamp: int(r.Timestamp),
		CommentID: r.CommentID,
	}, nil
}

func ReshapeAllReplies(replies []*comment.Reply) ([]*model.Reply, error) {
	var cs []*model.Reply
	for _, c := range replies {
		tmp, err := reformatReply(c)
		if err != nil {
			return nil, model.InternalServerException{Message: "error while reshape reply array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}
