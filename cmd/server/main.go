package main

import (
	"dev.justdrven/loadbalancer/cmd/engine"
	"dev.justdrven/loadbalancer/internal/application"
	"dev.justdrven/loadbalancer/internal/config"
	"dev.justdrven/loadbalancer/internal/orm"
	"dev.justdrven/loadbalancer/internal/service"
)

func main() {
	cnfError, cnf := config.CreateConfig()
	if cnfError != nil {
		panic(cnfError)
	}

	engine.CreateApplication(*cnf)
	app := engine.GetApplication()
	if app == nil {
		panic("The application can't be null!")
	}

	db, ormError := orm.Initialize(*cnf)
	if ormError != nil {
		panic(ormError)
	}

	service.LoadServices(db, app.ServiceType)
	engine.RegisterController()

	application.Start()

}
