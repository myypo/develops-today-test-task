package target

import (
	reqTarg "sca/internal/dto/request/target"
)

type AddTarget struct {
	MissionId string

	Name    string
	Country string
	Notes   string
}

func NewAddTargetFromRequest(
	req *reqTarg.AddTarget,
) *AddTarget {
	return &AddTarget{
		MissionId: req.MissionId,
		Name:      req.Name,
		Country:   req.Country,
		Notes:     req.Notes,
	}
}
