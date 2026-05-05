package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	sesTypes "github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	garageTypes "github.com/Steelyphil1/garage/types"
	"github.com/aws/aws-lambda-go/events"
)

func HandlePut(ctx context.Context, cfg garageTypes.GarageConfig, req garageTypes.GaragePutRequest) (*events.APIGatewayV2HTTPResponse, error) {
	err := req.Valid()
	if err != nil {
		return nil, err
	}

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("loading aws config: %w", err)
	}
	client := dynamodb.NewFromConfig(awsCfg)

	if req.State == "Open" {
		sendNotification(ctx, cfg)
	}

	now := time.Now()

	strNow := strconv.Itoa(int(now.Unix()))

	event := garageTypes.GarageEvent{
		EventTime:   now,
		GarageState: garageTypes.GarageState(req.State),
	}

	id := uuid.New().String()

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(cfg.TableName),
		Item: map[string]types.AttributeValue{
			"id":         &types.AttributeValueMemberS{Value: id},
			"entity":     &types.AttributeValueMemberS{Value: cfg.Partition},
			"event_time": &types.AttributeValueMemberN{Value: strNow},
			"status":     &types.AttributeValueMemberS{Value: req.State},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("putting garage state in dynamo: %w", err)
	}

	body, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("marshalling response: %w", err)
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(body),
	}, nil
}

func sendNotification(ctx context.Context, cfg garageTypes.GarageConfig) {
	awsCfg, _ := config.LoadDefaultConfig(ctx)
	client := sesv2.NewFromConfig(awsCfg)

	output, err := client.SendEmail(ctx, &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(cfg.EmailConfig.EmailFrom),
		Destination: &sesTypes.Destination{
			ToAddresses: []string{cfg.EmailConfig.EmailTo},
		},
		Content: &sesTypes.EmailContent{
			Simple: &sesTypes.Message{
				Subject: &sesTypes.Content{
					Data:    aws.String(cfg.EmailConfig.EmailSubject),
					Charset: aws.String("UTF-8"),
				},
				Body: &sesTypes.Body{
					Text: &sesTypes.Content{
						Data:    aws.String(cfg.EmailConfig.EmailBody),
						Charset: aws.String("UTF-8"),
					},
					Html: &sesTypes.Content{
						Data:    aws.String("<p>" + cfg.EmailConfig.EmailBody + "</p>"),
						Charset: aws.String("UTF-8"),
					},
				},
			},
		},
	})
	if err != nil {
		// Swallow error to continue with database transaction
		log.Println("email failed to send, continuing with db transaction")
	}

	log.Printf("Email sent. MessageID: %s", aws.ToString(output.MessageId))
}
