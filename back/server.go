package main

import (
	"log"
	"net/http"
	"os"
	"yes-sharifTube/graph"
	"yes-sharifTube/graph/generated"
	attachmentController "yes-sharifTube/internal/controller/attachment"
	commentController "yes-sharifTube/internal/controller/comment"
	contentController "yes-sharifTube/internal/controller/content"
	courseController "yes-sharifTube/internal/controller/course"
	pendingController "yes-sharifTube/internal/controller/pending"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"yes-sharifTube/pkg/database/mongodb"
)

const defaultPort = "8080"

func main() {

	courseController.GetCourseController().SetDBDriver(mongodb.NewCourseMongoDriver("yes-sharifTube", "courses"))
	contentController.GetContentController().SetDBDriver(mongodb.NewContentMongoDriver("yes-sharifTube", "courses"))
	pendingController.GetPendingController().SetDBDriver(mongodb.NewPendingMongoDriver("yes-sharifTube", "courses"))
	attachmentController.GetAttachmentController().SetDBDriver(mongodb.NewAttachmentMongoDriver("yes-sharifTube", "courses"))
	commentController.GetCommentController().SetDBDriver(mongodb.NewCommentMongoDriver("yes-sharifTube", "courses"))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
