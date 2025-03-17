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

type GetExpensePlanResponse_ExpensePlanRecord struct {
	ExpensePlanRecordId uint32 `json:"expense_plan_record_id"`
	AmountPaid          uint32 `json:"amount_paid"`
	PaymentDate         string `json:"payment_date"`
	PaidDate            string `json:"paid_date"`
	ExpensePlanSequence uint32 `json:"expense_plan_sequence"`
}

type GetExpensePlanResponse struct {
	ExpensePlanId      uint32                    `json:"expense_plan_id"`
	Title              string                    `json:"title"`
	AmountPlanned      uint32                    `json:"amount_planned"`
	RecurrencyType     *query.RecurrencyType     `json:"recurrency_type"`
	RecurrencyInterval uint32                    `json:"recurrency_interval"`
	Category           query.ExpensePlanCategory `json:"category"`

	FirstExpensePlanRecordId uint32                                    `json:"first_expense_plan_record_id"`
	FirstExpensePlanRecord   *GetExpensePlanResponse_ExpensePlanRecord `json:"first_expense_plan_record"`
	LastExpensePlanRecordId  uint32                                    `json:"last_expense_plan_record_id"`
	LastExpensePlanRecord    *GetExpensePlanResponse_ExpensePlanRecord `json:"last_expense_plan_record"`

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
		ExpensePlanId:            uint32(dbItem.ExpensePlanID),
		Title:                    dbItem.Title,
		AmountPlanned:            uint32(dbItem.AmountPlanned),
		RecurrencyInterval:       uint32(dbItem.RecurrencyInterval),
		Category:                 dbItem.Category,
		RecurrencyType:           nil,
		FirstExpensePlanRecordId: uint32(dbItem.FirstExpensePlanRecordID.Int32),
		FirstExpensePlanRecord:   nil,
		LastExpensePlanRecordId:  uint32(dbItem.LastExpensePlanRecordID.Int32),
		LastExpensePlanRecord:    nil,
		CreatedAt:                common.PgTimestamptzToISOString(&dbItem.CreatedAt),
		UpdatedAt:                common.PgTimestamptzToISOString(&dbItem.UpdatedAt),
	}
	if dbItem.RecurrencyType.Valid {
		response.RecurrencyType = &dbItem.RecurrencyType.RecurrencyType
	}
	if dbItem.FirstExpensePlanRecordID.Valid {
		response.FirstExpensePlanRecord = &GetExpensePlanResponse_ExpensePlanRecord{
			ExpensePlanRecordId: uint32(dbItem.FirstExpensePlanRecordID.Int32),
			AmountPaid:          uint32(dbItem.FirstAmountPaid.Int32),
			PaymentDate:         common.PgTimestamptzToISOString(&dbItem.FirstPaymentDate),
			PaidDate:            common.PgTimestamptzToISOString(&dbItem.FirstPaidDate),
			ExpensePlanSequence: uint32(dbItem.FirstExpensePlanSequence.Int32),
		}
	}
	if dbItem.LastExpensePlanRecordID.Valid {
		response.LastExpensePlanRecord = &GetExpensePlanResponse_ExpensePlanRecord{
			ExpensePlanRecordId: uint32(dbItem.LastExpensePlanRecordID.Int32),
			AmountPaid:          uint32(dbItem.LastAmountPaid.Int32),
			PaymentDate:         common.PgTimestamptzToISOString(&dbItem.LastPaymentDate),
			PaidDate:            common.PgTimestamptzToISOString(&dbItem.LastPaidDate),
			ExpensePlanSequence: uint32(dbItem.LastExpensePlanSequence.Int32),
		}
	}

	json.NewEncoder(w).Encode(response)
}
