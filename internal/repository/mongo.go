package repository

import (
	"context"

	"github.com/salesforceanton/products-data-manager/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(cfg *config.Config) (*mongo.Database, error) {
	// Set params and URI
	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.Uri)

	// Connect to DB and check connection with ping func
	dbClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	db := dbClient.Database(cfg.DB.DBName)
	return db, nil
}
