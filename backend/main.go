package main

import (
	"github.com/adem522/eight-sup/handlers"
	"github.com/adem522/eight-sup/utils"

	"github.com/labstack/echo"
)

func main() {
	client := utils.Connect()
	defer utils.Close(client)
	eventCol := handlers.Collection{
		C1: client.Database("eight-sup").Collection("event"),
		C2: client.Database("eight-sup").Collection("user"),
		N:  "event",
	}
	planCol := handlers.Collection{
		C1: client.Database("eight-sup").Collection("planInfo"),
		N:  "planInfo",
	}
	userCol := handlers.Collection{
		C1: client.Database("eight-sup").Collection("user"),
		N:  "user",
	}
	wantCol := handlers.Collection{
		C1: client.Database("eight-sup").Collection("want"),
		C2: client.Database("eight-sup").Collection("user"),
		N:  "want",
	}
	e := echo.New()
	e.Use(utils.CORSConfig(), utils.Logger()) //issues#6

	user := e.Group("/user")
	user.POST("/register", userCol.Register) //send models.UserStruct
	user.POST("/login", userCol.Login)       //send username and password

	create := e.Group("/create")
	create.POST("/event", eventCol.CreateEvent)      //send models.Event
	create.POST("/planInfo", planCol.CreatePlanInfo) //send models.PlanInfoStruct
	create.GET("/allPlan", planCol.CreateAllPlan)    //just get
	create.POST("/want", wantCol.CreateWant)         //send models.Want

	plan := e.Group("/plan")
	plan.POST("/push", userCol.PushPlan)      //models.UserStructAddPlan
	plan.PATCH("/update", userCol.UpdatePlan) //models.UserStructAddPlan

	list := e.Group("/list")
	list.GET("/user", userCol.ReturnAllUsername)        //just get
	list.GET("/planInfo", planCol.ReturnAllPlanInfo)    //just get
	list.GET("/planUnique", planCol.ReturnPlanUnique)   //just get
	list.PUT("/userStock", userCol.ReturnUserPlanStock) //username
	list.PUT("/event", eventCol.ReturnUserEvent)        //username and boolean type (streamer=true)

	e.Logger.Fatal(e.Start(":8080")) //issues#6
}
