package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Vegetable struct {
	Color    string  `json:"color" xml:"color"`
	Weight   float32 `json:"weight" xml:"weight"`
	Name     string  `json:"name" xml:"name"`
	Vitamins string  `json:"vitamins" xml:"vitamins"`
	Calories int     `json:"calories" xml:"calories"`
}

func main() {
	dbUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	dbPass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	dbAddress := os.Getenv("MONGO_CLUSTER_ADDRESS")
	dbName := os.Getenv("MONGO_INITDB_DATABASE")

	if dbUser == "" || dbPass == "" || dbAddress == "" || dbName == "" {
		log.Fatalf("db not configured properly: %s %s %s %s", dbUser, dbPass, dbAddress, dbName)
	}

	// uri := fmt.Sprintf("mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/%s?authSource=admin", dbUser, dbPass, dbAddress, dbName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = db.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	log.Print("Testing mongodb connection...")
	ping_err := db.Ping(ctx, nil)
	if ping_err != nil {
		log.Fatalln(ping_err)
	}

	log.Print("Writing data")
	vegetableCollection := db.Database(dbName).Collection("vegetables")
	veg := Vegetable{
		Color:    "green",
		Weight:   0.5,
		Name:     "cucumber",
		Vitamins: "None",
		Calories: 80,
	}
	vegetableCollection.InsertOne(context.TODO(), veg)

	log.Print("Reading data")
	var results []*Vegetable
	findOptions := options.Find()
	findOptions.SetLimit(5)
	cur, err := vegetableCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Vegetable
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	fmt.Printf("Doc0%+v\n", results[0])

	log.Print("Updating data")
	log.Print("Deleting data")
	log.Print("Reading data")
}
