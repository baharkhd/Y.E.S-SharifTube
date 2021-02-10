package graph

import (
	"context"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/middleware/auth"
	"yes-sharifTube/internal/middleware/ggcontext"
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
	var graphUser = &model.User{
		ID:      targetUser.ID.Hex(),
		Name:    &targetUser.Name,
		Email:   &targetUser.Email,
		Username: targetUser.Username,
		CourseIDs: targetUser.Courses,
	}
	return graphUser
}

func extractUsernameFromContext(ctx context.Context) string {
	ginContext, _ := ggcontext.GinContextFromContext(ctx)
	return auth.ForContext(ginContext)
}

func getUserName(ctx context.Context, name *string) string {
	var username string
	if name == nil || *name == "" {
		username = extractUsernameFromContext(ctx)
	} else {
		username = *name
	}
	return username
}

