package database

import (
	"context"
	"fmt"

	"github.com/adem522/eight-sup/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//update plan when client take plan
func UpdatePlan(u *Models.UserStructAddPlan, collection *mongo.Collection) error {
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"username": u.Username, "plan.package.unique": u.Unique},
		bson.D{
			{Key: "$inc", Value: bson.D{
				{Key: "plan.$.package.stock", Value: u.Number},
			}},
		},
	)
	if err != nil {
		return fmt.Errorf("error from adding plan streamer and error code = %g", err)
	}
	return nil
}
