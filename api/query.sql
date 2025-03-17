-- name: ListExpensePlans :many
SELECT 
   pla.*,
   frec.amount_paid as first_amount_paid,
   frec.payment_date as first_payment_date,
   frec.paid_date as first_paid_date,
   frec.expense_plan_sequence as first_expense_plan_sequence,
   lrec.amount_paid as last_amount_paid,
   lrec.payment_date as last_payment_date,
   lrec.paid_date as last_paid_date,
   lrec.expense_plan_sequence as last_expense_plan_sequence
FROM expense_plan pla
LEFT JOIN expense_plan_record frec ON pla.first_expense_plan_record_id = frec.expense_plan_record_id
LEFT JOIN expense_plan_record lrec ON pla.last_expense_plan_record_id = lrec.expense_plan_record_id;

-- name: GetExpensePlan :one
SELECT 
   pla.*,
   frec.amount_paid as first_amount_paid,
   frec.payment_date as first_payment_date,
   frec.paid_date as first_paid_date,
   frec.expense_plan_sequence as first_expense_plan_sequence,
   lrec.amount_paid as last_amount_paid,
   lrec.payment_date as last_payment_date,
   lrec.paid_date as last_paid_date,
   lrec.expense_plan_sequence as last_expense_plan_sequence
FROM expense_plan pla
LEFT JOIN expense_plan_record frec ON pla.first_expense_plan_record_id = frec.expense_plan_record_id
LEFT JOIN expense_plan_record lrec ON pla.last_expense_plan_record_id = lrec.expense_plan_record_id
WHERE pla.expense_plan_id = sqlc.arg('expense_plan_id');

-- name: CreateExpensePlan :one
INSERT INTO expense_plan (
   title,
   category,
   amount_planned,
   recurrency_type,
   recurrency_interval,
   created_at,
   updated_at
) VALUES (
   sqlc.arg('title'),
   sqlc.narg('category'),
   sqlc.arg('amount_planned'),
   sqlc.arg('recurrency_type'),
   sqlc.arg('recurrency_interval'),
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
   recurrency_interval = sqlc.arg('recurrency_interval'),
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

-- name: GetExpensePlanRecord :one
SELECT 
   rec.*,
   ep.title as expense_plan_title,
   ep.category as expense_plan_category,
   ep.amount_planned as expense_plan_amount_planned,
   ep.recurrency_type as expense_plan_recurrency_type,
   ep.recurrency_interval as expense_plan_recurrency_interval
FROM expense_plan_record rec
LEFT JOIN expense_plan ep ON rec.expense_plan_id = ep.expense_plan_id
WHERE rec.expense_plan_record_id = sqlc.arg('expense_plan_record_id');

-- name: CreateExpensePlanRecord :one
INSERT INTO expense_plan_record (
   expense_plan_id,
   amount_paid,
   payment_date,
   paid_date,
   expense_plan_sequence,
   created_at,
   updated_at
) VALUES (
   sqlc.arg('expense_plan_id'),
   sqlc.arg('amount_paid'),
   sqlc.arg('payment_date'),
   sqlc.arg('paid_date'),
   COALESCE((
      SELECT r.expense_plan_sequence
      FROM expense_plan_record r
      WHERE r.expense_plan_id = sqlc.arg('expense_plan_id')
      AND r.payment_date      < sqlc.arg('payment_date')
      ORDER BY r.payment_date DESC
      LIMIT 1
   ), 0) + 1,
   NOW(),
   NOW()
)
RETURNING *;

-- name: UpdateExpensePlanAfterRecord :one
UPDATE expense_plan SET
   first_expense_plan_record_id = COALESCE((
      SELECT r.expense_plan_record_id
      FROM expense_plan_record r
      WHERE r.expense_plan_id = sqlc.arg('expense_plan_id')
      ORDER BY r.payment_date ASC
      LIMIT 1
   ), NULL),
   last_expense_plan_record_id = COALESCE((
      SELECT r.expense_plan_record_id
      FROM expense_plan_record r
      WHERE r.expense_plan_id = sqlc.arg('expense_plan_id')
      ORDER BY r.payment_date DESC
      LIMIT 1
   ), NULL),
   updated_at = NOW()
WHERE expense_plan_id = sqlc.arg('expense_plan_id')
RETURNING *;

-- name: UpdateExpensePlanRecord :one
UPDATE expense_plan_record SET
   expense_plan_id = sqlc.arg('expense_plan_id'),
   amount_paid = sqlc.arg('amount_paid'),
   payment_date = sqlc.arg('payment_date'),
   paid_date = sqlc.arg('paid_date'),
   expense_plan_sequence = COALESCE((
      SELECT r.expense_plan_sequence
      FROM expense_plan_record r
      WHERE r.expense_plan_id = sqlc.arg('expense_plan_id')
      AND r.payment_date      < sqlc.arg('payment_date')
      ORDER BY r.payment_date DESC
      LIMIT 1
   ), 0) + 1,
   updated_at = NOW()
WHERE expense_plan_record_id = sqlc.arg('expense_plan_record_id')
RETURNING *;

-- name: DeleteExpensePlanRecord :execrows
DELETE FROM expense_plan_record WHERE expense_plan_record_id = sqlc.arg('expense_plan_record_id');
