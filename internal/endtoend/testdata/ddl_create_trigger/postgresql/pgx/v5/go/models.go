// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.1

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type TddTest struct {
	TestID    pgtype.UUID
	Title     string
	Descr     string
	TsCreated pgtype.Timestamptz
	TsUpdated pgtype.Timestamptz
}