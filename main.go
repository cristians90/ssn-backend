package main

import (
	"net/http"
	"ssnbackend/routes"
)

func main() {
	printAppInfo()
	http.ListenAndServe(":3000", routes.GetRoutes())
}
