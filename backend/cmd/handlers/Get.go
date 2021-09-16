package handlers

import (
	"net/http"

	"github.com/adem522/eight-sup/models"
	"github.com/adem522/eight-sup/pkg/db"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collection) ReturnAllPlanInfo(c echo.Context) error {
	data, err := db.ReturnAll(col.C1, "", nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
func (col *Collection) ReturnAllUsername(c echo.Context) error {
	data, err := db.ReturnAll(col.C1, "username", bson.M{"type": "streamer"})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (col *Collection) ReturnUserPlan(c echo.Context) error {
	temp := struct {
		Username string `bson:"username"`
	}{}
	err := c.Bind(&temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	data, err := db.ReturnAll(col.C1, "plan", bson.M{"username": temp.Username})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (col *Collection) ReturnUserEvent(c echo.Context) error {
	temp := struct {
		Username       string `bson:"username,omitempty" json:"username"`
		BuyerUsername  string `bson:"buyerUsername,omitempty"`
		SellerUsername string `bson:"sellerUsername,omitempty"`
		Type           string `bson:"type,omitempty"`
	}{}
	if err := c.Bind(&temp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	filter := bson.M{}
	if temp.Type != "" {
		if temp.Type == "streamer" {
			filter = bson.M{
				"sellerUsername": temp.Username,
			}
		} else {
			filter = bson.M{
				"buyerUsername": temp.Username,
			}
		}
	} else {
		filter = bson.M{
			"sellerUsername": temp.SellerUsername,
			"buyerUsername":  temp.BuyerUsername,
		}
	}
	data, err := db.ReturnAll(col.C1, "", filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (col *Collection) ReturnPlanUnique(c echo.Context) error {
	data, err := db.ReturnAll(col.C1, "unique", nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (col *Collection) ReturnUserWants(c echo.Context) error {
	var temp struct {
		BuyerUsername  string `bson:"buyerUsername"`
		SellerUsername string `bson:"sellerUsername"`
	}
	if err := c.Bind(&temp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	data, err := db.ReturnAll(
		col.C1, "plan.package.items", //projection
		bson.M{
			"type":                             "streamer",
			"username":                         temp.SellerUsername,
			"plan.package.items.buyerUsername": temp.BuyerUsername,
		}, //filter
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (col *Collection) ReturnAllItemsForClient(c echo.Context) error {
	var temp models.Want
	if err := c.Bind(&temp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	data, err := db.ReturnAllItemsForClient(col.C1, &temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}
func (col *Collection) ReturnAllItemsForStreamer(c echo.Context) error {
	var temp models.Want
	if err := c.Bind(&temp); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	data, err := db.ReturnAllItemsForStreamer(col.C1, &temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}
