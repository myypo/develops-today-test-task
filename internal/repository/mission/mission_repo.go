package cat

import (
	"sca/internal/context"
	domain "sca/internal/domain/mission"
	derr "sca/internal/error/data"

	"github.com/google/uuid"
)

type MissionRepo interface {
	CreateMission(
		ctx *context.Context,
		create *domain.CreateMission,
	) (*domain.Mission, derr.DataError)
	UpdateMission(
		ctx *context.Context,
		update *domain.UpdateMission,
	) (*domain.Mission, derr.DataError)
	GetMission(ctx *context.Context, id uuid.UUID) (*domain.Mission, derr.DataError)
	ListMissions(ctx *context.Context) ([]domain.Mission, derr.DataError)
	DeleteMission(ctx *context.Context, id uuid.UUID) derr.DataError
}
