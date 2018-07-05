package topic

import (
	"github.com/aws/aws-sdk-go/service/sns"
)

type Client struct {
	topicARN string
	client   *sns.SNS
}

func New(topicARN string, client *sns.SNS) *Client {
	return &Client{topicARN, client}
}

func (t *Client) PublishMessage(subject string, message string) (*string, error) {
	resp, err := t.client.Publish(&sns.PublishInput{
		Subject:  &subject,
		Message:  &message,
		TopicArn: &t.topicARN,
	})

	if err != nil {
		return nil, err
	} else {
		return resp.MessageId, nil
	}
}
