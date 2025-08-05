package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

type Config struct {
	App    appConfig
	Secret secretCfg
	Server serverCfg
	Syslog syslogCfg
	Tracer tracerCfg
	DBMain mariadbCfg
	Mongo  mongodbCfg
	Redis  redisCfg
	Aws    awsCfg
	OneId  oneIdCfg
}

type appConfig struct {
	Name       string
	Version    string
	Mode       string
	PrefixPath string
	StorageDir string
	LogLevel   logrus.Level
}

type secretCfg struct {
	EncryptKey string
	JwtKey     string
}

func (c *appConfig) IsDebug() bool {
	return c.LogLevel == logrus.DebugLevel
}

func LoadConfig(file string, version string) *Config {
	// Read configuration
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalln("load config file error:", err.Error())
	}

	// Set default configuration
	// Main app
	viper.SetDefault("app.name", "HealthID")
	viper.SetDefault("app.version", version)
	viper.SetDefault("app.mode", EnvDevelopment)
	viper.SetDefault("app.storage", "./storage")
	viper.SetDefault("app.prefix.path", "/api")
	viper.SetDefault("app.log.level", "info")

	// Secret
	viper.SetDefault("secret.encrypt.key", "")
	viper.SetDefault("secret.jwt.key", "")
	// Server
	viper.SetDefault("server.listen", ServerListen)
	viper.SetDefault("server.port", ServerPort)
	viper.SetDefault("server.timeout.read", ServerTimeoutRead)
	viper.SetDefault("server.timeout.write", ServerTimeoutWrite)
	viper.SetDefault("server.timeout.idle", ServerTimeoutIdle)
	viper.SetDefault("server.header", viper.GetString("app.name"))
	viper.SetDefault("server.proxy.header", fiber.HeaderXForwardedFor)
	viper.SetDefault("server.enable.cors", "false")

	// Syslog
	viper.SetDefault("syslog.enable", "false")
	viper.SetDefault("syslog.server", "127.0.0.1")
	viper.SetDefault("syslog.port", "514")
	viper.SetDefault("syslog.protocol", "udp")

	// DBMain
	viper.SetDefault("db.main.host", MariadbHost)
	viper.SetDefault("db.main.port", MariadbPort)
	viper.SetDefault("db.main.username", "")
	viper.SetDefault("db.main.password", "")
	viper.SetDefault("db.main.database", "")

	// DBLog
	viper.SetDefault("db.log.host", MariadbHost)
	viper.SetDefault("db.log.port", MariadbPort)
	viper.SetDefault("db.log.username", "")
	viper.SetDefault("db.log.password", "")
	viper.SetDefault("db.log.database", "")

	// Mongo
	viper.SetDefault("mongo.connection", "")

	// Redis
	viper.SetDefault("redis.host", RedisHost)
	viper.SetDefault("redis.port", RedisPort)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.database", RedisDb)

	// One ID
	viper.SetDefault("oneid.ref_code", "")
	viper.SetDefault("oneid.client_id", "")
	viper.SetDefault("oneid.client_secret", "")
	viper.SetDefault("oneid.url", "")
	viper.SetDefault("oneid.timeout", ServerTimeoutIdle)
	// Set configuration
	config := &Config{
		App: appConfig{
			Name:       viper.GetString("app.name"),
			Version:    viper.GetString("app.version"),
			Mode:       viper.GetString("app.mode"),
			PrefixPath: viper.GetString("app.prefix.path"),
			StorageDir: viper.GetString("app.storage"),
		},
		Secret: secretCfg{
			EncryptKey: viper.GetString("secret.encrypt.key"),
			JwtKey:     viper.GetString("secret.jwt.key"),
		},
		Server: serverCfg{
			ListenIp:     viper.GetString("server.listen"),
			Port:         viper.GetString("server.port"),
			TimeoutRead:  util.ParseDuration(viper.GetString("server.timeout.read")),
			TimeoutWrite: util.ParseDuration(viper.GetString("server.timeout.write")),
			TimeoutIdle:  util.ParseDuration(viper.GetString("server.timeout.idle")),
			ServerHeader: viper.GetString("server.header"),
			ProxyHeader:  viper.GetString("server.proxy.header"),
			EnableCORS:   viper.GetBool("server.enable.cors"),
		},
		Syslog: syslogCfg{
			Enable:   viper.GetString("syslog.enable") == "true",
			Server:   viper.GetString("syslog.server"),
			Port:     viper.GetString("syslog.port"),
			Protocol: viper.GetString("syslog.protocol"),
		},
		Tracer: tracerCfg{
			Enable: viper.GetString("tracer.enable") == "true",
			Url:    viper.GetString("tracer.url"),
		},
		DBMain: mariadbCfg{
			Host:     viper.GetString("db.main.host"),
			Port:     viper.GetString("db.main.port"),
			User:     viper.GetString("db.main.username"),
			Password: viper.GetString("db.main.password"),
			Database: viper.GetString("db.main.database"),
		},
		Mongo: mongodbCfg{
			Connection: viper.GetString("mongo.connection"),
		},
		Redis: redisCfg{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			Database: viper.GetString("redis.database"),
		},
		Aws: awsCfg{
			AccessKeyId:     viper.GetString("aws.access_key_id"),
			SecretAccessKey: viper.GetString("aws.secret_access_key"),
			DefaultRegion:   viper.GetString("aws.default_region"),
			Bucket:          viper.GetString("aws.bucket"),
			EndPoint:        viper.GetString("aws.end_point"),
		},
		OneId: oneIdCfg{
			RefCode:      viper.GetString("one_id.ref_code"),
			ClientId:     viper.GetString("one_id.client_id"),
			ClientSecret: viper.GetString("one_id.client_secret"),
			Timeout:      viper.GetInt("one_id.timeout"),
			Url:          viper.GetString("one_id.url"),
		},
	}

	config.App.LogLevel, err = logrus.ParseLevel(viper.GetString("app.log.level"))
	if err != nil {
		config.App.LogLevel = logrus.InfoLevel
	}
	return config
}
