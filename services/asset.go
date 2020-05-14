package services

import (
	apiHelpers "farnam-street-api-go/apiHelpers"
	// "Golang-Project-Structure/models"
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
