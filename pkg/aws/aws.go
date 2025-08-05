package aws

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"git.inet.co.th/ekyc-platform-backend/pkg/requests"
)

const packageName = "aws"

type Client struct {
	http            *requests.HttpClient
	log             *logrus.Entry
	tracer          trace.Tracer
	accessKeyId     string
	secretAccessKey string
	defaultRegion   string
	bucket          string
	endPoint        string
	s3Client        *s3.S3
}

func New(http *requests.HttpClient, logger *logrus.Entry, accessKeyId, secretAccessKey, defaultRegion, bucket, endPoint string, s3Client *s3.S3) *Client {
	return &Client{
		http:            http,
		log:             logger,
		tracer:          otel.Tracer(packageName),
		accessKeyId:     accessKeyId,
		secretAccessKey: secretAccessKey,
		defaultRegion:   defaultRegion,
		bucket:          bucket,
		endPoint:        endPoint,
		s3Client:        s3Client,
	}
}
