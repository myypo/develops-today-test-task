package jet

import (
	"errors"
	derr "sca/internal/error/data"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/lib/pq"
)

func ErrSpec(jetErr error) (err derr.DataErrorType, constaint string) {
	if errors.Is(jetErr, qrm.ErrNoRows) {
		return derr.NotFound, ""
	}

	var pqErr *pq.Error
	if errors.As(jetErr, &pqErr) {
		switch pqErr.Code {
		case uniqueViolationCode:
			return derr.Conflict, pqErr.Constraint
		case doesNotExistCode:
			return derr.NotFound, ""
		default:
			return derr.Internal, ""
		}
	}
	return derr.Internal, ""
}

const (
	uniqueViolationCode = "23505"
	doesNotExistCode    = "42703"
)
