package main

import (
	"dev.justdrven/loadbalancer/engine"
	"dev.justdrven/loadbalancer/manager"
)

func main() {

	err, cnf := manager.CCreateConfig()
	if err != nil {
		panic(err)
	}

	engine.CreateApplication(*cnf)

	manager.OrmInit(*cnf)
	manager.SvcLoadServices()

	engine.Start()

}
