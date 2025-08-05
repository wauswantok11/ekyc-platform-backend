package app

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/gorm"

	"git.inet.co.th/ekyc-platform-backend/pkg/database"
	gorm_logrus "git.inet.co.th/ekyc-platform-backend/pkg/gorm-logrus"
	"git.inet.co.th/ekyc-platform-backend/pkg/mongodb"
)

func (ctx *Context) NewMongoClient(logger *logrus.Entry) (*mongodb.Client, error) {
	logger.Infoln("[*] Initialize mongodb datastore")
	cfg := mongodb.Config{
		Connection: &ctx.Config.Mongo.Connection,
	}
	c := mongodb.NewWithConfig(cfg)
	if err := c.Connect(); err != nil {
		logger.Errorln("[*] mongo connection error:", err.Error())
		return nil, err
	}
	return &c, nil
}

func (ctx *Context) NewMariaDBClient(identifier string, h string, p string, user string, pass string, name string, debug bool, logger *logrus.Entry) (*database.Client, error) {
	logger.Infoln("[*] Initialize database", identifier)
	db := database.NewWithConfig(
		database.Config{
			Host:      h,
			Port:      p,
			Username:  user,
			Password:  pass,
			Name:      name,
			DebugMode: debug,
		},
		ctx.NewLogger(),
	)
	logger.Infoln("[*] Connecting to database", identifier, "...")
	if err := db.ConnectWithGormConfig(gorm.Config{
		Logger: gorm_logrus.New(identifier, logger, time.Second, ctx.Config.App.LogLevel == logrus.DebugLevel),
	}); err != nil {
		logger.Errorln("[x] database", identifier, "connection error:", err.Error())
		return nil, err
	}
	// logger.Infoln("[*] Connected to database", identifier, ctx.Config.DBMain.Host, "name", ctx.Config.DBMain.Database)
	if ctx.Tracer != nil {
		_ = db.Ctx().Use(otelgorm.NewPlugin())
	}
	return &db, nil
}

func (ctx *Context) NewDBMainClient(logger *logrus.Entry) (*database.Client, error) {
	return ctx.NewMariaDBClient(
		"DBMain",
		ctx.Config.DBMain.Host,
		ctx.Config.DBMain.Port,
		ctx.Config.DBMain.User,
		ctx.Config.DBMain.Password,
		ctx.Config.DBMain.Database,
		ctx.Config.App.IsDebug(),
		logger,
	)
}

func (ctx *Context) NewStorageS3Client(logger *logrus.Entry) {
	logger.Infoln("[*] Initialize cloud storage datastore")

}
