package module

import (
	"net/http"

	"git.inet.co.th/ekyc-platform-backend/app"
	"git.inet.co.th/ekyc-platform-backend/middleware"
	frontweb "git.inet.co.th/ekyc-platform-backend/module/frontweb"
	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

func Create(app *app.Context) error {
	logrus := app.NewLogger().WithField("module", "generic")
	redis, err := app.NewCacheClient(logrus)
	if err != nil {
		logrus.Errorln("[x] Init global caching module error -:", err)
		return err
	}
	aclSkipper := middleware.NewSkipperPath("")
	aclSkipper.Add("/api/-/health", http.MethodGet)
	app.Router.Use(middleware.NewACLMiddleware(&aclSkipper, redis))

	trSkipper := util.NewHttpSkipper()
	trSkipper.Add(http.MethodGet, "/api/-/health")
	app.Router.Use(middleware.GoFiberTracer(trSkipper))

	if err := frontweb.Create(app); err != nil {
		logrus.Errorln("[x] Create frontweb module error -:", err)
		return err
	}

	return nil
}
