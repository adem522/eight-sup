package handlers

import (
	"net/http"

	"github.com/adem522/eight-sup/database"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collection) ReturnAllPlanInfo(c echo.Context) error {
	data, err := database.ReturnAll(col.C1, "", nil)
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
	data, err := database.ReturnAll(col.C1, "username", bson.M{"type": "streamer"})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (col *Collection) ReturnUserPlanStock(c echo.Context) error {
	temp := struct {
		Username string `bson:"username"`
	}{}
	err := c.Bind(&temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	data, err := database.ReturnAll(col.C1, "plan", bson.M{"username": temp.Username})
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
	filter := bson.M{}
	err := c.Bind(&temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
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
	data, err := database.ReturnAll(col.C1, "", filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (col *Collection) ReturnPlanUnique(c echo.Context) error {
	data, err := database.ReturnAll(col.C1, "unique", nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't return user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}
