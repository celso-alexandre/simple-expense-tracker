package router

import (
	"net/http"

	"github.com/celso-alexandre/api/handler"
)

func ExpensePlanRecordSetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/expense-plan-record/list", handler.ListExpensePlanRecord)
	mux.HandleFunc("/expense-plan-record/get", handler.GetExpensePlanRecord)
	mux.HandleFunc("/expense-plan-record/create", handler.CreateExpensePlanRecord)
	mux.HandleFunc("/expense-plan-record/update", handler.UpdateExpensePlanRecord)
	mux.HandleFunc("/expense-plan-record/delete", handler.DeleteExpensePlanRecord)
}
