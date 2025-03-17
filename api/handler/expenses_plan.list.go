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
	RecurrencyType     *query.RecurrencyType     `json:"recurrency_type"`
	RecurrencyInterval uint32                    `json:"recurrency_interval"`

	FirstExpensePlanRecordId uint32                                    `json:"first_expense_plan_record_id"`
	FirstExpensePlanRecord   *GetExpensePlanResponse_ExpensePlanRecord `json:"first_expense_plan_record"`
	LastExpensePlanRecordId  uint32                                    `json:"last_expense_plan_record_id"`
	LastExpensePlanRecord    *GetExpensePlanResponse_ExpensePlanRecord `json:"last_expense_plan_record"`

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
	for i, dbItem := range dbItems {
		items[i] = ListExpensePlanResponse_ListExpensePlan{
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
			items[i].RecurrencyType = &dbItem.RecurrencyType.RecurrencyType
		}
		if dbItem.FirstExpensePlanRecordID.Valid {
			items[i].FirstExpensePlanRecord = &GetExpensePlanResponse_ExpensePlanRecord{
				ExpensePlanRecordId: uint32(dbItem.FirstExpensePlanRecordID.Int32),
				AmountPaid:          uint32(dbItem.FirstAmountPaid.Int32),
				PaymentDate:         common.PgTimestamptzToISOString(&dbItem.FirstPaymentDate),
				PaidDate:            common.PgTimestamptzToISOString(&dbItem.FirstPaidDate),
				ExpensePlanSequence: uint32(dbItem.FirstExpensePlanSequence.Int32),
			}
		}
		if dbItem.LastExpensePlanRecordID.Valid {
			items[i].LastExpensePlanRecord = &GetExpensePlanResponse_ExpensePlanRecord{
				ExpensePlanRecordId: uint32(dbItem.LastExpensePlanRecordID.Int32),
				AmountPaid:          uint32(dbItem.LastAmountPaid.Int32),
				PaymentDate:         common.PgTimestamptzToISOString(&dbItem.LastPaymentDate),
				PaidDate:            common.PgTimestamptzToISOString(&dbItem.LastPaidDate),
				ExpensePlanSequence: uint32(dbItem.LastExpensePlanSequence.Int32),
			}
		}
	}

	json.NewEncoder(w).Encode(&ListExpensePlanResponse{Items: &items})
}
