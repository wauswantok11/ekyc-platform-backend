package app

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"git.inet.co.th/ekyc-platform-backend/config"
)

type Context struct {
	Config   *config.Config
	Router   *fiber.App
	log      *logrus.Entry
	Tracer   *sdktrace.TracerProvider
	s3Client *s3.S3
}

type App struct {
	*Context
}

func New(cfg *config.Config) *App {
	l := logrus.New()
	l.SetLevel(cfg.App.LogLevel)
	return &App{Context: &Context{
		Config: cfg,
		log:    l.WithField("package", "app"),
	}}
}
