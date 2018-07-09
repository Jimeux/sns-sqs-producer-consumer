package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/Jimeux/sns-sqs-producer-consumer/consumer"
	"github.com/Jimeux/sns-sqs-producer-consumer/producer"
	"os"
)

var (
	region   = os.Getenv("AWS_DEFAULT_REGION")
	queueURL = os.Getenv("QUEUE_URL")
	topicARN = os.Getenv("TOPIC_ARN")
)

func main() {
	awsSession, err := session.NewSession(&aws.Config{Region: &region})

	if err != nil {
		panic(err)
	}

	snsClient := sns.New(awsSession)
	sqsClient := sqs.New(awsSession)

	p := producer.New(snsClient, topicARN)
	c := consumer.New(sqsClient, queueURL)

	go p.Produce()
	c.Consume()
}
