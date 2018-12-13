package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	findAllPharmacies()
}

//El problema de esto es que recorre toda la tabla
func findAllPharmacies() {
	svc := common.GetDynamoDbSession()

	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String("pharmacy"),
			},
		},
		KeyConditionExpression: aws.String("vertical = :v1"),
		ProjectionExpression:   aws.String("guid"),
		TableName:              aws.String("Profiles"),
	}

	result, err := svc.Query(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
