package application

import (
	"fmt"
	"net/http"

	"dev.justdrven/loadbalancer/pkg"
)

func Start() {
	port := pkg.PORT

	fmt.Printf("[APP-ENGINE] The application is listening at port %d\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
