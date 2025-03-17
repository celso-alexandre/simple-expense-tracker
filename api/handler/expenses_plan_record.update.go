package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type UpdateExpensePlanRecordRequest struct {
	ExpensePlanRecordId uint32 `json:"expense_plan_record_id"`
	ExpensePlanId       uint32 `json:"expense_plan_id"`
	AmountPaid          uint32 `json:"amount_paid"`
	PaymentDate         string `json:"payment_date"`
	PaidDate            string `json:"paid_date"`
}

type UpdateExpensePlanRecordResponse struct {
	ExpensePlanRecordId uint32 `json:"expense_plan_record_id"`
}

// UpdateExpensePlanRecord godoc
// @Router       /expense-plan-record/update [post]
// @Accept       json
// @Param        request body UpdateExpensePlanRecordRequest false "UpdateExpensePlanRecordRequest"
// @Produce      json
// @Success      200   {object}  UpdateExpensePlanRecordResponse
// @Summary      Update expense-plan-record item
// @Description  Update expense-plan-record item
// @Tags         ExpensePlanRecord
func UpdateExpensePlanRecord(w http.ResponseWriter, r *http.Request) {
	var req UpdateExpensePlanRecordRequest
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

	item, err := q.UpdateExpensePlanRecord(ctx, query.UpdateExpensePlanRecordParams{
		ExpensePlanRecordID: *common.Uint32ToPgInt(req.ExpensePlanRecordId),
		ExpensePlanID:       *common.Uint32ToPgInt(req.ExpensePlanId),
		AmountPaid:          *common.Uint32ToPgInt(req.AmountPaid),
		PaymentDate:         *common.ISOStringToPgTimestamptz(req.PaymentDate),
		PaidDate:            *common.ISOStringToPgTimestamptz(req.PaidDate),
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

	json.NewEncoder(w).Encode(&UpdateExpensePlanRecordResponse{ExpensePlanRecordId: uint32(item.ExpensePlanRecordID)})
}
