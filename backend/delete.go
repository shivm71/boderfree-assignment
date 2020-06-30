package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)


//defination delete_event
func delete_event(del_req Deleteid) (string,error){

	//making query with go attributes
	input := &dynamodb.DeleteItemInput{

		Key: map[string]*dynamodb.AttributeValue{

			"eventid": {
				S: aws.String(del_req.Id),
			},
		},
		TableName: aws.String(tablename),
	}

	//passing query into dynamodb session and handling errors
	_, err = db_session.DeleteItem(input)

	return "Successfully Deleted entry",err

}
