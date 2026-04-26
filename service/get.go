package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	garageTypes "github.com/phillipbay/garage/types"
)

func HandleGet(ctx context.Context, req garageTypes.GarageGetRequest) (*garageTypes.GarageEvent, error) {
	fmt.Printf("In HandleGet with: %+v", req)

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("loading aws config")
	}
	client := dynamodb.NewFromConfig(cfg)

	out, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("GarageStatus"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "garage"},
		},
	})
	if err != nil {
		return nil, err
	}

	state := garageTypes.GarageState(
		out.Item["status"].(*types.AttributeValueMemberS).Value,
	)

	tsStr := out.Item["event_time"].(*types.AttributeValueMemberN).Value

	tsInt, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &garageTypes.GarageEvent{
		GarageState: state,
		EventTime:   time.Unix(tsInt, 0),
	}, nil
}
