package apihelpers

import (
	"github.com/graphql-go/graphql"
	"log"
)

func ExecuteGraphQLQuery(query string, schema graphql.Schema)interface{}{
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,

	})
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
	}
	return result.Data
}