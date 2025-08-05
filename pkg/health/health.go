package health

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"git.inet.co.th/ekyc-platform-backend/pkg/requests"
)

const packageName = "health"

type Client struct {
	url       string
	http      *requests.HttpClient
	log       *logrus.Entry
	tracer    trace.Tracer
	clientId  string
	secretKey string
}

func New(http *requests.HttpClient, log *logrus.Entry, url, clientId, secretKey string) *Client {
	return &Client{
		url:       url,
		http:      http,
		log:       log,
		tracer:    otel.Tracer(packageName),
		clientId:  clientId,
		secretKey: secretKey,
	}
}
