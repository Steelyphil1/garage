package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type GarageState string

const (
	GarageStateOpen   GarageState = "Open"
	GarageStateClosed GarageState = "Closed"
)

type GarageGetRequest struct {
	// Intentionally empty
}

type GaragePutRequest struct {
	State string `json:"state"`
}

type GarageEvent struct {
	EventTime   time.Time   `json:"event_time"`
	GarageState GarageState `json:"garage_state"`
}

func (r *GaragePutRequest) Valid() error {
	if r.State != string(GarageStateOpen) && r.State != string(GarageStateClosed) {
		return fmt.Errorf("invalid garage state")
	}

	return nil
}

type BaseHTTPRequest struct {
	Method string
	Path   string
	Body   string
}

func ErrorResponse(statusCode int, message string) events.APIGatewayV2HTTPResponse {
	fmt.Println("IN ERROR RESPONSE")
	body, _ := json.Marshal(map[string]string{"error": message})
	return events.APIGatewayV2HTTPResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}
}
