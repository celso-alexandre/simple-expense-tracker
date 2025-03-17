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
