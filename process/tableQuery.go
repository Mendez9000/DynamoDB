package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	findAllPharmacies()
	findAllActives()
}

//Busca por la clave princial
func findAllPharmacies() {
	fmt.Println("findAllPharmacies")
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

//Busca por una clave secundaria
func findAllActives() {
	fmt.Println("\n\nfindAllActives")
	svc := common.GetDynamoDbSession()
	input := &dynamodb.QueryInput{
		TableName: aws.String("Profiles"),
		IndexName: aws.String("observationsindex"),

		KeyConditionExpression: aws.String("observations = :observations and vertical= :vertical"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":observations": {
				S: aws.String("This user is"),
			},
			":vertical": {
				S: aws.String("supermarket"),
			},
		},
		ScanIndexForward: aws.Bool(false), // sort by sort key in DESC order
		Limit:            aws.Int64(15),
	}

	result, err := svc.Query(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
