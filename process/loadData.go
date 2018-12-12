package main

import (
	"dynamodb_2/common"
	"dynamodb_2/model"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	loadData()
}

func loadData() {
	var profiles []model.Profile
	if err := json.Unmarshal(getJsonProfiles(), &profiles); err != nil {
		panic("Could not parse json profiles data")
	}

	svc := common.GetDynamoDbSession()

	for _, prof := range profiles {
		// marshal the profile struct into an aws attribute value map
		profileAVMap, err := dynamodbattribute.MarshalMap(prof)
		if err != nil {
			panic("Cannot marshal profile into AttributeValue map")
		}

		// create the api params
		params := &dynamodb.PutItemInput{
			TableName: aws.String("Profiles"),
			Item:      profileAVMap,
		}

		// put the item
		resp, err := svc.PutItem(params)
		if err != nil {
			fmt.Printf("Unable to add profile: %v\n", err.Error())
		} else {
			// print the response data
			fmt.Printf("Put item successful: '%s' (resp = '%+v')\n", prof.RawJsonData, resp)
		}
	}
}

func getJsonProfiles() []byte {
	return []byte(`[{
		"guid": "71e1515b-cec0-40f2-be1d-70c218fb6721",
		"vertical": "supermarket",
		"observations": "This user is..",
		"active": true,
		"raw_json_data": "{\"date\": \"2013-09-02T00:00:00Z\",\"Name\": \"Sergio\", \"starts\": 7, \"interestCategories\": [\"Carniceria\",\"Herramientas\",\"Deportes\"],\"image_url\": \"https://d2x5ku95bkycr3.cloudfront.net/App_Themes/Common/images/profile/0_200.png\"}"
	},

	{
		"guid": "43e1e209-32f6-428e-892b-2ae818ab4440",
		"vertical": "hardware-store",
		"active": true,
		"raw_json_data": "{\"date\": \"2013-09-02T00:00:00Z\",\"Name\": \"Fernando\", \"starts\": 7, \"interestCategories\": [\"Jardin\"],\"image_url\": \"https://static1.squarespace.com/static/52d4725ee4b0d4a5bfc88830/5756ec9b27d4bd18286a7336/5756ecf19f7266856a4389da/1473876715043/garden-phs.jpg\"}"
	}
]`)
}
