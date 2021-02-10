package graph

import (
	"context"
	"yes-sharifTube/internal/middleware/auth"
	"yes-sharifTube/internal/middleware/ggcontext"
)

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

