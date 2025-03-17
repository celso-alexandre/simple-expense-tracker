package common

import "github.com/jackc/pgx/v5/pgtype"

func Uint32ToPgInt(v uint32) *pgtype.Int4 {
	return &pgtype.Int4{Int32: int32(v), Valid: v > 0}
}
