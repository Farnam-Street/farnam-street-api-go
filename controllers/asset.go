package controllers

import (
	apiHelpers "farnam-street-api-go/handlers"
	assetService "farnam-street-api-go/services"
	//"fmt"

	//"github.com/graphql-go/graphql"
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

func Asset(c *gin.Context) {
	resp := assetService.Asset()

	apiHelpers.Respond(c.Writer, resp)
}

func AssetGraphql(c *gin.Context){


	queryParameters := c.Request.URL.Query()
	graphqlQuery := queryParameters["query"][0]
	resp := assetService.AssetGraphQL(graphqlQuery)
	apiHelpers.Respond(c.Writer, resp)
}
