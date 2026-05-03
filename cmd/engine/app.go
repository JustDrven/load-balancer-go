package engine

import (
	"fmt"
	"net/http"

	"dev.justdrven/loadbalancer/internal/application"
	"dev.justdrven/loadbalancer/internal/config"
	"dev.justdrven/loadbalancer/pkg"
)

var app *application.Application

func CreateApplication(cnf config.Config) {
	serviceType := cnf.ServiceType

	app = &application.Application{
		ServiceType: serviceType,
	}
}

func GetApplication() application.Application {
	if app == nil {
		panic("[APP-ENGINE] The application isn't initialized")
	}

	return *app
}

func Start() {
	port := pkg.PORT

	fmt.Printf("[APP-ENGINE] The application is listening at port %d\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
