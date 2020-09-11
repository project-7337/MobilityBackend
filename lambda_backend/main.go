package main

import (
	"backend/apihandler"
	"backend/config"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(BackendHandler)
}

//BackendHandler -> lambda handler
func BackendHandler(ctx context.Context) ([]string, error) {
	var configs config.Conf
	configs.ReadConf()

	finalList := apihandler.APIHandler(configs)
	return finalList, nil
}
