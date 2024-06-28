package target

import (
	"sca/internal/domain/common"
	reqComm "sca/internal/dto/request"
	reqTarg "sca/internal/dto/request/target"
)

type UpdateTarget struct {
	Id string

	Name    *string
	Country *string
	Notes   *string
	Status  *common.Status
}

func NewUpdateTargetFromRequest(
	sel *reqComm.InternalUUIDSelector,
	req *reqTarg.UpdateTarget,
) UpdateTarget {
	return UpdateTarget{
		Id: sel.ID,

		Name:    req.Name,
		Country: req.Country,
		Notes:   req.Notes,
		Status:  (*common.Status)(req.Status),
	}
}
