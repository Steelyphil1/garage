package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/phillipbay/garage/types"
)

func DecodePutGarageStateRequest(_ context.Context, r types.BaseHTTPRequest) (*types.GaragePutRequest, error) {
	fmt.Println("IN DECODE PUT")
	request := types.GaragePutRequest{}

	if r.Body != "" {
		if err := json.Unmarshal([]byte(r.Body), &request); err != nil {
			return nil, fmt.Errorf("decoding request")
		}
	}

	return &request, nil
}

func DecodeGetGarageStateRequest(_ context.Context, r types.BaseHTTPRequest) (*types.GarageGetRequest, error) {
	fmt.Println("IN DECODE GET")
	return &types.GarageGetRequest{}, nil
}

func PrepGetResponse(ctx context.Context, state types.GarageEvent) (*events.APIGatewayV2HTTPResponse, error) {
	fmt.Printf("In PrepGetResponse with: %+v", state)
	responseBody, err := json.Marshal(state)
	if err != nil {
		return nil, fmt.Errorf("marshalling state")
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(responseBody),
	}, nil
}
