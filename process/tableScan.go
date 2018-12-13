package main

import (
	"fmt"
	"github.com/DynamoDB/common"
	"github.com/DynamoDB/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func main() {
	itemsScan()
}

//El problema de esto es que recorre toda la tabla
func itemsScan() {
	svc := common.GetDynamoDbSession()

	filt := expression.Name("active").Equal(expression.Value(false))

	proj := expression.NamesList(expression.Name("vertical"), expression.Name("guid"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Profiles"),
	}

	result, err := svc.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, i := range result.Items {
		item := model.Profile{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Item: ", item)
		fmt.Println()
	}
}
