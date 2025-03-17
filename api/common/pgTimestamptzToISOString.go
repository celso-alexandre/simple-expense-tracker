package common

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func PgTimestamptzToISOString(t *pgtype.Timestamptz) string {
	if !t.Valid {
		return ""
	}
	return t.Time.UTC().Format(time.RFC3339)
}

func ISOStringToPgTimestamptz(s string) *pgtype.Timestamptz {
	if s == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		fmt.Println("ISOStringToPgTimestamptz time.Parse err", err)
		return &pgtype.Timestamptz{}
	}
	return &pgtype.Timestamptz{Time: t, Valid: true}
}
