package route

import (
	"context"
	"fmt"
	"net/http"

	garageHttp "github.com/Steelyphil1/garage/http"
	"github.com/Steelyphil1/garage/service"
	"github.com/Steelyphil1/garage/types"
	"github.com/aws/aws-lambda-go/events"
)

func RouteRequest(ctx context.Context, cfg types.GarageConfig, request types.BaseHTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	if request.Method == http.MethodGet {
		decodedReq, err := garageHttp.DecodeGetGarageStateRequest(ctx, request)
		if err != nil {
			return nil, err
		}

		events, err := service.HandleGet(ctx, cfg, *decodedReq)
		if err != nil {
			return nil, err
		}

		return garageHttp.PrepGetResponse(ctx, events)
	}
	if request.Method == http.MethodPut {
		decodedReq, err := garageHttp.DecodePutGarageStateRequest(ctx, request)
		if err != nil {
			return nil, err
		}

		return service.HandlePut(ctx, cfg, *decodedReq)
	}

	return nil, fmt.Errorf("invalid method")
}
