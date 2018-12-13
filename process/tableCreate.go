package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	tableCreate()
}

func tableCreate() {
	svc := common.GetDynamoDbSession()

	//HASH - partition key. That is used to distribute data storage
	//RANGE - sort key. That is used for sort data

	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Profiles"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("vertical"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("guid"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("vertical"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("guid"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("observations"), AttributeType: aws.String("S")},
		},

		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			{
				IndexName: aws.String("observationsindex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("vertical"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("observations"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("KEYS_ONLY"),
				},
			},
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
