package handlers

import (
	"net/http"

	"github.com/adem522/eight-sup/database"
	"github.com/adem522/eight-sup/models"
	"github.com/adem522/eight-sup/utils"
	"github.com/labstack/echo"
)

func (col *Collection) CreateEvent(c echo.Context) error {
	u := models.Event{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	err := database.CreateEvent(&u, col.C1, col.C2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, []string{
		"Event created Succesfull",
	})
}

func (col *Collection) CreateWant(c echo.Context) error {
	u := models.Want{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	data, err := database.CreateWant(&u, col.C1, col.C2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, []interface{}{
		"Want created Succesfull",
		data,
	})
}

func (col *Collection) CreatePlanInfo(c echo.Context) error {
	u := models.PlanInfoStruct{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	result, err := database.CreatePlanInfo(&u, col.C1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Event not created - "+err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

func (col *Collection) CreateAllPlan(c echo.Context) error {
	result, err := database.CreateAllPlan(col.C1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Plan Infos not created - "+err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (col *Collection) PushPlan(c echo.Context) error {
	u := models.UserStructAddPlan{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	err := database.PushPlan(&u, col.C1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Plan added")
}

func (col *Collection) Register(c echo.Context) error {
	// Binding context data
	u := models.UserStruct{}
	u.Plan = []models.PlanStruct{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	// Save to database
	err := database.RegisterUser(&u, col.C1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Register not completed because ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Register Completed for ": u.Username,
	})
}

func (col *Collection) Login(c echo.Context) error {
	// Binding context data
	user := models.UserStruct{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	check := database.LoginCheck(user.Username, user.Password, col.C1)
	// Throws unauthorized error
	if check == "" {
		return echo.ErrUnauthorized
	}
	//Create token
	token := utils.CreateToken(user)
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
		"type":  check,
	})
}
