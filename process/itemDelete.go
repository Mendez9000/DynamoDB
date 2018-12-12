package main

import (
	"dynamodb_2/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	itemDelete()
}

func itemDelete() {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"guid": {
				S: aws.String("b5e6d907-a953-4a79-be38-0f97b4ca29a4"),
			},
			"vertical": {
				S: aws.String("Pharmacy"),
			},
		},
		TableName: aws.String("Profiles"),
	}

	svc := common.GetDynamoDbSession()

	_, err := svc.DeleteItem(input)

	if err != nil {
		panic(err)
	}
}
