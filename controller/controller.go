package controller

import (
	"net/http"

	"api/services"
)

type ApiResponse struct {
	HTTPStatusCode string
	ResultMessage  string
}

func Endpoint1(w http.ResponseWriter, r *http.Request) {
	response := ApiResponse{"200", sayHello("world")}
	services.ResponseWithJson(w, http.StatusOK, response)
}

func Endpoint2(w http.ResponseWriter, r *http.Request) {
	response := ApiResponse{"200", sayHello("go")}
	services.ResponseWithJson(w, http.StatusOK, response)
}

func Endpoint3(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["prm"]
	prm := ""
	if ok {
		prm = params[0]
	}
	response := ApiResponse{"200", sayHello(prm)}
	services.ResponseWithJson(w, http.StatusOK, response)
}

func sayHello(name string) string {
	return "hello " + name
}
