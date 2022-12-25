package kinesis

import (
	"fmt"

	"github.com/FindHotel/emspy/internal/app/store"
	"github.com/FindHotel/emspy/pkg/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type KinesisStore struct {
	streamName *string
	kinesis    *kinesis.Kinesis
	logger     logger.Logger
}

func New(awsSession *session.Session, streamName string, logger logger.Logger) (store.Store, error) {
	kc := kinesis.New(awsSession)
	stream := aws.String(streamName)

	if err := kc.WaitUntilStreamExists(&kinesis.DescribeStreamInput{StreamName: stream}); err != nil {
		return nil, err
	}

	return &KinesisStore{streamName: stream, kinesis: kc, logger: logger.Named("kinesis store")}, nil
}

func (s *KinesisStore) InsertWebhook(source string, record interface{}) error {
	input := record.([]byte)

	_, err := s.kinesis.PutRecord(&kinesis.PutRecordInput{
		Data:         []byte(fmt.Sprintf(`{"webhook_source":"%s","webhook_data":%s}`, source, input)),
		StreamName:   s.streamName,
		PartitionKey: aws.String("shortcut"),
	})

	return err
}
