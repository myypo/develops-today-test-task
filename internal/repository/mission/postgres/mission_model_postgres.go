package postgres

import (
	domComm "sca/internal/domain/common"
	domMis "sca/internal/domain/mission"
	domTarg "sca/internal/domain/target"
	targRepo "sca/internal/repository/target/postgres"
	"sca/internal/util"

	model "sca/.gen/jet/sca/public/model"

	"github.com/google/uuid"
)

func newModelFromCreate(create *domMis.CreateMission) model.Missions {
	return model.Missions{
		Status: model.MissionStatus_InProgress,
	}
}

func newTargListFromCreate(misId uuid.UUID, create []domMis.CreateTarget) []model.Targets {
	toMod := func(dom domMis.CreateTarget) model.Targets {
		return model.Targets{
			MissionID: misId,

			Name:    dom.Name,
			Country: dom.Country,
			Notes:   dom.Notes,
			Status:  model.TargetStatus(model.MissionStatus_InProgress),
		}
	}
	return util.Map(create, toMod)
}

func newTargFromUpdate(update domTarg.UpdateTarget) model.Targets {
	mod := model.Targets{}
	if update.Name != nil {
		mod.Name = *update.Name
	}
	if update.Country != nil {
		mod.Country = *update.Country
	}
	if update.Notes != nil {
		mod.Notes = *update.Notes
	}
	if update.Status != nil {
		mod.Status = model.TargetStatus(*update.Status)
	}

	return mod
}

func newModelFromUpdate(update *domMis.UpdateMission) (model.Missions, error) {
	return model.Missions{
		Status: func() model.MissionStatus {
			if update.Status == nil {
				return ""
			} else {
				return model.MissionStatus(*update.Status)
			}
		}(),
		CatID: update.CatId,
	}, nil
}

func newDomainFromModel(modMis *SaturatedMission) *domMis.Mission {
	return &domMis.Mission{
		Id:     modMis.ID,
		Status: domComm.Status(modMis.Status),
		CatId: func() *string {
			if modMis.CatID == nil {
				return nil
			} else {
				s := modMis.CatID.String()
				return &s
			}
		}(),
		Targets: util.MapRef(modMis.Targets, targRepo.NewDomainFromModel),
	}
}

type SaturatedMission struct {
	model.Missions

	Targets []model.Targets
}
