package sqsproducer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Client struct {
	sqsClient *sqs.Client
	queue     string
}

func NewClient(ctx context.Context, queue string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	sqsClient := sqs.NewFromConfig(cfg)
	return &Client{
		sqsClient: sqsClient,
		queue:     queue,
	}, nil
}

func (c *Client) SendMessage(ctx context.Context, msg string) error {
	_, err := c.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: &msg,
		QueueUrl:    &c.queue,
	})

	return err
}
