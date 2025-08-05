package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Host       string
	Port       string
	Username   string
	Password   string
	DbName     string
	Connection *string
}

type Client struct {
	cfg Config
	db  *mongo.Client
}

func NewWithConfig(cfg Config) Client {
	return Client{
		cfg: cfg,
	}
}

func NewWithConnection(conn string) Client {
	return Client{
		cfg: Config{
			Host:       "",
			Port:       "",
			Username:   "",
			Password:   "",
			DbName:     "",
			Connection: &conn,
		},
	}
}
