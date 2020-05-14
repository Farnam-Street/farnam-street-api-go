package controllers

import (
	apiHelpers "farnam-street-api-go/apiHelpers"
	assetService "farnam-street-api-go/services"
	// "encoding/json"
	"github.com/gin-gonic/gin"
)

//UserList function will give you the list of users
func AssetList(c *gin.Context) {

	// //decode the request body into struct and failed if any error occur
	// err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	// if err != nil {
	// 	// u.Respond(c.Writer, u.Message(1, "Invalid request"))
	// 	return
	// }

	// //call service
	resp := assetService.AssetList()

	// //return response using api helper
	apiHelpers.Respond(c.Writer, resp)

}
