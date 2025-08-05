package repositories

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"git.inet.co.th/ekyc-platform-backend/app"
	"git.inet.co.th/ekyc-platform-backend/config"
	"git.inet.co.th/ekyc-platform-backend/pkg/aws"
	"git.inet.co.th/ekyc-platform-backend/pkg/cache"
	"git.inet.co.th/ekyc-platform-backend/pkg/database"
	"git.inet.co.th/ekyc-platform-backend/pkg/mongodb"
	oneid "git.inet.co.th/ekyc-platform-backend/pkg/one-id"
	"git.inet.co.th/ekyc-platform-backend/pkg/requests"
)

const moduleName = "frontweb"

type Repository struct {
	app    *app.Context
	http   *requests.HttpClient
	log    *logrus.Entry
	tracer trace.Tracer
	dbMain *database.Client
	mongo  *mongodb.Client
	cache  *cache.Redis
	aws    *aws.Client
	oneId  *oneid.Client
}

func (r Repository) Module() string {
	return moduleName
}

func (r Repository) AppCfg() *config.Config {
	return r.app.Config
}

func (r Repository) Log() *logrus.Entry {
	return r.log.Dup()
}

func (r Repository) Trace(ctx context.Context, spanName string, attributes ...trace.SpanStartOption) (context.Context, trace.Span) {
	return r.tracer.Start(ctx, spanName, attributes...)
}

func (r Repository) Aws() *aws.Client {
	return r.aws
}

func (r Repository) Cache() *cache.Redis {
	return r.cache
}

func (r Repository) OneId() *oneid.Client {
	return r.oneId
}

func New(app *app.Context) (*Repository, error) {
	logrus := app.NewLogger().WithField("module", moduleName)
	dbMain, err := app.NewDBMainClient(logrus)
	if err != nil {
		return nil, err
	}
	mongo, err := app.NewMongoClient(logrus)
	if err != nil {
		return nil, err
	}
	redis, err := app.NewCacheClient(logrus)
	if err != nil {
		return nil, err
	}
	dbAws, err := app.NewDBAwsClient(logrus)
	if err != nil {
		return nil, err
	}
	httpClient := requests.NewHttpClient(app.AddSyslogHook(logrus, moduleName))

	return &Repository{
		app:    app,
		http:   httpClient,
		log:    logrus,
		tracer: otel.Tracer(moduleName),
		dbMain: dbMain,
		aws:    aws.New(httpClient, logrus, app.Config.Aws.AccessKeyId, app.Config.Aws.SecretAccessKey, app.Config.Aws.DefaultRegion, app.Config.Aws.Bucket, app.Config.Aws.EndPoint, dbAws),
		mongo:  mongo,
		cache:  redis,
		oneId:  oneid.New(app.Config.OneId.Url, app.Config.OneId.ClientId, app.Config.OneId.ClientSecret, app.Config.OneId.RefCode, app.Config.OneId.Timeout, httpClient, logrus),
	}, nil
}
