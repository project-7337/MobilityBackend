package main

import (
	"backend/apihandler"
	"backend/models"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(BackendHandler)
}

//BackendHandler -> lambda handler
func BackendHandler(ctx context.Context) (models.FinalData, error) {

	finalList := apihandler.APIHandler()
	return finalList, nil
}
