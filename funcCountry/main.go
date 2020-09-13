package main

import (
	"context"
	"funcCountry/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(CountryHandler)
}

//CountryHandler -> lambda handler
func CountryHandler(ctx context.Context) ([]string, error) {

	finalList := handler.APIHandler()
	return finalList, nil
}
