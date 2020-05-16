package graphqlSchema

import (
	"farnam-street-api-go/models"
	"github.com/graphql-go/graphql"
	"log"
)

var assets []models.Asset = []models.Asset{
	models.Asset{
		ID: "resourceId",
		Name: "AssetName",
		Description: "AssetDescription",
		Configval1: "AssetConfigval1",
		Configval2: "AssetConfigval2",
		Configval3: "AssetConfigval3",

	},
	models.Asset{
		ID: "resourceId2",
		Name: "AssetName2",
		Description: "AssetDescription2",
		Configval1: "AssetConfigval1_2",
		Configval2: "AssetConfigval2_2",
		Configval3: "AssetConfigval3_2",

	},
}

var assetObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Asset",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return assets, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"configval1": &graphql.Field{
			Type: graphql.String,
		},
		"configval2": &graphql.Field{
			Type: graphql.String,
		},
		"configval3": &graphql.Field{
			Type: graphql.String,
		},
	},
})
var assetQueryObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "assets",
	Fields: graphql.Fields{
		"assets": &graphql.Field{
			Type: graphql.NewList(assetObject),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return assets, nil
			},
		},
	},
})

func GetAssetSchema() graphql.Schema{

	assetSchema, err := graphql.NewSchema(graphql.SchemaConfig{Query: assetQueryObject})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return assetSchema

}