package route

import (
	"context"
	"fmt"
	"net/http"

	garageHttp "github.com/Steelyphil1/garage/lambda/http"
	"github.com/Steelyphil1/garage/lambda/service"
	"github.com/Steelyphil1/garage/lambda/types"
	"github.com/aws/aws-lambda-go/events"
)

func RouteRequest(ctx context.Context, request types.BaseHTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	if request.Method == http.MethodGet {
		decodedReq, err := garageHttp.DecodeGetGarageStateRequest(ctx, request)
		if err != nil {
			return nil, err
		}

		state, err := service.HandleGet(ctx, *decodedReq)
		if err != nil {
			return nil, err
		}

		return garageHttp.PrepGetResponse(ctx, *state)
	}
	if request.Method == http.MethodPut {
		decodedReq, err := garageHttp.DecodePutGarageStateRequest(ctx, request)
		if err != nil {
			return nil, err
		}

		return service.HandlePut(ctx, *decodedReq)
	}

	return nil, fmt.Errorf("invalid method")
}
