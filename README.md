## SNS+SQSによる非同期サービス間通信

以下の環境変数を設定しておく。認証に関してはSpecifying Credentialsに参照(英語のみ)。

```
export AWS_ACCESS_KEY_ID="xxx"
export AWS_SECRET_ACCESS_KEY="xxx"
export AWS_DEFAULT_REGION="ap-northeast-1"
export TRACKING_TOPIC_ARN="arn:aws:sns:ap-northeast-1:xxx:TrackingTopic"
export DELIVERY_QUEUE_URL="https://sqs.ap-northeast-1.amazonaws.com/xxx/DeliveryQueue"
```
