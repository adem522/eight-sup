package main

import (
	"os"

	"github.com/adem522/eight-sup/cmd/handlers"
	"github.com/adem522/eight-sup/pkg/utils"

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
	wantCol := handlers.Collection{
		C1: client.Database("eight-sup").Collection("want"),
		C2: client.Database("eight-sup").Collection("user"),
		N:  "want",
	}
	planCol := handlers.Collection{
		C1: client.Database("eight-sup").Collection("planInfo"),
		N:  "planInfo",
	}
	userCol := handlers.Collection{
		C1: client.Database("eight-sup").Collection("user"),
		N:  "user",
	}
	e := echo.New()
	e.Use(utils.CORSConfig(), utils.Logger()) //issues#6

	user := e.Group("/user")
	user.POST("/register", userCol.Register)                    //send models.UserStruct
	user.POST("/login", userCol.Login)                          //send username and password
	user.GET("/createExampleUsers", userCol.CreateExampleUsers) //just get

	create := e.Group("/create")
	create.GET("/allPlan", planCol.CreateAllPlan)      //just get
	create.POST("/event", eventCol.CreateEvent)        //send models.Event
	create.POST("/planInfo", planCol.CreatePlanInfo)   //send models.PlanInfoStruct
	create.POST("/wantClient", wantCol.WantClient)     //send models.Want
	create.POST("/wantStreamer", wantCol.WantStreamer) //send models.Want

	plan := e.Group("/plan")
	plan.POST("/push", userCol.PushPlan)      //models.UserStructAddPlan
	plan.PATCH("/update", userCol.UpdatePlan) //models.UserStructAddPlan

	list := e.Group("/list")
	list.GET("/user", userCol.ReturnAllUsername)                        //just get
	list.GET("/planInfo", planCol.ReturnAllPlanInfo)                    //just get
	list.GET("/planUnique", planCol.ReturnPlanUnique)                   //just get
	list.PUT("/userPlan", userCol.ReturnUserPlan)                       //username
	list.PUT("/event", eventCol.ReturnUserEvent)                        //username and boolean type (streamer=true)
	list.PUT("/userWants", userCol.ReturnUserWants)                     //buyer and seller username
	list.PUT("/allItemsForClient", userCol.ReturnAllItemsForClient)     //buyer and seller username
	list.PUT("/allItemsForStreamer", userCol.ReturnAllItemsForStreamer) //buyer and seller username
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))                    //issues#6
}
