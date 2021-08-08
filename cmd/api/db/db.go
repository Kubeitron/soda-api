package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Mongodb struct {
		BaseCtx context.Context
		Conn    *mongo.Client
	}
)

func NewMongodb() *Mongodb {
	dbUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	dbPass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	dbAddress := os.Getenv("MONGO_CLUSTER_ADDRESS")
	dbName := os.Getenv("MONGO_INITDB_DATABASE")
	if dbUser == "" || dbPass == "" || dbAddress == "" || dbName == "" {
		log.Fatalf("db not configured properly: %s %s %s %s", dbUser, dbPass, dbAddress, dbName)
	}

	// uri := fmt.Sprintf("mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/%s?authSource=admin", dbUser, dbPass, dbAddress, dbName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Testing mongodb connection...")
	ping_err := client.Ping(ctx, nil)
	if ping_err != nil {
		log.Fatalln(ping_err)
	}
	log.Print("Ping to db successful")
	db := &Mongodb{
		ctx,
		client,
	}
	return db
}
