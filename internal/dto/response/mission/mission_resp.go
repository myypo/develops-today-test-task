package response

import (
	domMis "sca/internal/domain/mission"
	target "sca/internal/dto/response/target"
	"sca/internal/util"
)

type Mission struct {
	Id string `json:"id"`

	Status string  `json:"status"`
	CatId  *string `json:"cat_id"`

	Targets []target.Target `json:"targets"`
}

func NewMissionFromDomain(domMis *domMis.Mission) *Mission {
	return &Mission{
		Id: domMis.Id.String(),

		Status: string(domMis.Status),
		CatId:  domMis.CatId,

		Targets: util.MapRef(domMis.Targets, target.NewTargetFromDomain),
	}
}
