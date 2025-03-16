package router

import (
	"net/http"

	"github.com/celso-alexandre/api/handler"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/expense-plan/list", handler.ListExpensePlan)
	mux.HandleFunc("/expense-plan/create", handler.CreateExpensePlan)

	return mux
}
