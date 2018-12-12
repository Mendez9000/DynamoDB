package main

import (
	"dynamodb_2/common"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	tableCreate()
}

func tableCreate() {
	svc := common.GetDynamoDbSession()

	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Profiles"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("guid"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("vertical"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("guid"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("vertical"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(100),
		},
	}

	resp, err := svc.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
