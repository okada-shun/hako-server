package main

import (
	"hako-server/config"
	"hako-server/database"
	"hako-server/eventlogs"
	"hako-server/hkfinance"
	"hako-server/router"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	db, err := database.NewDatabase(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	hkfinanceTx, err := hkfinance.NewHkfinanceTx(config)
	if err != nil {
		panic(err)
	}
	eventDatabase := eventlogs.EventDatabase{
		DB: db,
	}
	err = eventDatabase.ReadEventLogs(config)
	if err != nil {
		panic(err)
	}
	r := router.CreateRouter(db, hkfinanceTx)
	r.Run(":8080")
}
