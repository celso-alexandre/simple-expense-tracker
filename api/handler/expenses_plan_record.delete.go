package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
)

type DeleteExpensePlanRecordsRequest struct {
	ExpensePlanRecordsId uint32 `json:"expense_plan_record_id" validate:"required"`
}

type DeleteExpensePlanRecordsResponse struct{}

// DeleteExpensePlanRecord godoc
// @Router       /expense-plan-record/delete [post]
// @Accept       json
// @Param        request body DeleteExpensePlanRecordsRequest false "DeleteExpensePlanRecordsRequest"
// @Produce      json
// @Success      200   {object}  DeleteExpensePlanRecordsResponse
// @Summary      Delete expense-plan-record item
// @Description  Delete expense-plan-record item
// @Tags         ExpensePlanRecords
func DeleteExpensePlanRecord(w http.ResponseWriter, r *http.Request) {
	var req DeleteExpensePlanRecordsRequest
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

	rowcount, err := q.DeleteExpensePlanRecord(ctx, int32(req.ExpensePlanRecordsId))
	if err != nil {
		fmt.Println("rowcount: ", rowcount, "id: ", req.ExpensePlanRecordsId)
		msg := fmt.Sprintf("failed to Delete expense plan record: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if rowcount != 1 {
		msg := fmt.Sprintf("failed to Delete expense plan record rowcount: %d", rowcount)
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

	json.NewEncoder(w).Encode(&DeleteExpensePlanRecordsResponse{})
}
