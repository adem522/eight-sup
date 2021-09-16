package handlers

import (
	"log"
	"net/http"

	"github.com/adem522/eight-sup/models"
	"github.com/adem522/eight-sup/pkg/db"
	"github.com/labstack/echo"
)

func (col *Collection) UpdatePlan(c echo.Context) error {
	u := models.UserStructAddPlan{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	err := db.UpdatePlan(&u, col.C1)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Plan updated")
}
