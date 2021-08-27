package Models

import "time"

type PackageStruct struct {
	Date   time.Time `bson:"date,omitempty"`
	Stock  int       `bson:"stock,omitempty default:0"`
	Unique string    `bson:"unique"`
	Items  []string  `bson:"items,omitempty"`
}

type PlanStruct struct {
	Package        PackageStruct
	SellerUsername string `bson:"sellerusername" `
}
