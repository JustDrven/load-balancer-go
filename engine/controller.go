package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dev.justdrven/loadbalancer/data"
	"dev.justdrven/loadbalancer/manager"
)

func RegisterController() {
	http.HandleFunc("/", middleware)

	fmt.Println("[APP-ENGINE] Loaded service controller")
}

func middleware(res http.ResponseWriter, req *http.Request) {
	service := manager.RefGetBestService()
	if service == nil {
		res.WriteHeader(http.StatusNotFound)
		setResponseBody(res, data.ErrorResponse{
			Error:   true,
			Message: "The service is unavailable!",
		})
		return
	}

	manager.RefUse(service)
	url := createURLToService(*service, *req)

	resp, err := sendRequest(url, *req)
	if err != nil {

		res.WriteHeader(http.StatusInternalServerError)
		setResponseBody(res, data.ErrorResponse{
			Error:   true,
			Message: "The request failed",
		})

		return
	}

	defer resp.Body.Close()
	for k, v := range resp.Header {
		for _, val := range v {
			res.Header().Add(k, val)
		}
	}

	res.WriteHeader(resp.StatusCode)
	io.Copy(res, resp.Body)

	manager.RefClose(service)
}

func sendRequest(addr string, oldReq http.Request) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(oldReq.Method, addr, oldReq.Body)
	if err != nil {
		return nil, err
	}

	setupHeaders(oldReq, req)

	return client.Do(req)
}

func setupHeaders(oldReq http.Request, newRequest *http.Request) {

	for key, values := range oldReq.Header {
		for _, v := range values {
			newRequest.Header.Add(key, v)
		}
	}

}

func setResponseBody(newResponse http.ResponseWriter, responseBody any) {
	newResponse.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(newResponse)

	jsonEncoder.Encode(responseBody)
}

func createURLToService(service data.ManagedService, req http.Request) string {
	addr := service.Address
	endpoint := req.URL.Path
	query := req.URL.RawQuery

	if len(query) != 0 {
		query = "?" + query
	} else {
		query = ""
	}

	return fmt.Sprintf("%s%s%s", addr, endpoint, query)
}
