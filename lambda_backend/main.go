package main

import (
	"backend/apihandler"
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response ->
type Response events.APIGatewayProxyResponse

func main() {
	lambda.Start(BackendHandler)
}

//BackendHandler -> lambda handler
func BackendHandler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	param1 := request.QueryStringParameters["country"]
	finalList := apihandler.APIHandler(param1)
	var buf bytes.Buffer
	body, err := json.Marshal(finalList)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}
	return resp, nil
}
