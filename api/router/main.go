package router

import (
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	ExpensePlanSetupRoutes(mux)
	ExpensePlanRecordSetupRoutes(mux)

	return mux
}
