package config

type mariadbCfg struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type mongodbCfg struct {
	Connection string
}

type redisCfg struct {
	Host     string
	Port     string
	Password string
	Database string
}
