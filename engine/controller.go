package engine

import (
	"fmt"
	"net/http"
)

func RegisterController() {
	http.HandleFunc("/", middleware)
	fmt.Println("[APP-ENGINE] Loaded service controller")

}

func middleware(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world"))
}
