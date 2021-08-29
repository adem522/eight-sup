package database

import (
	"context"
	"errors"

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

func ReturnAllArray(collection *mongo.Collection) (interface{}, error) {
	opt := options.Find().SetProjection(bson.M{"plan.package.items": 1})
	filterCursor, err := collection.Find(context.TODO(), bson.M{}, opt)
	if err != nil {
		return nil, errors.New("error from ReturnAll/filterCursor.All and error code=" + err.Error())
	}
	var filtered []bson.M
	if err = filterCursor.All(context.TODO(), &filtered); err != nil {
		return nil, errors.New("error from ReturnAll/filterCursor.All and error code=" + err.Error())
	}
	return filtered, nil
}
