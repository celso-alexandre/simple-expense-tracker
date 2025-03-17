package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type UpdateExpensePlanRequest struct {
	ExpensePlanId      uint32                    `json:"expense_plan_id" validate:"required"`
	Title              string                    `json:"title" validate:"required"`
	Category           query.ExpensePlanCategory `json:"category" validate:"required"`
	AmountPlanned      uint32                    `json:"amount_planned" validate:"required"`
	RecurrencyType     *query.RecurrencyType     `json:"recurrency_type,omitempty"`
	RecurrencyInterval uint32                    `json:"recurrency_interval,omitempty"`
}

type UpdateExpensePlanResponse struct {
	ExpensePlanId uint32 `json:"expense_plan_id"`
}

// UpdateExpensePlan godoc
// @Router       /expense-plan/update [post]
// @Accept       json
// @Param        request body UpdateExpensePlanRequest false "UpdateExpensePlanRequest"
// @Produce      json
// @Success      200   {object}  UpdateExpensePlanResponse
// @Summary      Update expense-plan item
// @Description  Update expense-plan item
// @Tags         ExpensePlan
func UpdateExpensePlan(w http.ResponseWriter, r *http.Request) {
	var req UpdateExpensePlanRequest
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
	if req.RecurrencyType != nil && *req.RecurrencyType != "" {
		recurrencyType.Valid = true
		recurrencyType.RecurrencyType = *req.RecurrencyType
	}
	item, err := q.UpdateExpensePlan(ctx, query.UpdateExpensePlanParams{
		ExpensePlanID:      int32(req.ExpensePlanId),
		Title:              req.Title,
		Category:           query.NullExpensePlanCategory{Valid: true, ExpensePlanCategory: req.Category},
		AmountPlanned:      int32(req.AmountPlanned),
		RecurrencyType:     recurrencyType,
		RecurrencyInterval: int32(req.RecurrencyInterval),
	})
	if err != nil {
		msg := fmt.Sprintf("failed to Update expense plan: %v", err)
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

	json.NewEncoder(w).Encode(&UpdateExpensePlanResponse{ExpensePlanId: uint32(item.ExpensePlanID)})
}
