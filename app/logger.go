package app

import (
	"fmt"
	"log/syslog"

	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"

	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

func (ctx *Context) NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(ctx.Config.App.LogLevel)
	return l
}

func (ctx *Context) AddSyslogHook(l *logrus.Entry, tag string) *logrus.Entry {
	e := l.Dup()
	if ctx.Config.Syslog.Enable {
		if !util.Contains([]string{"udp", "tcp"}, ctx.Config.Syslog.Protocol) {
			e.Errorln("Syslog protocol error -: should be udp or tcp but got", ctx.Config.Syslog.Protocol)
			return e
		}
		if hook, err := lSyslog.NewSyslogHook(
			"udp",
			fmt.Sprintf("%s:%s", ctx.Config.Syslog.Server, ctx.Config.Syslog.Port),
			syslog.LOG_INFO,
			tag,
		); err == nil {
			e.Logger.AddHook(hook)
		}
	}
	return e
}
