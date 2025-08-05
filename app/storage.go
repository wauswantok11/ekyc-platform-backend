package app

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

func (ctx *Context) NewAwsClient(identifier string, endpoint string, region string, accessKey string, secretKey string, bucket string, debug bool, logger *logrus.Entry) (*s3.S3, error) {
	if region == "" {
		return nil, errors.New("region cannot be empty")
	}

	if accessKey == "" || secretKey == "" {
		return nil, errors.New("access key and secret key cannot be empty")
	}

	logger.Infoln("[*] Initialize aws", identifier)

	awsConfig := &aws.Config{
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		ctx.log.Errorf("Error creating AWS session: %v", err)
		return nil, err
	}

	/// Create an S3 service client
	s3Client := s3.New(sess)

	ctx.s3Client = s3Client

	if debug {
		ctx.log = logger
	}

	ctx.log.Infoln("[*] AWS S3 connection ")

	// Optional: Verify bucket existence (not strictly necessary, but useful for debugging)
	if _, err = s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}); err != nil {
		logger.Errorf("Error accessing bucket %s: %v", bucket, err)
		return nil, err
	}

	return s3Client, nil
}

func (ctx *Context) NewDBAwsClient(logger *logrus.Entry) (*s3.S3, error) {
	return ctx.NewAwsClient(
		"DBAws",
		ctx.Config.Aws.EndPoint,
		ctx.Config.Aws.DefaultRegion,
		ctx.Config.Aws.AccessKeyId,
		ctx.Config.Aws.SecretAccessKey,
		ctx.Config.Aws.Bucket,
		ctx.Config.App.IsDebug(),
		logger,
	)
}
