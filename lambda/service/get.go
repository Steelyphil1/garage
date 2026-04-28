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

func HandleGet(ctx context.Context, cfg garageTypes.GarageConfig, req garageTypes.GarageGetRequest) (*garageTypes.GarageEvent, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("loading aws config: %w", err)
	}
	client := dynamodb.NewFromConfig(awsCfg)

	out, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(cfg.TableName),
		IndexName:              aws.String("entity-event_time-index"),
		KeyConditionExpression: aws.String("entity = :entity"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":entity": &types.AttributeValueMemberS{Value: cfg.Partition},
		},
		ScanIndexForward: aws.Bool(false),
		Limit:            aws.Int32(1),
	})
	if err != nil {
		return nil, err
	}

	if len(out.Items) == 0 {
		return nil, fmt.Errorf("no events found")
	}

	item := out.Items[0]

	status, err := extractStringFromItem(item, "status")
	if err != nil {
		return nil, err
	}
	state := garageTypes.GarageState(status)

	eventTime, err := extractNumberFromItem(item, "event_time")
	if err != nil {
		return nil, err
	}

	return &garageTypes.GarageEvent{
		GarageState: state,
		EventTime:   time.Unix(int64(eventTime), 0),
	}, nil
}

func extractStringFromItem(item map[string]types.AttributeValue, key string) (string, error) {
	val, ok := item[key]
	if !ok {
		return "", fmt.Errorf("%s field not found in item", key)
	}
	member, ok := val.(*types.AttributeValueMemberS)
	if !ok {
		return "", fmt.Errorf("%s field is not a string type", key)
	}

	return member.Value, nil
}

func extractNumberFromItem(item map[string]types.AttributeValue, key string) (int64, error) {
	val, ok := item[key]
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
