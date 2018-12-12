package main

import (
	"dynamodb_2/common"
	"dynamodb_2/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	createItem()
}

func createItem() {
	profile := model.Profile{
		GUI:         "b5e6d907-a953-4a79-be38-0f97b4ca29a4",
		Vertical:    "Pharmacy",
		Active:      true,
		RawJsonData: "{\"date\": \"2013-09-02T00:00:00Z\",\"Name\": \"Alberto\", \"starts\": 9, \"interestCategories\": [\"Receta Verde\"],\"image_url\": \"http://c3.thejournal.ie/media/2018/08/shutterstock_717437125-2-390x285.jpg\"}",
	}

	av, err := dynamodbattribute.MarshalMap(profile)

	if err != nil {
		panic("Got error marshalling map profile")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Profiles"),
	}

	svc := common.GetDynamoDbSession()

	_, err = svc.PutItem(input)

	if err != nil {
		panic("Got error calling PutItem:")
	}

	fmt.Println("Successfully added")
}
