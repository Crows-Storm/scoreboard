package config

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ReadMongoClient *mongo.Client
var WriteMongoClient *mongo.Client

var ReadMongoCtx = context.Background()
var WriteMongoCtx = context.Background()

var MongoConfig struct {
	host     string
	port     int
	password string

	database string
}
