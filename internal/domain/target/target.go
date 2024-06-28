package target

import (
	"sca/internal/domain/common"

	"github.com/google/uuid"
)

type Target struct {
	Id uuid.UUID

	Name    string
	Country string
	Notes   string
	Status  common.Status

	MissionId uuid.UUID
}
