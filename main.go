package main

import (
	routes "api/router"
	"fmt"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	fmt.Println("Start server")
	http.ListenAndServe(":1234", router)
}
