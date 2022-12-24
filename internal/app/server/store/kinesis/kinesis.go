package kinesis

import (
	"fmt"

	"github.com/FindHotel/emspy/internal/app/server/store"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type KinesisStore struct {
	streamName *string
	kinesis    *kinesis.Kinesis
}

func New(awsSession *session.Session, streamName string) (store.Store, error) {
	//s := session.New(&aws.Config{Region: aws.String(*region)})
	kc := kinesis.New(awsSession)
	stream := aws.String(streamName)

	out, err := kc.CreateStream(&kinesis.CreateStreamInput{
		ShardCount: aws.Int64(1),
		StreamName: stream,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", out)

	if err := kc.WaitUntilStreamExists(&kinesis.DescribeStreamInput{StreamName: stream}); err != nil {
		return nil, err
	}

	streams, err := kc.DescribeStream(&kinesis.DescribeStreamInput{StreamName: stream})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", streams)
	return &KinesisStore{streamName: stream, kinesis: kc}, nil
}

func (s *KinesisStore) InsertWebhook(record interface{}) error {
	input := record.([]byte)

	putOutput, err := s.kinesis.PutRecord(&kinesis.PutRecordInput{
		Data:         input,
		StreamName:   s.streamName,
		PartitionKey: aws.String("shortcut"),
	})
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", putOutput)

	return nil
}
