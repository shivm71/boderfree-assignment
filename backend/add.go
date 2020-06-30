package main
//imports
import (
   "encoding/json"
   "fmt"
   "github.com/aws/aws-sdk-go/aws"
   "github.com/aws/aws-sdk-go/service/dynamodb"
   "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
   "os"
)
// main add function
func add(eve Event) (string, error) {
   // mapping go attribute into dynamo attribute
   av, err := dynamodbattribute.MarshalMap(eve)

   if err != nil {
      fmt.Println("Got error marshalling new eve item:")
      fmt.Println(err.Error())
      os.Exit(1)
   }

   //passing dynamoattribute into dynamodb suitable query
   input := &dynamodb.PutItemInput{
      Item:      av,
      TableName: aws.String(tablename),
   }

   //passing query into dynamodb session
   _, err = db_session.PutItem(input)
   //handling error
   if err != nil {
      fmt.Println("Got error calling PutItem:")
      fmt.Println(err.Error())
      os.Exit(1)
   }
   //parsing go entity into json as string([]byte to string)
   body,err2 := json.Marshal(eve)
   //returning reponse with any error
   return string(body),err2


}
