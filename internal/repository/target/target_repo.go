package cat

import (
	"sca/internal/context"
	domain "sca/internal/domain/target"
	derr "sca/internal/error/data"

	"github.com/google/uuid"
)

type TargetRepo interface {
	GetTarget(
		ctx *context.Context,
		id uuid.UUID,
	) (*domain.Target, derr.DataError)
	AddTarget(
		ctx *context.Context,
		add *domain.AddTarget,
	) (*domain.Target, derr.DataError)
	UpdateTarget(
		ctx *context.Context,
		update *domain.UpdateTarget,
	) (*domain.Target, derr.DataError)
	DeleteTarget(ctx *context.Context, id uuid.UUID) derr.DataError
}
