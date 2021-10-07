package main

import (
	"backend/apihandler"
	"bytes"
	"encoding/json"

	"go.appmanch.org/commons/logging"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var logger = logging.GetLogger()

// Response ->
type Response events.APIGatewayProxyResponse

func main() {
	lambda.Start(BackendHandler)
}

//BackendHandler -> lambda handler
func BackendHandler(request events.APIGatewayProxyRequest) (Response, error) {
	param1, found := request.QueryStringParameters["country"]
	if !found {
		logger.ErrorF("No QP Available")
	}
	finalList := apihandler.APIHandler(param1)
	var buf bytes.Buffer
	body, err := json.Marshal(finalList)
	if err != nil {
		logger.ErrorF("Error while Unmarshaling finaldata")
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"X-MyCompany-Func-Reply":      "hello-handler",
			"Access-Control-Allow-Origin": "*",
		},
	}
	return resp, nil
}
