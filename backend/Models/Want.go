package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Want struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Date           time.Time          `bson:"date,omitempty"`
	SellerUsername string             `bson:"sellerUsername"`
	BuyerUsername  string             `bson:"buyerUsername"`
	Unique         string             `bson:"unique"`
	Prop           string             `bson:"prop"`
	Status         string             `bson:"status"`
}
