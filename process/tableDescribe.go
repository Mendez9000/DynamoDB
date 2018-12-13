package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	tableDescribe()
}

func tableDescribe() {
	svc := common.GetDynamoDbSession()

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("Profiles"),
	}

	resp, err := svc.DescribeTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
