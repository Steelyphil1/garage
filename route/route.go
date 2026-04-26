package route

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	garageHttp "github.com/phillipbay/garage/http"
	"github.com/phillipbay/garage/service"
	"github.com/phillipbay/garage/types"
)

func RouteRequest(ctx context.Context, request types.BaseHTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	fmt.Println("In RouteRequest")
	if request.Method == http.MethodGet {
		fmt.Println("In ROUTEGET")
		decodedReq, err := garageHttp.DecodeGetGarageStateRequest(ctx, request)
		if err != nil {
			return nil, err
		}

		fmt.Println("About to handle get")
		state, err := service.HandleGet(ctx, *decodedReq)
		if err != nil {
			return nil, err
		}

		return garageHttp.PrepGetResponse(ctx, *state)
	}
	if request.Method == http.MethodPut {
		fmt.Println("In ROUTEPUT")
		decodedReq, err := garageHttp.DecodePutGarageStateRequest(ctx, request)
		if err != nil {
			return nil, err
		}

		fmt.Println("About to handle put")
		return service.HandlePut(ctx, *decodedReq)
	}

	return nil, fmt.Errorf("invalid method")
}
