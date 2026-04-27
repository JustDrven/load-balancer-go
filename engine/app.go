package engine

import (
	"fmt"
	"net/http"

	"dev.justdrven/loadbalancer/data"
)

const PORT = 9090

var application *data.Application

func CreateApplication(cnf data.Config) {
	serviceType := cnf.ServiceType

	application = &data.Application{
		ServiceType: serviceType,
	}
}

func GetApplication() data.Application {
	if application == nil {
		panic("[APP-ENGINE] The application isn't initialized")
	}

	return *application
}

func Start() {
	fmt.Printf("[APP-ENGINE] The application is listening at port %d\n", PORT)

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
