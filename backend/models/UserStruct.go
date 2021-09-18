package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStruct struct {
	Plan           []PlanStruct
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `bson:"username,omitempty" json:"username"`
	Password       string             `bson:"password,omitempty" json:"password"`
	Type           string             `bson:"type,omitempty"`
	Email          string             `bson:"email,omitempty" json:"email"`
	Name           string             `bson:"name,omitempty"`
	TwitchUsername string             `bson:"twitchUsername,omitempty"`
}

type UserStructAddPlan struct {
	Username       string `bson:"username,omitempty" json:"username"`
	Unique         string `bson:"unique,omitempty"`
	SellerUsername string `bson:"sellerusername,omitempty"`
	Type           string `bson:"type,omitempty"`
	Number         int    `bson:"number,omitempty"`
}
