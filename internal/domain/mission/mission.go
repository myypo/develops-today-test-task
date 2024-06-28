package mission

import (
	"sca/internal/domain/common"
	"sca/internal/domain/target"

	"github.com/google/uuid"
)

type Mission struct {
	Id uuid.UUID

	Status common.Status
	CatId  *string

	Targets []target.Target
}
