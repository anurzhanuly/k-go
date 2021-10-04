package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

// Options настройки монго
type Options struct {
	URL                      string
	Database                 string
	ConnectTimeoutMs         int
	WriteTimeoutMs           int
	SocketTimeoutMs          int
	ServerSelectionTimeoutMs int
}

// Dial подключение к указанной коллекции монго
func Dial(opt Options) (collection *mongo.Database, err error) {
	client, err := mongo.NewClient(
		options.Client().
			ApplyURI(opt.URL).
			SetSocketTimeout(time.Millisecond * time.Duration(opt.SocketTimeoutMs)).
			SetServerSelectionTimeout(time.Millisecond * time.Duration(opt.ServerSelectionTimeoutMs)).
			SetConnectTimeout(time.Millisecond * time.Duration(opt.ConnectTimeoutMs)).
			SetWriteConcern(writeconcern.New(
				writeconcern.J(true),
				writeconcern.WTimeout(time.Duration(opt.WriteTimeoutMs)*time.Millisecond),
			)),
	)

	if err != nil {
		return
	}

	if err = client.Connect(context.Background()); err != nil {
		return
	}

	if err = client.Ping(context.Background(), readpref.PrimaryPreferred()); err != nil {
		return
	}

	return client.Database(opt.Database), nil
}
