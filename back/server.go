package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/coocood/freecache"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"yes-sharifTube/graph"
	"yes-sharifTube/graph/generated"
	"yes-sharifTube/internal/middleware/auth"
	"yes-sharifTube/internal/middleware/ggcontext"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/comment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/pending"
	"yes-sharifTube/internal/model/user"
	"yes-sharifTube/pkg/database/mongodb"
	"yes-sharifTube/pkg/objectstorage/baremetal"
)

const defaultPort = "8080"
const queryComplexity = 8

func main() {
	//setting default Object storage for content model
	if driver, err := baremetal.New("localhost:22", "kycilius", "/home/kycilius/Documents/dev-null/shariftube");err!=nil{
		panic(err)
	}else {
		course.OSD=driver
	}


	//setting a mongodb driver for DBDriver filed of user model
	user.DBD = mongodb.NewUserMongoDriver("yes-sharifTube", "users")
	course.DBD = mongodb.NewCourseMongoDriver("yes-sharifTube", "courses")
	content.DBD = mongodb.NewContentMongoDriver("yes-sharifTube", "courses")
	pending.DBD = mongodb.NewPendingMongoDriver("yes-sharifTube", "courses")
	attachment.DBD = mongodb.NewAttachmentMongoDriver("yes-sharifTube", "courses")
	comment.DBD = mongodb.NewCommentMongoDriver("yes-sharifTube", "courses")
	// adding the deleted account in database
	if err := user.SetDeletedAccount(); err != nil {
		panic(err)
	}

	// set 1Gb cache
	cacheSize := 1000 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	course.Cache = cache
	content.Cache = cache
	user.Cache = cache

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
