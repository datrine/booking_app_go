// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Availability struct {
	Weekday   int16
	StartTime pgtype.Time
	EndTime   pgtype.Time
}

type Booking struct {
	ID        int64
	StartTime pgtype.Timestamp
	EndTime   pgtype.Timestamp
	Email     string
	CreatedAt pgtype.Timestamp
}
