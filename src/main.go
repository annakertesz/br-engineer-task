package main

import (
	"fmt"
	"github.com/annakertesz/br-engineer-task/src/config"
	"github.com/annakertesz/br-engineer-task/src/controller"
	"github.com/annakertesz/br-engineer-task/src/persistence"
	"time"
)

var configfile = "src/config/limit_config.json"


func main() {
	//get config
	config, err := config.GetConfigFromFile(configfile)
	if err != nil {
		panic("Could't load config file")
	}

	var db persistence.Persistence
	var c controller.Controller
	db = persistence.NewDumbPersistence()
	c = controller.NewDumbController(db, *config)

	fmt.Println("*Create users*")
	fmt.Println("*Create public app*")
	userA := c.CreateUser("Myfirst User", "free")
	c.CreateUser("MySecond User", "organization")
	appA, _ := c.CreateApp(userA.GetId(), "First Public", true)
	db.Print()

	fmt.Print("*Public app limits:   ")
	fmt.Println(c.GetLimit(appA.GetId()))
	c.ChangeLimits(appA.GetId(), 10, time.Hour,10, 10)
	fmt.Print("*Updated app limits:  ")
	fmt.Println(c.GetLimit(appA.GetId()))
	c.UsePrivateLimits(appA.GetId())
	fmt.Print("*use private limits:   ")
	fmt.Println(c.GetLimit(appA.GetId()))
	fmt.Println("*Create private app*")
	c.CreateApp(userA.GetId(), "Second Private", false)
	db.Print()
}