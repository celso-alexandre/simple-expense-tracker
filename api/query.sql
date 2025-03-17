-- name: ListExpensePlans :many
SELECT * FROM expense_plan;

-- name: GetExpensePlan :one
SELECT * FROM expense_plan WHERE expense_plan_id = sqlc.arg('expense_plan_id');

-- name: CreateExpensePlan :one
INSERT INTO expense_plan (
   title,
   category,
   amount_planned,
   recurrency_type,
   created_at,
   updated_at
) VALUES (
   sqlc.arg('title'),
   sqlc.narg('category'),
   sqlc.arg('amount_planned'),
   sqlc.arg('recurrency_type'),
   NOW(),
   NOW()
)
RETURNING *;

-- name: UpdateExpensePlan :one
UPDATE expense_plan SET
   title = sqlc.arg('title'),
   category = sqlc.narg('category'),
   amount_planned = sqlc.arg('amount_planned'),
   recurrency_type = sqlc.arg('recurrency_type'),
   updated_at = NOW()
WHERE expense_plan_id = sqlc.arg('expense_plan_id')
RETURNING *;

-- name: DeleteExpensePlan :execrows
DELETE FROM expense_plan WHERE expense_plan_id = sqlc.arg('expense_plan_id');

-- name: ListExpensePlanRecords :many
SELECT 
   rec.*,
   ep.title as expense_plan_title,
   ep.category as expense_plan_category,
   ep.amount_planned as expense_plan_amount_planned,
   ep.recurrency_type as expense_plan_recurrency_type,
   ep.recurrency_interval as expense_plan_recurrency_interval
FROM expense_plan_record rec
LEFT JOIN expense_plan ep ON rec.expense_plan_id = ep.expense_plan_id;
