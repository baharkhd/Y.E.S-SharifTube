package auth

import (
"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const AuthHeaderKey = "whoami"
const usernameKey = "username"



func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(AuthHeaderKey)
		if token==""{
			c.Next()
			return
		}
		//Set username in the context
		parts := strings.Split(token, "Bearer ")
		jwtToken := ""
		if len(parts) == 2 {
			jwtToken = parts[1]
		}

		username, err := jwt.ParseToken(jwtToken)
		if err!=nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		c.Set(usernameKey,username)
		c.Next()
	}
}


// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx *gin.Context) string {
	return ctx.GetString(usernameKey)
}