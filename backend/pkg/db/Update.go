package db

import (
	"context"
	"errors"

	"github.com/adem522/eight-sup/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//update plan when client take plan
func UpdatePlan(u *models.UserStructAddPlan, collection *mongo.Collection) error {
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
		return errors.New("error from adding plan streamer and error code = " + err.Error())
	}
	return nil
}
