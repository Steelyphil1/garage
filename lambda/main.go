package main

import (
	"context"
	"net/http"
	"os"

	"github.com/Steelyphil1/garage/route"
	"github.com/Steelyphil1/garage/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	cfg, err := buildConfig()
	if err != nil {
		return types.ErrorResponse(500, err.Error()), nil
	}

	req := types.DeconstructAPIGatewayRequest(request)

	resp, err := route.RouteRequest(ctx, *cfg, req)
	if err != nil {
		return types.ErrorResponse(http.StatusBadRequest, err.Error()), nil
	}

	return *resp, nil
}

func main() {
	lambda.Start(handler)
}

func buildConfig() (*types.GarageConfig, error) {
	_ = godotenv.Load()

	return &types.GarageConfig{
		TableName: os.Getenv("DYNAMO_TABLE_NAME"),
		Partition: os.Getenv("DYNAMO_PARTITION_KEY"),
	}, nil
}
