package engine

import (
	"dev.justdrven/loadbalancer/internal/application"
	"dev.justdrven/loadbalancer/internal/config"
)

var app *application.Application

func CreateApplication(cnf config.Config) {
	serviceType := cnf.ServiceType

	app = &application.Application{
		ServiceType: serviceType,
	}
}

func GetApplication() *application.Application {
	if app == nil {
		panic("[APP-ENGINE] The application isn't initialized")
	}

	return app
}
