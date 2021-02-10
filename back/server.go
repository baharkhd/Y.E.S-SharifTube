package main

import (
	"log"
	"net/http"
	"os"
	"yes-sharifTube/graph"
	"yes-sharifTube/graph/generated"
	"yes-sharifTube/internal/model/user"
	"yes-sharifTube/pkg/database/mongodb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	user.UserDBD=mongodb.NewUserMongoDriver("yes-blog", "users");

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
