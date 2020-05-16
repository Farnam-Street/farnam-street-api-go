package routers

import (
	apiController "farnam-street-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Giving access to storage folder
	r.Static("/storage", "storage")

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	// API route for version 1
	api := r.Group("/api")

	// v1.Use(middlewares.UserMiddlewares())
	// {
	// 	v1.POST("user-list", apiControllerV1.UserList)
	// }
	api.GET("assets", apiController.AssetList)
	api.GET("asset", apiController.Asset)
	api.GET("assetgraphql", apiController.AssetGraphql)

	//If you want to pass your route through specific middlewares

	return r

}
