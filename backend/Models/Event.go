package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Date           time.Time          `bson:"date,omitempty"`
	BuyerUsername  string             `bson:"buyerUsername,omitempty"`
	SellerUsername string             `bson:"sellerUsername,omitempty"`
	Unique         string             `bson:"unique,omitempty"`
	Items          []Items            `bson:"items,omitempty"`
}
