package main

import (
	"github.com/Jimeux/sns-sqs-tracker/queue"
	"github.com/Jimeux/sns-sqs-tracker/topic"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/Jimeux/sns-sqs-tracker/events"
	"os"
	"fmt"
	"encoding/json"
)

var (
	region           = os.Getenv("AWS_DEFAULT_REGION")
	deliveryQueueUrl = os.Getenv("DELIVERY_QUEUE_URL")
)

func main() {
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		panic(err)
	}

	q := queue.New(deliveryQueueUrl, sqs.New(awsSession))

	consumeEvent(q)
}

func consumeEvent(q *queue.Client) {
	for {
		snsMessage, err := q.ReceiveMessage()

		if err != nil {
			fmt.Errorf("problem consuming event")
		} else if snsMessage != nil {
			if message, err := parseMessage(snsMessage); err != nil {
				fmt.Errorf("problem parsing event")
			} else {
				processMessage(message)
				q.DeleteMessage(snsMessage.ReceiptHandle)
			}
		}
	}
}

func processMessage(message topic.Message) {
	switch message.Subject {
	case events.TrackImpressionEvent:
		fmt.Println(message.Subject, message.Message)
	case events.TrackConversionEvent:
		fmt.Println(message.Subject, message.Message)
	default:
		fmt.Errorf("warning: unsupported message type %s", message.Subject)
	}
}

func parseMessage(snsMessage *sqs.Message) (topic.Message, error) {
	var message topic.Message
	err := json.Unmarshal([]byte(*snsMessage.Body), &message)
	return message, err
}
