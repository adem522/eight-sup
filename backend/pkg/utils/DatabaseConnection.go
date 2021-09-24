package utils

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Close(client *mongo.Client) {
	// client provides a method to close
	// a mongoDB connection.
	defer func() {
		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func Connect() *mongo.Client {
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb+srv://"+os.Getenv("DBUSERNAME")+":"+os.Getenv("DBPASSWORD")+"@cluster0.4lioy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"),
	)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

/*
func ping(client *mongo.Client, ctx context.Context) error {
	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occored, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}
*/
