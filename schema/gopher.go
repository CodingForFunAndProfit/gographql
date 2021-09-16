package schemas

import "github.com/graphql-go/graphql"

var GopherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gopher",

	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "The ID that is used to identify unique gophers",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the gopher",
		},
		"hired": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "True if the Gopher is employeed",
		},
		"profession": &graphql.Field{
			Type:        graphql.String,
			Description: "The gophers last/current profession",
		},
	},
})
