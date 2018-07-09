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

// AWS SDKに必要な環境変数を取得する。
// AWS_ACCESS_KEY_IDとAWS_SECRET_ACCESS_KEYはSDKに自動で認識される。
var (
	region   = os.Getenv("AWS_DEFAULT_REGION")
	queueURL = os.Getenv("DELIVERY_QUEUE_URL")
	topicARN = os.Getenv("TRACKING_TOPIC_ARN")
)

// AWSセッション、SNSクライント、SQSクライントを作成し、
// 非同期的にメッセージの発行・受信を行う。
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
