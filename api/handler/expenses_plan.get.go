package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type GetExpensePlanRequest struct {
	ExpensePlanId uint32 `json:"expense_plan_id" validate:"required"`
}

type GetExpensePlanResponse struct {
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

// GetExpensePlan godoc
// @Router       /expense-plan/get [post]
// @Accept       json
// @Param        request body GetExpensePlanRequest false "GetExpensePlanRequest"
// @Produce      json
// @Success      200   {object}  GetExpensePlanResponse
// @Summary      Get all expense-plan items
// @Description  Get all expense-plan items (using cursor-based pagination)
// @Tags         ExpensePlan
func GetExpensePlan(w http.ResponseWriter, r *http.Request) {
	var req GetExpensePlanRequest
	err := common.ParseReqBody(r, &req)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse request body: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	err = validate.Struct(req)
	if err != nil {
		msg := fmt.Sprintf("validate.Struct failed: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

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

	dbItem, err := q.GetExpensePlan(ctx, int32(req.ExpensePlanId))
	if err != nil {
		msg := fmt.Sprintf("q.GetExpensePlan() failed: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	response := &GetExpensePlanResponse{
		ExpensePlanId:      uint32(dbItem.ExpensePlanID),
		Title:              dbItem.Title,
		AmountPlanned:      uint32(dbItem.AmountPlanned),
		LastAmountSpent:    uint32(dbItem.LastAmountSpent),
		FirstPaymentDate:   common.PgTimestamptzToISOString(&dbItem.FirstPaymentDate),
		LastPaymentDate:    common.PgTimestamptzToISOString(&dbItem.LastPaymentDate),
		LastPaidDate:       common.PgTimestamptzToISOString(&dbItem.LastPaidDate),
		PaidCount:          uint32(dbItem.PaidCount),
		RecurrencyInterval: uint32(dbItem.RecurrencyInterval),
		Category:           dbItem.Category,
		CreatedAt:          common.PgTimestamptzToISOString(&dbItem.CreatedAt),
		UpdatedAt:          common.PgTimestamptzToISOString(&dbItem.UpdatedAt),
	}
	if dbItem.RecurrencyType.Valid {
		response.RecurrencyType = &dbItem.RecurrencyType.RecurrencyType
	}

	json.NewEncoder(w).Encode(response)
}
