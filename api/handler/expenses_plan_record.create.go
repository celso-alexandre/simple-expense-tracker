package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type CreateExpensePlanRecordRequest struct {
	ExpensePlanId uint32 `json:"expense_plan_id"`
	AmountPaid    uint32 `json:"amount_paid"`
	PaymentDate   string `json:"payment_date"`
	PaidDate      string `json:"paid_date"`
}

type CreateExpensePlanRecordResponse struct {
	ExpensePlanRecordId uint32 `json:"expense_plan_record_id"`
}

// CreateExpensePlanRecord godoc
// @Router       /expense-plan-record/create [post]
// @Accept       json
// @Param        request body CreateExpensePlanRecordRequest false "CreateExpensePlanRecordRequest"
// @Produce      json
// @Success      200   {object}  CreateExpensePlanRecordResponse
// @Summary      Create expense-plan-record item
// @Description  Create expense-plan-record item
// @Tags         ExpensePlanRecord
func CreateExpensePlanRecord(w http.ResponseWriter, r *http.Request) {
	var req CreateExpensePlanRecordRequest
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

	item, err := q.CreateExpensePlanRecord(ctx, query.CreateExpensePlanRecordParams{
		ExpensePlanID: int32(req.ExpensePlanId),
		AmountPaid:    int32(req.AmountPaid),
		PaymentDate:   *common.ISOStringToPgTimestamptz(req.PaymentDate),
		PaidDate:      *common.ISOStringToPgTimestamptz(req.PaidDate),
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

	json.NewEncoder(w).Encode(&CreateExpensePlanRecordResponse{ExpensePlanRecordId: uint32(item.ExpensePlanRecordID)})
}
