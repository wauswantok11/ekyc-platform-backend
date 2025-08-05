package aws

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (ctx *Client) AwsUploadFile(key string, img image.Image, format string) error {
	var buf bytes.Buffer
	switch strings.ToLower(format) {
	case "jpeg":
		if err := jpeg.Encode(&buf, img, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(&buf, img); err != nil {
			return err
		}
	default:
		return errors.New("unsupported format")
	}
	input := &s3.PutObjectInput{
		Bucket: aws.String(ctx.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
	}

	if _, err := ctx.s3Client.PutObject(input); err != nil {
		return err
	}
	return nil
}

func (ctx *Client) AwsReadFile(key string) (string, error) {
	fmt.Printf("Final S3 key being used: %s\n", key) // Log the key for debugging

	input := &s3.GetObjectInput{
		Bucket: aws.String(ctx.bucket),
		Key:    aws.String(key),
	}

	result, err := ctx.s3Client.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				return "", fmt.Errorf("object not found at key: %s", key)
			default:
				return "", fmt.Errorf("error getting object from S3: %w", err)
			}
		}
		return "", fmt.Errorf("error getting object from S3: %w", err)
	}
	defer result.Body.Close()

	// Read the object data into a buffer
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, result.Body)
	if err != nil {
		return "", fmt.Errorf("error reading object data: %w", err)
	}

	// Convert the data to a Base64 encoded string
	encodedData := base64.StdEncoding.EncodeToString(buf.Bytes())

	return encodedData, nil
}
