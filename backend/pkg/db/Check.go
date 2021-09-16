package db

import (
	"context"
	"errors"

	"github.com/adem522/eight-sup/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkClient(collectionUser *mongo.Collection, event *models.Event, chan1 chan error) {
	var filter2 models.PlanStruct
	var filter models.UserStruct
	opt := options.FindOne().SetProjection(bson.M{"plan": 1})
	collectionUser.FindOne(
		context.TODO(),
		bson.M{
			"username": event.BuyerUsername,
		}, opt,
	).Decode(&filter)
	if filter.Plan != nil {
		for _, filter2 = range filter.Plan {
			if filter2.Package.Unique == event.Unique && filter2.SellerUsername == event.SellerUsername {
				chan1 <- errors.New("error from checking buyer stock and error code - already have")
				return
			}
		}
	}
	chan1 <- nil
}

func checkStreamer(collectionUser *mongo.Collection, event *models.Event, chan1 chan error) {
	var filter2 models.PlanStruct
	var filter models.UserStruct
	collectionUser.FindOne(
		context.TODO(),
		bson.M{
			"username": event.SellerUsername,
			"type":     "streamer",
		},
		options.FindOne().SetProjection(bson.M{"plan.package.items": 0}),
	).Decode(&filter)
	if filter.Plan != nil {
		for _, filter2 = range filter.Plan {
			if filter2.Package.Unique == event.Unique && filter2.Package.Stock > 0 {
				chan1 <- nil
				return
			}
		}
	}
	chan1 <- errors.New("error from checking seller stock and error code -seller don't have stock")
}

func pushStreamer(collectionUser *mongo.Collection, event *models.Event) error {
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
		return errors.New("error from updating seller stock and error code - " + err.Error())
	}
	return nil
}

func pushClient(collection *mongo.Collection, event *models.Event) error {
	plan := models.PlanStruct{
		Package: models.PackageStruct{
			Date:   event.Date,
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
		return errors.New("error from updating buyer stock and error code - " + err.Error())
	}
	return nil
}

func LoginCheck(data1, data2 string, col *mongo.Collection) string {
	var result bson.M
	opt := options.FindOne().SetProjection(bson.M{
		"_id":  0,
		"plan": 0,
	})
	col.FindOne(context.TODO(), bson.M{
		"username": data1,
		"password": data2,
	}, opt).Decode(&result)
	if result == nil {
		return ""
	}
	return result["type"].(string)
}

func checkPlan(u *models.UserStructAddPlan, collection *mongo.Collection) bool {
	var filter models.UserStruct
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

func pushPlanIfExist(u *models.UserStructAddPlan, collection *mongo.Collection) bool {
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
