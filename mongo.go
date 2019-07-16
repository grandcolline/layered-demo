package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoConnect はmongoDBに接続します
func mongoConnect() *mongo.Database {
	var err error

	// mongoDBの設定の読み込み
	var conf mongoConf
	conf.init()

	// DB接続
	url := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
	)
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	db := client.Database(conf.Database)

	return db
}
