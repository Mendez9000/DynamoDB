package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"io/ioutil"
)

func main() {
	itemUpdateMaxContent()
}

func itemUpdateMaxContent() {
	svc := common.GetDynamoDbSession()

	jsonFile, err := ioutil.ReadFile("/home/sergio/go/src/github.com/DynamoDB/data/itemUpdateMaxContent.json")

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
				S: aws.String(string(jsonFile)),
			},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	_, err = svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated")
}
