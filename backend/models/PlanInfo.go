package models

type PlanInfoStruct struct {
	Unique string   `bson:"unique,omitempty default:defaultUnique"`
	Name   string   `bson:"name,omitempty default:default Name"`
	Desc   string   `bson:"desc,omitempty default:default Desc"`
	Color  string   `bson:"color"`
	Cost   float64  `bson:"cost,omitempty default:0"`
	Items  []string `bson:"items"`
}
