package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	itemUpdatePartialContent()
}

func itemUpdatePartialContent() {
	svc := common.GetDynamoDbSession()

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"vertical": {
				S: aws.String("Pharmacy"),
			},
			"guid": {
				S: aws.String("b5e6d907-a953-4a79-be38-0f97b4ca29a4"),
			},
		},
		TableName:        aws.String("Profiles"),
		UpdateExpression: aws.String("set raw_json_data = :r"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String("{\"date\": \"2013-09-21T00:00:00Z\",\"Name\": \"Alberto-Updated\", \"starts\": 9, \"interestCategories\": [\"Jardin\"],\"image_url\": \"https://static1.squarespace.com/static/52d4725ee4b0d4a5bfc88830/5756ec9b27d4bd18286a7336/5756ecf19f7266856a4389da/1473876715043/garden-phs.jpg\"}"),
			},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated")
}
