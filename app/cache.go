package app

import (
	"github.com/sirupsen/logrus"

	"git.inet.co.th/ekyc-platform-backend/pkg/cache"
	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

func (ctx *Context) NewCacheClient(logger *logrus.Entry) (*cache.Redis, error) {
	logger.Infoln("[*] Initialize caching")
	c := cache.NewWithCfg(cache.Config{
		Host:     ctx.Config.Redis.Host,
		Port:     ctx.Config.Redis.Port,
		Password: ctx.Config.Redis.Password,
		Db:       util.AtoI(ctx.Config.Redis.Database, 0),
	})
	if err := c.Ping(); err != nil {
		logger.Errorln("[x] caching connection error:", err.Error())
		return nil, err
	}
	return &c, nil
}
