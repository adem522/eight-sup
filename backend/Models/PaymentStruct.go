package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentStruct struct {
	PaymentType PaymentTypeStruct
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Value       float64            `bson:"value,omitempty"`
}
