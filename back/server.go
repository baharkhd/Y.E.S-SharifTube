package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"yes-sharifTube/graph"
	"yes-sharifTube/graph/generated"
	attachmentController "yes-sharifTube/internal/controller/attachment"
	commentController "yes-sharifTube/internal/controller/comment"
	contentController "yes-sharifTube/internal/controller/content"
	courseController "yes-sharifTube/internal/controller/course"
	pendingController "yes-sharifTube/internal/controller/pending"
	"yes-sharifTube/internal/middleware/auth"
	"yes-sharifTube/internal/middleware/ggcontext"
	"yes-sharifTube/internal/model/user"

	"yes-sharifTube/pkg/database/mongodb"
)

const defaultPort = "8080"
const queryComplexity = 8

func main() {
	//setting a mongodb driver for DBDriver filed of user model
	user.DBD = mongodb.NewUserMongoDriver("yes-sharifTube", "users")
	courseController.GetCourseController().SetDBDriver(mongodb.NewCourseMongoDriver("yes-sharifTube", "courses"))
	contentController.GetContentController().SetDBDriver(mongodb.NewContentMongoDriver("yes-sharifTube", "courses"))
	pendingController.GetPendingController().SetDBDriver(mongodb.NewPendingMongoDriver("yes-sharifTube", "courses"))
	attachmentController.GetAttachmentController().SetDBDriver(mongodb.NewAttachmentMongoDriver("yes-sharifTube", "courses"))
	commentController.GetCommentController().SetDBDriver(mongodb.NewCommentMongoDriver("yes-sharifTube", "courses"))


	// Setting up Gin
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{auth.AuthHeaderKey,
			"content-type"},
		AllowCredentials: true,
	}))
	r.Use(ggcontext.GinContextToContextMiddleware())
	r.Use(auth.Middleware())

	// routing
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	//let it begin
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	r.Run(":" + defaultPort)

}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//srv.Use(extension.FixedComplexityLimit(queryComplexity))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	pg := playground.Handler("Yes-blog playground", "/query")

	return func(c *gin.Context) {
		pg.ServeHTTP(c.Writer, c.Request)
	}
}
