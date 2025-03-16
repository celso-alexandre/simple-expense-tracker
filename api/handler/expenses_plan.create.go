package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type CreateExpensePlanRequest struct {
	Title          string                    `json:"title" validate:"required"`
	Category       query.ExpensePlanCategory `json:"category" validate:"required"`
	AmountPlanned  uint32                    `json:"amount_planned" validate:"required"`
	RecurrencyType *query.RecurrencyType     `json:"recurrency_type" validate:"required"`
}

type CreateExpensePlanResponse struct {
	ExpensePlanId uint32 `json:"expense_plan_id"`
}

// CreateExpensePlan godoc
// @Router       /expense-plan/create [post]
// @Accept       json
// @Param        request body CreateExpensePlanRequest false "CreateExpensePlanRequest"
// @Produce      json
// @Success      200   {object}  CreateExpensePlanResponse
// @Summary      Create expense-plan item
// @Description  Create expense-plan item
// @Tags         ExpensePlan
func CreateExpensePlan(w http.ResponseWriter, r *http.Request) {
	var req CreateExpensePlanRequest
	err := common.ParseReqBody(r, &req)
	if err != nil {
		msg := fmt.Sprintf("failed to parse request body: %v", err)
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

	recurrencyType := query.NullRecurrencyType{}
	if req.RecurrencyType != nil {
		recurrencyType.Valid = true
		recurrencyType.RecurrencyType = *req.RecurrencyType
	}
	item, err := q.CreateExpensePlan(ctx, query.CreateExpensePlanParams{
		Title:          req.Title,
		Category:       query.NullExpensePlanCategory{Valid: true, ExpensePlanCategory: req.Category},
		AmountPlanned:  int32(req.AmountPlanned),
		RecurrencyType: recurrencyType,
	})
	if err != nil {
		msg := fmt.Sprintf("failed to create expense plan: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		msg := fmt.Sprintf("failed to commit transaction: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&CreateExpensePlanResponse{ExpensePlanId: uint32(item.ExpensePlanID)})
}
