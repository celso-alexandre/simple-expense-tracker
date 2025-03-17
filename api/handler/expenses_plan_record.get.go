package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type GetExpensePlanRecordRequest struct {
	ExpensePlanRecordId uint32 `json:"expense_plan_id" validate:"required"`
}

type GetExpensePlanRecordResponse_ExpensePlan struct {
	ExpensePlanId      uint32                `json:"expense_plan_id"`
	Title              string                `json:"title"`
	Category           string                `json:"category"`
	AmountPlanned      uint32                `json:"amount_planned"`
	RecurrencyType     *query.RecurrencyType `json:"recurrency_type"`
	RecurrencyInterval uint32                `json:"recurrency_interval"`
}

type GetExpensePlanRecordResponse struct {
	ExpensePlanRecordId uint32                                    `json:"expense_plan_record_id"`
	ExpensePlanId       uint32                                    `json:"expense_plan_id"`
	ExpensePlan         *GetExpensePlanRecordResponse_ExpensePlan `json:"expense_plan"`
	AmountPaid          uint32                                    `json:"amount_paid"`
	PaymentDate         string                                    `json:"payment_date"`
	PaidDate            string                                    `json:"paid_date"`
	ExpensePlanSequence uint32                                    `json:"expense_plan_sequence"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// GetExpensePlanRecord godoc
// @Router       /expense-plan-record/get [post]
// @Accept       json
// @Param        request body GetExpensePlanRecordRequest false "GetExpensePlanRecordRequest"
// @Produce      json
// @Success      200   {object}  GetExpensePlanRecordResponse
// @Summary      Get all expense-plan-record items
// @Description  Get all expense-plan-record items (using cursor-based pagination)
// @Tags         ExpensePlanRecord
func GetExpensePlanRecord(w http.ResponseWriter, r *http.Request) {
	var req GetExpensePlanRecordRequest
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

	dbItem, err := q.GetExpensePlanRecord(ctx, int32(req.ExpensePlanRecordId))
	if err != nil {
		msg := fmt.Sprintf("q.GetExpensePlanRecord() failed: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	response := &GetExpensePlanRecordResponse{
		ExpensePlanRecordId: uint32(dbItem.ExpensePlanRecordID),
		ExpensePlanId:       uint32(dbItem.ExpensePlanID),
		AmountPaid:          uint32(dbItem.AmountPaid),
		PaymentDate:         common.PgTimestamptzToISOString(&dbItem.PaymentDate),
		PaidDate:            common.PgTimestamptzToISOString(&dbItem.PaidDate),
		ExpensePlanSequence: uint32(dbItem.ExpensePlanSequence),
		ExpensePlan: &GetExpensePlanRecordResponse_ExpensePlan{
			ExpensePlanId:      uint32(dbItem.ExpensePlanID),
			Title:              dbItem.ExpensePlanTitle.String,
			Category:           string(dbItem.ExpensePlanCategory.ExpensePlanCategory),
			AmountPlanned:      uint32(dbItem.ExpensePlanAmountPlanned.Int32),
			RecurrencyType:     &dbItem.ExpensePlanRecurrencyType.RecurrencyType,
			RecurrencyInterval: uint32(dbItem.ExpensePlanRecurrencyInterval.Int32),
		},
		CreatedAt: common.PgTimestamptzToISOString(&dbItem.CreatedAt),
		UpdatedAt: common.PgTimestamptzToISOString(&dbItem.UpdatedAt),
	}

	json.NewEncoder(w).Encode(response)
}
