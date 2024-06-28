package cat

import (
	"sca/internal/context"
	domain "sca/internal/domain/cat"
	derr "sca/internal/error/data"

	"github.com/google/uuid"
)

type CatRepo interface {
	CreateCat(ctx *context.Context, create *domain.CreateCat) (*domain.Cat, derr.DataError)
	UpdateCat(ctx *context.Context, update *domain.UpdateCat) (*domain.Cat, derr.DataError)
	GetCat(ctx *context.Context, id uuid.UUID) (*domain.Cat, derr.DataError)
	ListCats(ctx *context.Context) ([]domain.Cat, derr.DataError)
	DeleteCat(ctx *context.Context, id uuid.UUID) derr.DataError
}
