package producer

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"strconv"
	"time"
	"fmt"
)

type Producer struct {
	client   *sns.SNS
	topicARN string
}

func New(client *sns.SNS, topicARN string) *Producer {
	return &Producer{client, topicARN}
}

// 毎秒1メッセージを10回まで発行する
func (p *Producer) Produce() {
	for i := 0; i < 10; i++ {
		msg := "Message " + strconv.Itoa(i)
		p.publishSns("Test.Event", msg)
		time.Sleep(1 * time.Second)
	}
}

// AWS SDKを利用してSNSへメッセージを発行する
func (p *Producer) publishSns(subject string, message string) (string, error) {
	resp, err := p.client.Publish(&sns.PublishInput{
		Subject:  &subject,
		Message:  &message,
		TopicArn: &p.topicARN,
	})

	if err != nil {
		return "", err
	}

	fmt.Println("発行：" + message)
	return *resp.MessageId, nil
}
