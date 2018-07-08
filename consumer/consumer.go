package consumer

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"encoding/json"
	"fmt"
)

type Consumer struct {
	client   *sqs.SQS
	queueURL string
}

func New(client *sqs.SQS, queueURL string) *Consumer {
	return &Consumer{client, queueURL}
}

func (c *Consumer) Consume() {
	for {
		snsMessage, err := c.ReceiveMessage()

		if err != nil {
			fmt.Errorf("メッセージを受信できませんでした")
		} else if snsMessage == nil {
			fmt.Println("メッセージはありませんでした")
		} else {
			c.processSNSMessage(snsMessage)
		}
	}
}

func (c *Consumer) DeleteMessage(handle string, message string) {
	_, err := c.client.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &c.queueURL,
		ReceiptHandle: &handle,
	})

	if err != nil {
		fmt.Errorf("メッセージを削除できませんでした：" + handle)
		return
	}

	fmt.Println("削除：" + message)
}

func (c *Consumer) ReceiveMessage() (*sqs.Message, error) {
	result, err := c.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &c.queueURL,
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

func (c *Consumer) processSNSMessage(snsMessage *sqs.Message) {
	if message, err := ParseMessageJSON(snsMessage); err != nil {
		fmt.Errorf("メッセージJSONをパースできませんでした")
	} else {
		fmt.Println("受信：" + message.Message)
		c.DeleteMessage(*snsMessage.ReceiptHandle, message.Message)
	}
}

func ParseMessageJSON(snsMessage *sqs.Message) (SnsMessage, error) {
	var message SnsMessage
	err := json.Unmarshal([]byte(*snsMessage.Body), &message)
	return message, err
}
