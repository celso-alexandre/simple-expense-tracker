package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type ListExpensePlanRequest struct{}

type ListExpensePlanResponse_ListExpensePlan struct {
	ExpensePlanId      uint32                    `json:"expense_plan_id"`
	Title              string                    `json:"title"`
	Category           query.ExpensePlanCategory `json:"category"`
	AmountPlanned      uint32                    `json:"amount_planned"`
	LastAmountSpent    uint32                    `json:"last_amount_spent"`
	FirstPaymentDate   string                    `json:"first_payment_date"`
	LastPaymentDate    string                    `json:"last_payment_date"`
	LastPaidDate       string                    `json:"last_paid_date"`
	PaidCount          uint32                    `json:"paid_count"`
	RecurrencyType     *query.RecurrencyType     `json:"recurrency_type"`
	RecurrencyInterval uint32                    `json:"recurrency_interval"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListExpensePlanResponse struct {
	Items *[]ListExpensePlanResponse_ListExpensePlan `json:"items"`
}

// ListExpensePlan godoc
// @Router       /expense-plan/list [post]
// @Accept       json
// @Param        request body ListExpensePlanRequest false "ListExpensePlanRequest"
// @Produce      json
// @Success      200   {object}  ListExpensePlanResponse
// @Summary      List all expense-plan items
// @Description  List all expense-plan items (using cursor-based pagination)
// @Tags         ExpensePlan
func ListExpensePlan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	db, err := database.NewCustomDB()
	if err != nil {
		msg := fmt.Sprintf("database.NewCustomDB() failed: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	tx, q, err := db.NewTxQuery(ctx)
	if err != nil {
		msg := fmt.Sprintf("database.NewTxQuery() failed: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(ctx)

	dbItems, err := q.ListExpensePlans(ctx)
	if err != nil {
		msg := fmt.Sprintf("q.ListExpensePlans() failed: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	var items = make([]ListExpensePlanResponse_ListExpensePlan, len(dbItems))
	for i, item := range dbItems {
		items[i] = ListExpensePlanResponse_ListExpensePlan{
			ExpensePlanId:      uint32(item.ExpensePlanID),
			Title:              item.Title,
			AmountPlanned:      uint32(item.AmountPlanned),
			LastAmountSpent:    uint32(item.LastAmountSpent),
			FirstPaymentDate:   common.PgTimestamptzToISOString(&item.FirstPaymentDate),
			LastPaymentDate:    common.PgTimestamptzToISOString(&item.LastPaymentDate),
			LastPaidDate:       common.PgTimestamptzToISOString(&item.LastPaidDate),
			PaidCount:          uint32(item.PaidCount),
			RecurrencyInterval: uint32(item.RecurrencyInterval),
			Category:           item.Category,
			CreatedAt:          common.PgTimestamptzToISOString(&item.CreatedAt),
			UpdatedAt:          common.PgTimestamptzToISOString(&item.UpdatedAt),
		}
		if item.RecurrencyType.Valid {
			items[i].RecurrencyType = &item.RecurrencyType.RecurrencyType
		}
	}

	json.NewEncoder(w).Encode(&ListExpensePlanResponse{Items: &items})
}
