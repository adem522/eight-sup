package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Want struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	SellerUsername string             `bson:"sellerUsername"`
	BuyerUsername  string             `bson:"buyerUsername"`
	Unique         string             `bson:"unique"`
	Prop           string             `bson:"prop"`
}
