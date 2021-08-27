package handlers

import "go.mongodb.org/mongo-driver/mongo"

type Collection struct {
	C1 *mongo.Collection
	C2 *mongo.Collection
	N  string
}
