package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
	"github.com/celso-alexandre/api/query"
)

type ListExpensePlanRecordRequest struct{}

type ListExpensePlanRecordResponse_ListExpensePlan struct {
	ExpensePlanId      uint32                `json:"expense_plan_id"`
	Title              string                `json:"title"`
	Category           string                `json:"category"`
	AmountPlanned      uint32                `json:"amount_planned"`
	RecurrencyType     *query.RecurrencyType `json:"recurrency_type"`
	RecurrencyInterval uint32                `json:"recurrency_interval"`
}

type ListExpensePlanRecordResponse_ListExpensePlanRecord struct {
	ExpensePlanRecordId uint32                                         `json:"expense_plan_record_id"`
	ExpensePlanId       uint32                                         `json:"expense_plan_id"`
	ExpensePlan         *ListExpensePlanRecordResponse_ListExpensePlan `json:"expense_plan"`
	AmountPaid          uint32                                         `json:"amount_paid"`
	PaymentDate         string                                         `json:"payment_date"`
	PaidDate            string                                         `json:"paid_date"`
	ExpensePlanSequence uint32                                         `json:"expense_plan_sequence"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListExpensePlanRecordResponse struct {
	Items *[]ListExpensePlanRecordResponse_ListExpensePlanRecord `json:"items"`
}

// ListExpensePlanRecord godoc
// @Router       /expense-plan-record/list [post]
// @Accept       json
// @Param        request body ListExpensePlanRecordRequest false "ListExpensePlanRecordRequest"
// @Produce      json
// @Success      200   {object}  ListExpensePlanRecordResponse
// @Summary      List all expense-plan-record items
// @Description  List all expense-plan-record items (using cursor-based pagination)
// @Tags         ExpensePlanRecord
func ListExpensePlanRecord(w http.ResponseWriter, r *http.Request) {
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

	dbItems, err := q.ListExpensePlanRecords(ctx)
	if err != nil {
		msg := fmt.Sprintf("q.ListExpensePlanRecords() failed: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	var items = make([]ListExpensePlanRecordResponse_ListExpensePlanRecord, len(dbItems))
	for i, item := range dbItems {
		items[i] = ListExpensePlanRecordResponse_ListExpensePlanRecord{
			ExpensePlanRecordId: uint32(item.ExpensePlanRecordID),
			ExpensePlanId:       uint32(item.ExpensePlanID),
			AmountPaid:          uint32(item.AmountPaid),
			PaymentDate:         common.PgTimestamptzToISOString(&item.PaymentDate),
			PaidDate:            common.PgTimestamptzToISOString(&item.PaidDate),
			ExpensePlanSequence: uint32(item.ExpensePlanSequence),
			CreatedAt:           common.PgTimestamptzToISOString(&item.CreatedAt),
			UpdatedAt:           common.PgTimestamptzToISOString(&item.UpdatedAt),
			ExpensePlan: &ListExpensePlanRecordResponse_ListExpensePlan{
				ExpensePlanId:      uint32(item.ExpensePlanID),
				Title:              item.ExpensePlanTitle.String,
				Category:           string(item.ExpensePlanCategory.ExpensePlanCategory),
				AmountPlanned:      uint32(item.ExpensePlanAmountPlanned.Int32),
				RecurrencyInterval: uint32(item.ExpensePlanRecurrencyInterval.Int32),
				RecurrencyType:     nil,
			},
		}
		if item.ExpensePlanRecurrencyType.Valid {
			items[i].ExpensePlan.RecurrencyType = &item.ExpensePlanRecurrencyType.RecurrencyType
		}
	}

	json.NewEncoder(w).Encode(&ListExpensePlanRecordResponse{Items: &items})
}
