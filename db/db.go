package db

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database keep database
type Database struct {
	*mongo.Database
	*mongo.Client
}

//New creates an instance of mongodb
func New() (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("DatabaseURI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("quickstart")
	return &Database{db, client}, nil
}

//Disconnect from database
func (db *Database) Disconnect() error {
	return db.Client.Disconnect(context.Background())
}
