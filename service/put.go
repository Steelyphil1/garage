package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-lambda-go/events"
	garageTypes "github.com/phillipbay/garage/types"
)

func HandlePut(ctx context.Context, req garageTypes.GaragePutRequest) (*events.APIGatewayV2HTTPResponse, error) {
	fmt.Printf("In HandlePut with: %+v", req)
	err := req.Valid()
	if err != nil {
		return nil, err
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("loading aws config")
	}
	client := dynamodb.NewFromConfig(cfg)

	now := time.Now().Unix()

	_, err = client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String("GarageStatus"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "garage"},
		},
		UpdateExpression: new("SET #s = :status, event_time = :time"),
		ExpressionAttributeNames: map[string]string{
			"#s": "status",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":status": &types.AttributeValueMemberS{Value: req.State},
			":time":   &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", now)},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("updating garage state in dynamo: %w", err)
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}
