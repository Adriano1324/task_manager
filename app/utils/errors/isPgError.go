package pgErrorComparison

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func IsDuplicatedKeyError(err error) bool {
	var perr *pgconn.PgError
	if errors.As(err, &perr) {
		return perr.Code == "23505"
	}
	return false
}
