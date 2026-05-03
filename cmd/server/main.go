package main

import (
	"dev.justdrven/loadbalancer/cmd/engine"
	"dev.justdrven/loadbalancer/internal/config"
	"dev.justdrven/loadbalancer/internal/orm"
	"dev.justdrven/loadbalancer/internal/service"
)

func main() {

	err, cnf := config.CCreateConfig()
	if err != nil {
		panic(err)
	}

	engine.CreateApplication(*cnf)

	db := orm.OrmInit(*cnf)
	service.SvcLoadServices(*db, engine.GetApplication().ServiceType)

	engine.RegisterController()
	engine.Start()

}
