package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/DynamoDB/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	itemGet()
}

func itemGet() {
	svc := common.GetDynamoDbSession()

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Profiles"),
		Key: map[string]*dynamodb.AttributeValue{
			"vertical": {
				S: aws.String("pharmacy"),
			},
			"guid": {
				S: aws.String("b5e6d907-a953-4a79-be38-0f97b4ca29a4"),
			},
		},
	})

	if err != nil {
		panic(err.Error())
	}

	item := model.Profile{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if item.GUID == "" {
		fmt.Println("Could not find")
		return
	}

	fmt.Println("Found item:")
	fmt.Println("Vertical:", item.Vertical)
	fmt.Println("GUID:  ", item.GUID)
	fmt.Println("Active: ", item.Active)
	fmt.Println("Observations:  ", item.Observations)
	fmt.Println("RawJsonData:  ", item.RawJsonData)
	fmt.Println("\n\nresult.Item:  ", result.Item)

	fmt.Println(item)
}
