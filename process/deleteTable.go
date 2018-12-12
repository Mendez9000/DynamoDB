package main

import (
	"dynamodb_2/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	deleteTable()
}

func deleteTable() error {
	svc := common.GetDynamoDbSession()
	_, err := svc.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String("Profiles"),
	})
	if err != nil {
		return err
	}
	err = svc.WaitUntilTableNotExistsWithContext(
		aws.BackgroundContext(),
		&dynamodb.DescribeTableInput{
			TableName: aws.String("Profiles"),
		},
	)
	return err
}
