package sqsconsumer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqstype "github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const (
	MaxNumberOfMessages = 10
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

func (c *Client) StartConsuming(ctx context.Context, handler func(msg string) error) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			resp, err := c.sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
				QueueUrl:            &c.queue,
				MaxNumberOfMessages: MaxNumberOfMessages,
			})
			if err != nil {
				return err
			}

			for _, msg := range resp.Messages {
				// start a go routine to handle the message
				go func(msg sqstype.Message) {
					if err := handler(*msg.Body); err != nil {
						fmt.Println("error handling message: ", err)
						return
					}

					// delete the message from the queue
					_, err := c.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
						QueueUrl:      &c.queue,
						ReceiptHandle: msg.ReceiptHandle,
					})
					if err != nil {
						fmt.Println("error deleting message: ", err)
					}
				}(msg)
			}
		}
	}
}
