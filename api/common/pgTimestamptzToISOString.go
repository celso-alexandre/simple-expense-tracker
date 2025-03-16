package common

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func PgTimestamptzToISOString(t *pgtype.Timestamptz) string {
	if !t.Valid {
		return ""
	}
	return t.Time.UTC().Format(time.RFC3339)
}
