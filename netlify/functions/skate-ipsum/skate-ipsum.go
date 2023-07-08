package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/madflow/skate-ipsum/generator"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Running skate ipsum function")

	paragraphs := generator.IpsumJson(10)
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json; charset=utf-8"},
		Body:            paragraphs,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
