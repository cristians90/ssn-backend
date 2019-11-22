package ssn_backend

import (
	"net/http"
	"ssn-backend/routes"
)

func main() {
	printAppInfo()
	http.ListenAndServe(":3000", routes.GetRoutes())
}
