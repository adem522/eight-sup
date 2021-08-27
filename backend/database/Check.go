package database

import (
	"context"
	"fmt"
	"time"

	"github.com/adem522/eight-sup/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkClient(collection *mongo.Collection, event *Models.Event) error {
	var filter2 Models.PlanStruct
	var filter Models.UserStruct
	opt := options.FindOne().SetProjection(bson.M{"plan": 1})
	collection.FindOne(
		context.TODO(),
		bson.M{
			"username": event.BuyerUsername,
		}, opt,
	).Decode(&filter)
	if filter.Plan != nil {
		for _, filter2 = range filter.Plan {
			if filter2.Package.Unique == event.Unique && filter2.SellerUsername == event.SellerUsername {
				return fmt.Errorf("error from updating buyer stock and error code - already have")
			}
		}
	}
	return nil
}

func checkStreamer(collection *mongo.Collection, event *Models.Event) error {
	var filter2 Models.PlanStruct
	var filter Models.UserStruct
	collection.FindOne(
		context.TODO(),
		bson.M{
			"username": event.SellerUsername,
			"type":     "streamer",
		},
	).Decode(&filter)
	if filter.Plan != nil {
		for _, filter2 = range filter.Plan {
			if filter2.Package.Unique == event.Unique && filter2.Package.Stock > 0 {
				return nil
			}
		}
	}
	return fmt.Errorf("error from updating seller stock and error code -seller don't have stock")
}

func pushStreamer(collectionUser *mongo.Collection, event *Models.Event) error {
	_, err := collectionUser.UpdateOne(
		context.TODO(),
		bson.M{
			"username":            event.SellerUsername,
			"plan.package.unique": event.Unique,
		},
		bson.D{
			{Key: "$inc", Value: bson.D{
				{Key: "plan.$.package.stock", Value: -1},
			}},
		},
	)
	if err != nil {
		return fmt.Errorf("error from updating seller stock and error code - %g", err)
	}
	return nil
}

func pushClient(collection *mongo.Collection, event *Models.Event) error {
	plan := Models.PlanStruct{
		Package: Models.PackageStruct{
			Date:   time.Now().Add(time.Hour * 3),
			Unique: event.Unique,
			Stock:  1,
			Items:  event.Items,
		},
		SellerUsername: event.SellerUsername,
	}
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{
			"username": event.BuyerUsername,
		},
		bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "plan", Value: plan},
			}},
		},
	)
	if err != nil {
		return fmt.Errorf("error from updating buyer stock and error code - %g", err)
	}
	return nil
}

func LoginCheck(data1, data2 string, collection *mongo.Collection) string {
	var result bson.M
	opt := options.FindOne().SetProjection(bson.M{"_id": 0, "plan": 0})
	collection.FindOne(context.TODO(), bson.M{
		"username": data1,
		"password": data2,
	},
		opt).Decode(&result)
	if result == nil {
		return ""
	}
	return result["type"].(string)
}

func checkPlan(u *Models.UserStructAddPlan, collection *mongo.Collection) bool {
	var filter Models.UserStruct
	opt := options.FindOne().SetProjection(bson.M{"plan": 1})
	collection.FindOne(
		context.TODO(),
		bson.M{
			"username":            u.Username,
			"plan.package.unique": u.Unique,
		}, opt,
	).Decode(&filter)
	return filter.Plan != nil
}

func pushPlanIfExist(u *Models.UserStructAddPlan, collection *mongo.Collection) bool {
	doc := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"username": u.Username,
		},
		bson.D{
			{Key: "$inc", Value: bson.D{
				{Key: "plan.$[elem].package.stock", Value: u.Number},
			}},
		},
		options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{
				{Key: "elem.package.unique", Value: u.Unique},
			}},
		}).SetReturnDocument(1),
	)
	return doc != nil
}
