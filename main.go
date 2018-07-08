package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"github.com/Jimeux/sns-sqs-tracker/consumer"
	"github.com/aws/aws-sdk-go/service/sqs"
	"fmt"
)

var (
	region   = os.Getenv("AWS_DEFAULT_REGION")
	queueURL = os.Getenv("DELIVERY_QUEUE_URL")
	topicARN = os.Getenv("TRACKING_TOPIC_ARN")
)

func main() {
	awsSession, err1 := session.NewSession(&aws.Config{Region: &region})

	if err1 != nil {
		panic(err1)
	}


/*	snsClient := sns.New(awsSession)

	subject := "Test Subject"
	message := "Test Message"

	_, err2 := snsClient.Publish(&sns.PublishInput{
		Subject:  &subject,
		Message:  &message,
		TopicArn: &topicARN,
	})

	if err2 != nil {
		fmt.Errorf("メッセージを発行できませんでした")
	}

	fmt.Println("Published: " + message)
*/
	//p := producer.New(snsClient, topicARN)

	sqsClient := sqs.New(awsSession)
	c := consumer.New(sqsClient, queueURL)

snsMessage, _ := c.ReceiveMessage()
message, _ := consumer.ParseMessageJSON(snsMessage)

_, err := sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
	QueueUrl:      &queueURL,
	ReceiptHandle: snsMessage.ReceiptHandle,
})

if err != nil {
	fmt.Errorf("メッセージを削除できませんでした：" + *snsMessage.ReceiptHandle)
}

fmt.Println("削除：" + message.Message) // 削除：Test Message

	//go p.Produce()
	//c.Consume()
}
