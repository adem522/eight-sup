package database

import (
	"context"
	"fmt"
	"time"

	"github.com/adem522/eight-sup/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateEvent(data *models.Event, event *mongo.Collection, user *mongo.Collection) error {
	data.Date = time.Now().Add(3 * time.Hour)
	if err := checkClient(user, data); err != nil {
		return fmt.Errorf("error from check client - %w", err)
	}
	if err := checkStreamer(user, data); err != nil {
		return fmt.Errorf("error from check streamer - %w", err)
	}
	if err := pushClient(user, data); err != nil {
		return fmt.Errorf("error from push client - %w", err)
	}
	if err := pushStreamer(user, data); err != nil {
		return fmt.Errorf("error from push streamer - %w", err)
	}
	_, err := event.InsertOne(
		context.TODO(), data,
	)
	if err != nil {
		return fmt.Errorf("error from create event - %w", err)
	}
	return nil
}

func CreatePlanInfo(data *models.PlanInfoStruct, planInfo *mongo.Collection) (interface{}, error) {
	result, err := planInfo.InsertOne(
		context.TODO(), data,
	)
	if err != nil {
		return nil, fmt.Errorf("error from create event and error code - %w", err)
	}
	return result.InsertedID, nil
}

//add plan when streamer take plan
func PushPlan(u *models.UserStructAddPlan, collection *mongo.Collection) error {
	deneme := models.PlanStruct{
		Package: models.PackageStruct{
			Stock:  u.Number,
			Date:   time.Now().Add(3 * time.Hour),
			Unique: u.Unique,
		},
		SellerUsername: "system",
	}
	if checkPlan(u, collection) {
		if pushPlanIfExist(u, collection) {
			return nil
		}
		return fmt.Errorf("error from increasing plan streamer and error code ")
	}
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"username": u.Username},
		bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "plan", Value: deneme},
			}},
		},
	)
	if err != nil {
		return fmt.Errorf("error from adding plan streamer and error code = %g", err)
	}
	return nil
}

func RegisterUser(data1 *models.UserStruct, collection *mongo.Collection) error {
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"username": data1.Username}).Decode(&result)
	if result != nil {
		return fmt.Errorf("already registered user %g", err)
	}
	_, err = collection.InsertOne(context.TODO(), data1)
	if err != nil {
		return fmt.Errorf("error from handlers/register %g", err)
	}
	return nil
}

func CreateAllPlan(collection *mongo.Collection) (interface{}, error) {
	collection.Drop(context.TODO()) //if exists
	items := []string{
		"Leaving an address in chat about channel",
		"Talk about the channel",
		"Invite to the channel",
		"Visit the channel",
		"Host the channel",
		"Play together",
	}
	data := []interface{}{
		models.PlanInfoStruct{
			Unique: "bronze",
			Name:   "Bronze Package",
			Desc:   "Bronze Desc",
			Color:  "#CD7F32",
			Cost:   2.99,
			Items:  items[:1],
		},
		models.PlanInfoStruct{
			Unique: "silver",
			Name:   "Silver Package",
			Desc:   "Silver Desc",
			Color:  "#C0C0C0",
			Cost:   3.99,
			Items:  items[:2],
		},
		models.PlanInfoStruct{
			Unique: "gold",
			Name:   "Gold Package",
			Desc:   "Gold Desc",
			Color:  "#E8B923",
			Cost:   4.99,
			Items:  items[:3],
		},
		models.PlanInfoStruct{
			Unique: "emerald",
			Name:   "Emerald Package",
			Desc:   "Emerald Desc",
			Color:  "#50C878",
			Cost:   5.99,
			Items:  items[:4],
		},
		models.PlanInfoStruct{
			Unique: "vibranium",
			Name:   "Vibranium Package",
			Desc:   "Vibranium Desc",
			Color:  "#A5A5AB",
			Cost:   6.99,
			Items:  items[:5],
		},
		models.PlanInfoStruct{
			Unique: "diamond",
			Name:   "Diamond Package",
			Desc:   "Diamond Desc",
			Color:  "#B9F2FF",
			Cost:   7.99,
			Items:  items[:6],
		},
	}
	return collection.InsertMany(context.TODO(), data)
}

func CreateWant(want *models.Want, col1, col2 *mongo.Collection) (interface{}, error) {
	/*if err := deleteProp(want, col2); err != nil {
		return nil, fmt.Errorf("error from deleteProp and err %w", err)
	}*/
	err := deletePackage(want, col2)
	if err != nil {
		fmt.Println("deletePackage error ", err)
	}
	return col1.InsertOne(context.TODO(), want)
}

func deleteProp(want *models.Want, col *mongo.Collection) error {
	return col.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"username": want.BuyerUsername,
		},
		bson.D{
			{Key: "$pull", Value: bson.D{
				{Key: "plan.$[elem].package.items", Value: want.Prop},
			}},
		},
		options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{
				{Key: "elem.package.unique", Value: want.Unique},
				{Key: "elem.sellerusername", Value: want.SellerUsername},
			}},
		}),
	).Err()
}

func deletePackage(want *models.Want, col *mongo.Collection) error {
	return col.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"username": want.BuyerUsername,
		},
		bson.D{
			{Key: "$pull", Value: bson.D{
				{Key: "plan.$[elem]", Value: want.Unique},
			}},
		},
		options.FindOneAndUpdate().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{
				{Key: "elem.package.unique", Value: want.Unique},
				{Key: "elem.sellerusername", Value: want.SellerUsername},
			}},
		}),
	).Err()
}
