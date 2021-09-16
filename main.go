package main

import (
	"log"
	"net/http"

	"github.com/CodingForFunAndProfit/gographql/gopher"
	"github.com/CodingForFunAndProfit/gographql/job"
	schemas "github.com/CodingForFunAndProfit/gographql/schema"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	gopherService := gopher.NewService(
		gopher.NewMemoryRepository(),
		job.NewMemoryRepository(),
	)

	schema, err := schemas.GenerateSchema(&gopherService)
	if err != nil {
		panic(err)
	}

	StartServer(schema)
}

func StartServer(schema *graphql.Schema) {

	h := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
