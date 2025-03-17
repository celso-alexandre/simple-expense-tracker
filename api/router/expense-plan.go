package router

import (
	"net/http"

	"github.com/celso-alexandre/api/handler"
)

func ExpensePlanSetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/expense-plan/list", handler.ListExpensePlan)
	mux.HandleFunc("/expense-plan/get", handler.GetExpensePlan)
	mux.HandleFunc("/expense-plan/create", handler.CreateExpensePlan)
	mux.HandleFunc("/expense-plan/update", handler.UpdateExpensePlan)
	mux.HandleFunc("/expense-plan/delete", handler.DeleteExpensePlan)
}
