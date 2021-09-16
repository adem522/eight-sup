package db

import (
	"context"
	"errors"

	"github.com/adem522/eight-sup/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReturnAll(collection *mongo.Collection, projection string, filter bson.M) (interface{}, error) {
	var findOpt *options.FindOptions
	var filterCursor *mongo.Cursor
	var err error
	if projection != "" {
		findOpt = options.Find().SetProjection(bson.M{"_id": 0, projection: 1})
	}
	if filter != nil {
		filterCursor, err = collection.Find(context.TODO(), filter, findOpt)
	} else {
		filterCursor, err = collection.Find(context.TODO(), bson.M{}, findOpt)
	}
	if err != nil {
		return nil, errors.New("error from ReturnAll/find and error code= " + err.Error())
	}
	var filtered []bson.M
	if err = filterCursor.All(context.TODO(), &filtered); err != nil {
		return nil, errors.New("error from ReturnAll/filterCursor.All and error code= " + err.Error())
	}
	return filtered, nil
}

func ReturnAllItemsForClient(collection *mongo.Collection, want *models.Want) (interface{}, error) {
	temp := []models.UserStruct{}
	filterCursor, err := collection.Find(context.TODO(), bson.M{
		"username": want.BuyerUsername,
		//"plan.package.items.buyerUsername": want.SellerUsername,
	})
	if err != nil {
		return nil, errors.New("error from ReturnAllItemsForClient/filterCursor and error code=" + err.Error())
	}
	if err = filterCursor.All(context.TODO(), &temp); err != nil {
		return nil, errors.New("error from ReturnAllItemsForClient/filterCursor.All and error code=" + err.Error())
	}
	wants := []models.Want{}
	wants2 := models.Want{}
	for _, data := range temp[0].Plan {
		if data.Package.Items != nil {
			for _, data2 := range data.Package.Items {
				if want.Status != "" && want.Status == data2.Status {
					wants2.SellerUsername = data.SellerUsername
					wants2.Unique = data.Package.Unique
					wants2.Prop = data2.Prop
					wants2.BuyerUsername = want.BuyerUsername
					wants2.Status = data2.Status
					wants = append(wants, wants2)
				}
			}
		}
	}
	return wants, nil
}
func ReturnAllItemsForStreamer(collection *mongo.Collection, want *models.Want) (interface{}, error) {
	temp := []models.UserStruct{}
	filterCursor, err := collection.Find(context.TODO(), bson.M{
		"username": want.SellerUsername,
		//"plan.package.items.buyerUsername": want.SellerUsername,
	})
	if err != nil {
		return nil, errors.New("error from ReturnAllItemsForClient/filterCursor and error code=" + err.Error())
	}
	if err = filterCursor.All(context.TODO(), &temp); err != nil {
		return nil, errors.New("error from ReturnAllItemsForClient/filterCursor.All and error code=" + err.Error())
	}
	wants := []models.Want{}
	wants2 := models.Want{}
	for _, data := range temp[0].Plan {
		if data.Package.Items != nil {
			for _, data2 := range data.Package.Items {
				if want.Status != "" && want.Status == data2.Status {
					wants2.SellerUsername = data.SellerUsername
					wants2.Unique = data.Package.Unique
					wants2.Prop = data2.Prop
					wants2.BuyerUsername = want.BuyerUsername
					wants2.Status = data2.Status
					wants = append(wants, wants2)
				}
			}
		}
	}
	return wants, nil
}
