package queue

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

type Client struct {
	queueURL string
	client *sqs.SQS
}

func New(queueURL string, client *sqs.SQS) *Client {
	return &Client{queueURL, client}
}

func (q *Client) SendMessage(body string) (*string, error) {
	result, err := q.client.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody: aws.String(body),
		QueueUrl:    &q.queueURL,
	})

	if err != nil {
		return nil, err
	}

	return result.MessageId, nil
}

func (q *Client) ReceiveMessage() (*sqs.Message, error) {
	result, err := q.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &q.queueURL,
		AttributeNames: aws.StringSlice([]string{
			sqs.MessageSystemAttributeNameSentTimestamp,
		}),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			sqs.QueueAttributeNameAll,
		}),
		WaitTimeSeconds: aws.Int64(10),
	})

	if err != nil {
		return nil, err
	} else if len(result.Messages) > 0 {
		return result.Messages[0], nil
	} else {
		return nil, nil
	}
}

func (q *Client) ReceiveMessages(limit int) (*[]*sqs.Message, error) {
	result, err := q.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &q.queueURL,
		AttributeNames: aws.StringSlice([]string{
			sqs.MessageSystemAttributeNameSentTimestamp,
		}),
		MaxNumberOfMessages: aws.Int64(int64(limit)),
		MessageAttributeNames: aws.StringSlice([]string{
			sqs.QueueAttributeNameAll,
		}),
		WaitTimeSeconds: aws.Int64(10),
	})

	if err != nil {
		return nil, err
	} else if len(result.Messages) > 0 {
		return &result.Messages, nil
	} else {
		return nil, nil
	}
}

func (q *Client) DeleteMessage(handle *string) {
	_, err := q.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &q.queueURL,
		ReceiptHandle: handle,
	})

	if err != nil {
		fmt.Println("Delete Error", err)
		return
	}

	fmt.Println("Message Deleted")
}
