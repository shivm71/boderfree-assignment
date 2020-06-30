package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func list(uid string)(string,error) {
	//blocking multi call of golang bug
	if uid == ""{

		return "Invalid/Empty userid",nil
	}

	//making filter according to need
	filter := expression.Name("userid").Equal(expression.Value(uid))
	// turning filter into expression to scan recognizable entity
	// can use projection to get specific colunm
	expr, err := expression.NewBuilder().WithFilter(filter).Build()

	// handling error for expression(query) building correctness
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return "", err
	}

	//mapping all condition into a single query
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tablename),
	}

	//passing query into dynamodb session and handling errors
	result, err := db_session.Scan(params)
	response_list := []Event{}
	for _, i := range result.Items {

		//mapping reference declaration
		response_event := Event{}

		//maping dynamodb reponse into go EVENT struct
		err = dynamodbattribute.UnmarshalMap(i, &response_event)

		if err != nil {
			return "", err
		}
		response_list = append(response_list, response_event)
	}
	body, err := json.Marshal(&Eventlist{
		Eventlist: response_list,
	})
	return string(body), err
}





