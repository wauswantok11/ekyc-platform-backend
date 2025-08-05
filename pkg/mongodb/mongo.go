package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Client) Connect() error {
	var dsn string
	if c.cfg.Connection != nil && *c.cfg.Connection != "" {
		// Use conn to connect
		dsn = *c.cfg.Connection
	} else {
		dsn = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", c.cfg.Username, c.cfg.Password, c.cfg.Host, c.cfg.Port, c.cfg.DbName)
	}
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(dsn),
		options.Client().SetConnectTimeout(time.Second*5),
		options.Client().SetTimeout(time.Second*60), // Every query should not over 1min
	)
	if err != nil {
		return err
	}
	c.db = client
	return c.Ping()
}

func (c *Client) Ping() error {
	return c.db.Ping(context.TODO(), nil)
}

func (c *Client) Ctx() *mongo.Client {
	return c.db
}

func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Disconnect(context.TODO())
	}
	return nil
}
