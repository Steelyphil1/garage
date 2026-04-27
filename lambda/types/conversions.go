package types

import "github.com/aws/aws-lambda-go/events"

func DeconstructAPIGatewayRequest(r events.APIGatewayV2HTTPRequest) BaseHTTPRequest {
	return BaseHTTPRequest{
		Method: r.RequestContext.HTTP.Method,
		Path:   r.RequestContext.HTTP.Path,
		Body:   r.Body,
	}
}
