package services

import (
	apiHelpers "farnam-street-api-go/handlers"
	gqSchema "farnam-street-api-go/graphqlSchema"

	res "farnam-street-api-go/resources"
)

//UserService struct
// type UserService struct {
// 	User models.User
// }

//UserList function returns the list of users
func AssetList() map[string]interface{} {

	assetData := res.AssetResponse{
		ID:   1,
		Name: "test",
	}
	response := apiHelpers.Message(0, "This is from version 1 api")
	response["data"] = assetData
	return response
}

func Asset() map[string] interface{} {
	assetData := res.AssetResponse{
		ID:   1,
		Name: "test",
	}
	response := apiHelpers.Message(0, "This is from version 1 api")
	response["data"] = assetData
	return response

}

func AssetGraphQL(graphqlQuery string) map[string]interface{} {

	result := apiHelpers.ExecuteGraphQLQuery(graphqlQuery, gqSchema.GetAssetSchema())
	response := apiHelpers.Message(0, "This is from graphql")
	response["data"] = result
	return response

}
