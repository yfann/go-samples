package sample

import (
	"context"
	"fmt"
	"github.com/google/martian/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


func TestMongo(){


	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://xxx:xxx@xxxxxx:3717/hsc-service-broker"))
	if err != nil {
		log.Errorf("",err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Errorf("",err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Errorf("",err)
	}
	fmt.Println(databases)
}