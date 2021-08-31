package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/adem522/eight-sup/database"
	"github.com/adem522/eight-sup/models"
	"github.com/adem522/eight-sup/utils"
	"github.com/labstack/echo"
)

func (col *Collection) CreateEvent(c echo.Context) error {
	u := models.Event{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := database.CreateEvent(&u, col.C1, col.C2); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, []string{
		"Event created Succesfull",
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
	if err := database.PushPlan(&u, col.C1); err != nil {
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
	if err := database.RegisterUser(&u, col.C1); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Register not completed because ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, []string{
		"Register Completed for ",
		u.Username,
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

func (col *Collection) CreateExampleUsers(c echo.Context) error {
	err := database.DropIfExist(col.C1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Register not completed because ": err.Error(),
		})
	}
	u := models.UserStruct{}
	for i := 1; i < 5; i++ {
		if i%2 == 1 {
			u = models.UserStruct{
				Username: "user" + strconv.Itoa(i),
				Password: "1",
				Type:     "streamer",
			}
			u.Plan = []models.PlanStruct{}
		} else {
			u = models.UserStruct{
				Username: "user" + strconv.Itoa(i),
				Password: "1",
				Type:     "client",
			}
			u.Plan = []models.PlanStruct{}
		}
		// Save to database
		if err := database.RegisterUser(&u, col.C1); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"Register not completed because ": err.Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, []string{
		"Register Completed for all example users",
	})
}

func (col *Collection) WantClient(c echo.Context) error {
	u := models.Want{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("error from CreateWant/bind and error = "+err.Error()).Error())
	}
	err := database.ControllerWantClient(&u, col.C1, col.C2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("error from CreateWant/database and error = "+err.Error()).Error())
	}
	return c.JSON(http.StatusCreated, []string{
		"Want created Succesfull",
	})
}

func (col *Collection) WantStreamer(c echo.Context) error {
	u := models.Want{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("error from CompleteWantForStreamer/bind and error = "+err.Error()).Error())
	}
	err := database.ControllerWantStreamer(&u, col.C1, col.C2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("error from CompleteWantForStreamer/database and error = "+err.Error()).Error())
	}
	return c.JSON(http.StatusCreated, []string{
		"Want created Succesfull",
	})
}
