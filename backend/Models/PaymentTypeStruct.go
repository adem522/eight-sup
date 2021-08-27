package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentTypeStruct struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	TypeName string             `bson:"typeName,omitempty"`
}
