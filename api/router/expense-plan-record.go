package router

import (
	"net/http"

	"github.com/celso-alexandre/api/handler"
)

func ExpensePlanRecordSetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/expense-plan-record/list", handler.ListExpensePlanRecord)
}
