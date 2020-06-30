package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
	"os"
)

//Global variable declaration
var (
	res           = os.Getenv("TABLE_NAME")
	tablename     = "event_name"
	err           error
	event         = Event{}
	deleterequest = Deleteid{}
	db_session    *dynamodb.DynamoDB

	)

//initialization
func init(){

	//getting value of region from environment
	reg := os.Getenv("REGION")

	// initialization aws session with region and credentials(shared)
	aws_session,_ := session.NewSession(&aws.Config{

		Region:      aws.String(reg),
	})

	//starting dyanamodb session with aws session
	db_session = dynamodb.New(aws_session)
}

//lambda handler function
func route(ctx context.Context ,req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {

	//Restful diversion to various function depending type of REQUEST Performed
	switch method := req.HTTPMethod; method{

	//GET
	case "GET":
		user_id := req.QueryStringParameters["userid"]
		//returning list
		res,err = list(user_id)

	//POST
	case "POST":
		// Generate unique id for every event
		id := uuid.NewV4()
		event.Eventid=id.String()

		_ = json.Unmarshal([]byte(req.Body),&event)

		res,err = add(event)

	//PUT
	case "PUT":

		_ = json.Unmarshal([]byte(req.Body), &event)
		res,err = add(event)

	//DELETE
	case "DELETE":
		_ = json.Unmarshal([]byte(req.Body), &deleterequest)
		res,err = delete_event(deleterequest)

	//OPTIONS
	case "OPTIONS":
		res,err = "",nil
	
	// Default to Block not allowed methods && can also be handle from api gateway.
	default:
		return events.APIGatewayProxyResponse{
			Body: "Method NOT Allowed",
			StatusCode: 405,
		},err

	}

	//Mapping response into APIGatewayProxyResponse
	response:= events.APIGatewayProxyResponse{
		Body: res,
		StatusCode: 200,
		Headers: make(map[string]string),
	}

	//Adding "Access-Control-Allow-Origin" Header to allow CORS Policy
	response.Headers["Access-Control-Allow-Origin"] = "*"

	// prefly checking
	response.Headers["Access-Control-Allow-Methods"] = "GET,POST,PUT,DELETE,OPTIONS"
	//response.Headers["Content-Type"] = "application/json"

	//Return response to user with error
	return response,err
}

func main(){
	//lambda handler
	lambda.Start(route)
}
