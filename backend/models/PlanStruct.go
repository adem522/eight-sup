package models

import "time"

type Items struct {
	BuyerUsername string `bson:"buyerUsername"`
	Status        string `bson:"status"`
	Prop          string `bson:"prop"`
}

type PackageStruct struct {
	Date   time.Time `bson:"date,omitempty"`
	Stock  int       `bson:"stock,omitempty default:0"`
	Unique string    `bson:"unique"`
	Items  []Items   `bson:"items,omitempty"`
}

type PlanStruct struct {
	Package        PackageStruct
	SellerUsername string `bson:"sellerusername" `
}
