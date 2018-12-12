package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	itemUpdateExtraAttribute()
}

func itemUpdateExtraAttribute() {
	svc := common.GetDynamoDbSession()

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"guid": {
				S: aws.String("b5e6d907-a953-4a79-be38-0f97b4ca29a4"),
			},
			"vertical": {
				S: aws.String("Pharmacy"),
			},
		},
		TableName:        aws.String("Profiles"),
		UpdateExpression: aws.String("set new_attribute = :r"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String("New Attr attached"),
			},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully add")
}
