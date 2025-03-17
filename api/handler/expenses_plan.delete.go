package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celso-alexandre/api/common"
	"github.com/celso-alexandre/api/database"
)

type DeleteExpensePlanRequest struct {
	ExpensePlanId uint32 `json:"expense_plan_id" validate:"required"`
}

type DeleteExpensePlanResponse struct{}

// DeleteExpensePlan godoc
// @Router       /expense-plan/delete [post]
// @Accept       json
// @Param        request body DeleteExpensePlanRequest false "DeleteExpensePlanRequest"
// @Produce      json
// @Success      200   {object}  DeleteExpensePlanResponse
// @Summary      Delete expense-plan item
// @Description  Delete expense-plan item
// @Tags         ExpensePlan
func DeleteExpensePlan(w http.ResponseWriter, r *http.Request) {
	var req DeleteExpensePlanRequest
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

	rowcount, err := q.DeleteExpensePlan(ctx, int32(req.ExpensePlanId))
	if err != nil {
		fmt.Println("rowcount: ", rowcount, "id: ", req.ExpensePlanId)
		msg := fmt.Sprintf("failed to Delete expense plan: %v", err)
		fmt.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if rowcount != 1 {
		msg := fmt.Sprintf("failed to Delete expense plan rowcount: %d", rowcount)
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

	json.NewEncoder(w).Encode(&DeleteExpensePlanResponse{})
}
