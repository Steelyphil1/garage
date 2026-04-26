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

	garageTypes "github.com/Steelyphil1/garage/types"
)

func HandleGet(ctx context.Context, req garageTypes.GarageGetRequest) (*garageTypes.GarageEvent, error) {
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

	status, err := extractStringFromDynamoOutput(*out, "status")
	if err != nil {
		return nil, err
	}
	state := garageTypes.GarageState(status)

	eventTime, err := extractNumberFromDynamoOutput(*out, "event_time")
	if err != nil {
		return nil, err
	}

	return &garageTypes.GarageEvent{
		GarageState: state,
		EventTime:   time.Unix(int64(eventTime), 0),
	}, nil
}

func extractStringFromDynamoOutput(out dynamodb.GetItemOutput, key string) (string, error) {
	val, ok := out.Item[key]
	if !ok {
		return "", fmt.Errorf("%s field not found in item", key)
	}
	member, ok := val.(*types.AttributeValueMemberS)
	if !ok {
		return "", fmt.Errorf("%s field is not a string type", key)
	}

	return member.Value, nil
}

func extractNumberFromDynamoOutput(out dynamodb.GetItemOutput, key string) (int64, error) {
	val, ok := out.Item[key]
	if !ok {
		return -1, fmt.Errorf("%s field not found in item", key)
	}
	member, ok := val.(*types.AttributeValueMemberN)
	if !ok {
		return -1, fmt.Errorf("%s field is not a number type", key)
	}

	num, err := strconv.ParseInt(member.Value, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("unable to parse int from dynamo")
	}

	return num, nil
}
