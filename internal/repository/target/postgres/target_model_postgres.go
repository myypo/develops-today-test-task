package postgres

import (
	domComm "sca/internal/domain/common"
	domTarg "sca/internal/domain/target"
	"sca/internal/util"

	model "sca/.gen/jet/sca/public/model"

	"github.com/google/uuid"
)

func NewDomainFromModel(mod *model.Targets) *domTarg.Target {
	return &domTarg.Target{
		Id: mod.ID,

		Name:    mod.Name,
		Country: mod.Country,
		Notes:   mod.Notes,
		Status:  domComm.Status(mod.Status),

		MissionId: mod.MissionID,
	}
}

func newModelFromAdd(add *domTarg.AddTarget) (model.Targets, error) {
	u, err := uuid.FromBytes([]byte(add.MissionId))
	if err != nil {
		return model.Targets{}, err
	}

	return model.Targets{
		MissionID: u,
		Name:      add.Name,
		Country:   add.Country,
		Notes:     add.Notes,

		Status: model.TargetStatus(domComm.InProgress),
	}, nil
}

func newModelFromUpdate(update *domTarg.UpdateTarget) model.Targets {
	return model.Targets{
		Name:    util.DerefOrDefault(update.Name),
		Country: util.DerefOrDefault(update.Country),
		Notes:   util.DerefOrDefault(update.Notes),
		Status:  model.TargetStatus(util.DerefOrDefault(update.Status)),
	}
}
