package consumer

import "time"

// SNSトピックに発行されたメッセージ。件名と本文に加え、
// JSONドキュメントのメッセージに関するメタデータが含まれる
type SnsMessage struct {
	Type             string               // Notification
	MessageId        string               // 63a3f6b6-d533-4a47-aef9-fcf5cf758c76
	TopicArn         string               // arn:aws:sns:us-west-2:123456789012:MyTopic
	Subject          string               // Testing publish to subscribed queues
	Message          string               // Hello world!
	Timestamp        time.Time            // 2012-03-29T05:12:16.901Z
	SignatureVersion int `json:",string"` // 1
	Signature        string               // EXAMPLEnTrFPa37tnVO0FF9Iau3MGzjlJLRfySEoWz4uZHSj6ycK4ph71Zmdv0NtJ4dC/El9FOGp3VuvchpaTraNHWhhq/OsN1HVz20zxmF9b88R8GtqjfKB5woZZmz87HiM6CYDTo3l7LMwFT4VU7ELtyaBBafhPTg9O5CnKkg=
	SigningCertURL   string               // https://sns.us-west-2.amazonaws.com/SimpleNotificationService-f3ecfb7224c7233fe7bb5f59f96de52f.pem
	UnsubscribeURL   string               // https://sns.us-west-2.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-west-2:123456789012:MyTopic:c7fe3a54-ab0e-4ec2-88e0-db410a0f2bee
}
