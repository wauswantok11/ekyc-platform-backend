package one_id

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"git.inet.co.th/ekyc-platform-backend/pkg/requests"
)

const packageName = "one_id"

func New(url, clientId, clientSecret, refCode string, timeout int, http *requests.HttpClient, logger *logrus.Entry) *Client {
	logger.Infoln("[*] Initialize one-id package")
	return &Client{
		url:          url,
		clientId:     clientId,
		clientSecret: clientSecret,
		refCode:      refCode,
		timeOut:      timeout,
		http:         http,
		log:          logger.Dup().WithField("package", packageName),
		tracer:       otel.Tracer(packageName),
	}
}
