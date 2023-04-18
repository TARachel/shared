package sqsconsumer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/panjf2000/ants/v2"
)

const (
	MaxNumberOfMessages = 10
)

type Client struct {
	sqsClient *sqs.Client
	queue     string
	pool      *ants.Pool
}

func NewClient(ctx context.Context, queue string, poolSize int) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	pool, err := ants.NewPool(poolSize, ants.WithPreAlloc(true))
	if err != nil {
		return nil, err
	}

	sqsClient := sqs.NewFromConfig(cfg)

	return &Client{
		sqsClient: sqsClient,
		queue:     queue,
		pool:      pool,
	}, nil
}

func (c *Client) StartConsuming(ctx context.Context, handler func(msg string) error) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stopping consumer for queue %s\n", c.queue)
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
				if err := c.pool.Submit(func() {
					if err := handler(*msg.Body); err != nil {
						fmt.Println("error handling message: ", err)
						return
					}

					_, err := c.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
						QueueUrl:      &c.queue,
						ReceiptHandle: msg.ReceiptHandle,
					})
					if err != nil {
						fmt.Println("error deleting message: ", err)
					}

				}); err != nil {
					fmt.Println("error submitting message to pool: ", err)
					return err
				}
			}
		}
	}
}
