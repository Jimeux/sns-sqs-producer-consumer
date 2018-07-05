package main

import (
	"github.com/Jimeux/sns-sqs-tracker/topic"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"math/rand"
	"strconv"
	"time"
	"os"
	"github.com/Jimeux/sns-sqs-tracker/events"
)

var (
	region           = os.Getenv("AWS_DEFAULT_REGION")
	trackingTopicArn = os.Getenv("TRACKING_TOPIC_ARN")
)

func main() {
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		panic(err)
	}

	t := topic.New(trackingTopicArn, sns.New(awsSession))

	publishEvent(t)
}

func publishEvent(t *topic.Client) {
	for {
		n := rand.Intn(100)
		if n%2 == 0 {
			t.PublishMessage(events.TrackImpressionEvent, "ImpTest "+strconv.Itoa(n))
		} else {
			t.PublishMessage(events.TrackConversionEvent, "ImpTest "+strconv.Itoa(n))
		}

		time.Sleep(3 * time.Second)
	}
}
