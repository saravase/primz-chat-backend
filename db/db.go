package db

import (
	"context"
	"log"

	"github.com/saravase/primz-chat-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewConnection(cfg config.Config) (Connection, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Dsn()))
	if err != nil {
		return nil, err
	}

	database := client.Database(cfg.DbName())

	return &conn{
		client:   client,
		database: database,
	}, nil
}

func (c *conn) Close() {
	err := c.client.Disconnect(context.TODO())
	if err != nil {
		log.Panic(err)
	}
}

func (c *conn) DB() *mongo.Database {
	return c.database
}
