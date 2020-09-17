package main

import (
	"bytes"
	"context"
	"encoding/json"
	"funcCountry/handler"

	"github.com/appmanch/go-commons/logging"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var logger = logging.GetLogger()

// Response ->
type Response events.APIGatewayProxyResponse

func main() {
	lambda.Start(CountryHandler)
}

//CountryHandler -> lambda handler
func CountryHandler(ctx context.Context) (Response, error) {

	finalList := handler.APIHandler()
	var buf bytes.Buffer
	body, err := json.Marshal(finalList)
	if err != nil {
		logger.ErrorF("Error During Unmarshaling Final Response")
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
