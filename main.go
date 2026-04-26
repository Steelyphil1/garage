package main

import (
	"context"
	"fmt"

	"github.com/Steelyphil1/garage/route"
	"github.com/Steelyphil1/garage/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	req := types.DeconstructAPIGatewayRequest(request)
	fmt.Printf("Decoded Request: %+v", req)

	resp, err := route.RouteRequest(ctx, req)
	if err != nil {
		return types.ErrorResponse(500, err.Error()), nil
	}

	return *resp, nil
}

func main() {
	lambda.Start(handler)
}
