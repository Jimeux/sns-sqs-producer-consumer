package consumer

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"encoding/json"
	"fmt"
)

const (
	LongPollingWaitTimeSeconds = 10
	MaxMessages                = 1
)

type Consumer struct {
	client   *sqs.SQS
	queueURL string
}

func New(client *sqs.SQS, queueURL string) *Consumer {
	return &Consumer{client, queueURL}
}

// 無限ループでメッセージを受信する
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

// AWS SDKを利用してSQSキューからメッセージを削除する
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

// AWS SDKを利用してSQSキューからメッセージを一つ受信する
func (c *Consumer) ReceiveMessage() (*sqs.Message, error) {
	result, err := c.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &c.queueURL,
		AttributeNames: aws.StringSlice([]string{
			sqs.MessageSystemAttributeNameSentTimestamp,
		}),
		MaxNumberOfMessages: aws.Int64(MaxMessages),
		MessageAttributeNames: aws.StringSlice([]string{
			sqs.QueueAttributeNameAll,
		}),
		// 一秒以上だとロングポーリングが有効にされる
		WaitTimeSeconds: aws.Int64(LongPollingWaitTimeSeconds),
	})

	if err != nil {
		return nil, err
	} else if len(result.Messages) > 0 {
		return result.Messages[0], nil
	} else {
		return nil, nil
	}
}

// SQSキューから受信されたSNSメッセージをパースしキューから削除する
func (c *Consumer) processSNSMessage(snsMessage *sqs.Message) {
	if message, err := ParseMessageJSON(snsMessage); err != nil {
		fmt.Errorf("メッセージJSONをパースできませんでした")
	} else {
		fmt.Println("受信：" + message.Message)
		c.DeleteMessage(*snsMessage.ReceiptHandle, message.Message)
	}
}

// SNSのメッセージをJSONからパースする
func ParseMessageJSON(snsMessage *sqs.Message) (SnsMessage, error) {
	var message SnsMessage
	err := json.Unmarshal([]byte(*snsMessage.Body), &message)
	return message, err
}
