package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/phillipbay/garage/route"
	"github.com/phillipbay/garage/types"
)

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	req := types.DeconstructAPIGatewayRequest(request)
	fmt.Printf("Decoded Request: %+v", req)

	resp, err := route.RouteRequest(ctx, req)
	if err != nil {
		return types.ErrorResponse(500, "BOOM"), err
	}

	return *resp, err
}

func main() {
	lambda.Start(handler)
}
