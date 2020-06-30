package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

)


func route(_ context.Context ,req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
	var temp []byte
	var err error
	temp,err = json.Marshal(req)
	res := string(temp)
	//Mapping response into APIGatewayProxyResponse
	response:= events.APIGatewayProxyResponse{
		Body: res,
		StatusCode: 200,
		Headers: make(map[string]string),
	}
	//Adding "Access-Control-Allow-Origin" Header to allow CORS Policy
	response.Headers["Access-Control-Allow-Origin"] = "*"

	//Return response to user with error
	return response,err
}

func main(){
	//lambda handler
	lambda.Start(route)
}
