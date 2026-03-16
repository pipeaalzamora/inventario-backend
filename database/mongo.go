package database

import (
	"context"
	"sofia-backend/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(s *config.Config) *mongo.Database {

	ctx := context.TODO()

	opt := options.Client().ApplyURI(s.Mongo.Host)

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}
	db := client.Database(s.Mongo.DBName)

	return db

}
