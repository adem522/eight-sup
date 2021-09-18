package db

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/adem522/eight-sup/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateEvent(data *models.Event, event, user *mongo.Collection) error {
	chan1 := make(chan error)
	chan2 := make(chan error)
	b1, b2 := false, false
	go checkClient(user, data, chan1)
	go checkStreamer(user, data, chan2)
	for {
		select {
		case receiver := <-chan1:
			{
				if receiver != nil {
					return errors.New("error from create event - " + receiver.Error())
				} else {
					b1 = true
					break
				}
			}
		case receiver := <-chan2:
			{
				if receiver != nil {
					return errors.New("error from create event - " + receiver.Error())
				} else {
					b2 = true
					break
				}
			}
		}
		if b1 && b2 {
			var wg2 sync.WaitGroup
			data.Date = time.Now().Add(3 * time.Hour)
			wg2.Add(1)
			go func() error {
				err := pushClient(user, data)
				if err != nil {
					wg2.Done()
					return err
				}
				wg2.Done()
				return nil
			}()
			wg2.Add(1)
			go func() error {
				err := pushStreamer(user, data)
				if err != nil {
					wg2.Done()
					return err
				}
				wg2.Done()
				return nil
			}()
			wg2.Wait()
			_, err := event.InsertOne(
				context.TODO(), data,
			)
			if err != nil {
				return errors.New("error from create event - " + err.Error())
			}
			return nil
		}
	}
}

func CreatePlanInfo(data *models.PlanInfoStruct, planInfo *mongo.Collection) (interface{}, error) {
	result, err := planInfo.InsertOne(
		context.TODO(), data,
	)
	if err != nil {
		return nil, errors.New("error from create event and error code - " + err.Error())
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
		return errors.New("error from increasing plan streamer and error code ")
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
		return errors.New("error from adding plan streamer and error code = " + err.Error())
	}
	return nil
}

func RegisterUser(data1 *models.UserStruct, collection *mongo.Collection) error {
	var result bson.M
	err := collection.FindOne(
		context.TODO(),
		bson.M{"username": data1.Username},
	).Decode(&result)
	if result != nil {
		return errors.New("already registered user " + err.Error())
	}
	_, err = collection.InsertOne(context.TODO(), data1)
	if err != nil {
		return errors.New("error from handlers/register " + err.Error())
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

func DropIfExist(collection *mongo.Collection) error {
	return collection.Drop(context.TODO())
}

//Client
func ControllerWantClient(want *models.Want, col1, col2 *mongo.Collection) error {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() error {
		if err := createWant(want, col1); err != nil {
			wg.Done()
			wg.Done()
			wg.Done()
			return errors.New("error from createWant and err " + err.Error())
		} else {
			wg.Done()
			return nil
		}
	}()
	if want.Status == "want" {
		want.Status = "requested"
		go func() error {
			if err := updateClientProp(want, col2); err != nil {
				wg.Done()
				wg.Done()
				wg.Done()
				return errors.New("error from updatePropClient and err " + err.Error())
			} else {
				wg.Done()
				return nil
			}
		}()
		go func() error {
			if err := insertStreamerProp(want, col2); err != nil {
				wg.Done()
				wg.Done()
				wg.Done()
				return errors.New("error from insertStreamerProp and err " + err.Error())
			} else {
				wg.Done()
				return nil
			}
		}()
		wg.Wait()
	} else if want.Status == "getBack" {
		want.Status = "available"
		go func() error {
			if err := getBackPropClient(want, col2); err != nil {
				wg.Done()
				wg.Done()
				wg.Done()
				return errors.New("error from getBackPropClient and err " + err.Error())
			} else {
				wg.Done()
				return nil
			}
		}()
		go func() error {
			if err := getBackPropStreamer(want, col2); err != nil {
				wg.Done()
				wg.Done()
				wg.Done()
				return errors.New("error from getBackPropStreamer and err " + err.Error())
			} else {
				wg.Done()
				return nil
			}
		}()
		wg.Wait()
	} else if want.Status == "delete" {
		go func() error {
			if err := deletePropClient(want, col2); err != nil {
				wg.Done()
				wg.Done()
				wg.Done()
				return errors.New("error from deletePropClient and err " + err.Error())
			} else {
				wg.Done()
				return nil
			}
		}()
		go func() error {
			if err := deletePropStreamer(want, col2); err != nil {
				wg.Done()
				wg.Done()
				wg.Done()
				return errors.New("error from deletePropStreamer and err " + err.Error())
			} else {
				wg.Done()
				return nil
			}
		}()
		wg.Wait()
	}
	return nil
}

func createWant(want *models.Want, col *mongo.Collection) error {
	want.Date = time.Now().Add(time.Hour * 3)
	_, err := col.InsertOne(context.TODO(), want)
	if err != nil {
		return errors.New("error from database/insertPropClient " + err.Error())
	}
	return nil
}

func updateClientProp(want *models.Want, col *mongo.Collection) error {
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.BuyerUsername,
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "plan.$[elem].package.items.$[elem2].status", Value: want.Status},
			}},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{
					{Key: "elem.package.unique", Value: want.Unique},
					{Key: "elem.sellerusername", Value: want.SellerUsername},
				},
				bson.D{
					{Key: "elem2.prop", Value: want.Prop},
					{Key: "elem2.status", Value: "available"},
				},
			},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("this prop requested")
	}
	return nil
}

func insertStreamerProp(want *models.Want, col *mongo.Collection) error {
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.SellerUsername,
		},
		bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "plan.$[elem].package.items", Value: bson.D{
					{Key: "prop", Value: want.Prop},
					{Key: "buyerUsername", Value: want.BuyerUsername},
					{Key: "status", Value: "requested"},
				}},
			}},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{bson.D{
				{Key: "elem.package.unique", Value: want.Unique},
			}},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("this prop available")
	}
	return nil
}

func getBackPropClient(want *models.Want, col *mongo.Collection) error {
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.BuyerUsername,
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "plan.$[elem].package.items.$[elem2].status", Value: want.Status},
			}},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{
					{Key: "elem.package.unique", Value: want.Unique},
					{Key: "elem.sellerusername", Value: want.SellerUsername},
				},
				bson.D{
					{Key: "elem2.prop", Value: want.Prop},
					{Key: "elem2.status", Value: "requested"},
				},
			},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("can't delete")
	}
	return nil
}

func getBackPropStreamer(want *models.Want, col *mongo.Collection) error {
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.SellerUsername,
		},
		bson.M{"$pull": bson.M{"plan.$[elem].package.items": bson.M{
			"status":        "requested",
			"prop":          want.Prop,
			"buyerUsername": want.BuyerUsername,
		}}},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{
					{Key: "elem.package.unique", Value: want.Unique},
				},
			},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("can't delete")
	}
	return nil
}

func deletePropClient(want *models.Want, col *mongo.Collection) error {
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.BuyerUsername,
		},
		bson.M{"$pull": bson.M{"plan.$[elem].package.items": bson.M{
			"status": "completed",
			"prop":   want.Prop,
		}}},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{
					{Key: "elem.package.unique", Value: want.Unique},
					{Key: "elem.sellerusername", Value: want.SellerUsername},
				},
			},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("can't delete")
	}
	return nil
}

func deletePropStreamer(want *models.Want, col *mongo.Collection) error {
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.SellerUsername,
		},
		bson.M{"$pull": bson.M{"plan.$[elem].package.items": bson.M{
			"status":        "completed",
			"prop":          want.Prop,
			"buyerUsername": want.BuyerUsername,
		}}},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{
					{Key: "elem.package.unique", Value: want.Unique},
				},
			},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("can't delete")
	}
	return nil
}

////Streamer
func ControllerWantStreamer(want *models.Want, col1, col2 *mongo.Collection) error {
	if want.Status == "complete" {
		want.Status = "completed"
	} else if want.Status == "getBack" {
		want.Status = "requested"
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() error {
		if err := createWant(want, col1); err != nil {
			wg.Done()
			wg.Done()
			wg.Done()
			return errors.New("error from createWant and err " + err.Error())
		} else {
			wg.Done()
			return nil
		}
	}()
	go func() error {
		if err := updatePropForStreamer(want, col2); err != nil { //available to requested
			wg.Done()
			wg.Done()
			wg.Done()
			return errors.New("error from updatePropForStreamer and err " + err.Error())
		} else {
			wg.Done()
			return nil
		}
	}()
	go func() error {
		if err := updatePropClient2(want, col2); err != nil { //client hangi probu aldıysa
			wg.Done()
			wg.Done()
			wg.Done()
			return errors.New("error from updatePropClient2 and err " + err.Error())
		} else {
			wg.Done()
			return nil
		}
	}()
	wg.Wait()
	return nil
}

func updatePropForStreamer(want *models.Want, col *mongo.Collection) error {
	filter := ""
	if want.Status == "completed" {
		filter = "requested"
	} else if want.Status == "requested" {
		filter = "completed"
	}
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.SellerUsername,
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "plan.$[elem].package.items.$[elem2].status", Value: want.Status},
			}},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{
					{Key: "elem.package.unique", Value: want.Unique},
				},
				bson.D{
					{Key: "elem2.prop", Value: want.Prop},
					{Key: "elem2.status", Value: filter},
					{Key: "elem2.buyerUsername", Value: want.BuyerUsername},
				},
			},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("this prop not requested")
	}
	return nil
}
func updatePropClient2(want *models.Want, col *mongo.Collection) error {
	filter := ""
	if want.Status == "completed" {
		filter = "requested"
	} else if want.Status == "requested" {
		filter = "completed"
	}
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{
			"username": want.BuyerUsername,
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "plan.$[elem].package.items.$[elem2].status", Value: want.Status},
			}},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: []interface{}{
				bson.D{
					{Key: "elem.package.unique", Value: want.Unique},
					{Key: "elem.sellerusername", Value: want.SellerUsername},
				},
				bson.D{
					{Key: "elem2.prop", Value: want.Prop},
					{Key: "elem2.status", Value: filter},
				},
			},
		}),
	)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("this prop requested")
	}
	return nil
}
