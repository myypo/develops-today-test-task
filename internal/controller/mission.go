package controller

import (
	"errors"
	"sca/internal/context"
	"sca/internal/domain/common"
	domain "sca/internal/domain/mission"
	commReq "sca/internal/dto/request"
	request "sca/internal/dto/request/mission"
	response "sca/internal/dto/response/mission"
	herr "sca/internal/error/http"
	misRepo "sca/internal/repository/mission"
	"sca/internal/util"

	"github.com/google/uuid"
)

type MissionController struct {
	misRepo misRepo.MissionRepo
}

func NewMissionController(
	misRepo misRepo.MissionRepo,
) MissionController {
	return MissionController{
		misRepo,
	}
}

func (c *MissionController) CreateMission(
	ctx *context.Context,
	req *request.CreateMission,
) (*response.Mission, herr.ErrorHttp) {
	dom, err := domain.NewCreateMissionFromRequest(req)
	if err != nil {
		return nil, err
	}

	misDom, derr := c.misRepo.CreateMission(ctx, dom)
	if derr != nil {
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewMissionFromDomain(misDom), nil
}

func (c *MissionController) UpdateMission(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
	req *request.UpdateMission,
) (*response.Mission, herr.ErrorHttp) {
	uuid, uerr := uuid.Parse(sel.ID)
	if uerr != nil {
		return nil, herr.NewErrBadRequest(uerr)
	}
	domMis, derr := c.misRepo.GetMission(ctx, uuid)
	if derr != nil {
		if derr.NotFound() {
			return nil, herr.NewErrNotFound(derr)
		}
		return nil, herr.NewErrInternal(derr)
	}
	if domMis.Status == common.Completed {
		return nil, herr.NewErrBadRequest(
			errors.New("the requested mission to be updated is already completed"),
		)
	}

	dom, err := domain.NewUpdateMissionFromRequest(ctx, sel, req)
	if err != nil {
		return nil, err
	}
	domMis, derr = c.misRepo.UpdateMission(ctx, dom)
	if derr != nil {
		if derr.NotFound() {
			return nil, herr.NewErrNotFound(derr)
		}
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewMissionFromDomain(domMis), nil
}

func (c *MissionController) GetMission(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
) (*response.Mission, herr.ErrorHttp) {
	uuid, uerr := uuid.Parse(sel.ID)
	if uerr != nil {
		return nil, herr.NewErrBadRequest(uerr)
	}

	domMis, derr := c.misRepo.GetMission(ctx, uuid)
	if derr != nil {
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewMissionFromDomain(domMis), nil
}

func (c *MissionController) ListMissions(
	ctx *context.Context,
) ([]response.Mission, herr.ErrorHttp) {
	domMissionList, derr := c.misRepo.ListMissions(ctx)
	if derr != nil {
		return nil, herr.NewErrInternal(derr)
	}

	return util.MapRef(domMissionList, response.NewMissionFromDomain), nil
}

func (c *MissionController) DeleteMission(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
) herr.ErrorHttp {
	uuid, uerr := uuid.Parse(sel.ID)
	if uerr != nil {
		return herr.NewErrBadRequest(uerr)
	}
	domMis, derr := c.misRepo.GetMission(ctx, uuid)
	if derr != nil {
		if derr.NotFound() {
			return herr.NewErrNotFound(derr)
		}
		return herr.NewErrInternal(derr)
	}
	if domMis.CatId != nil {
		return herr.NewErrBadRequest(
			errors.New("the requested mission to be deleted has already been assigned to a cat"),
		)
	}

	derr = c.misRepo.DeleteMission(ctx, uuid)
	if derr != nil {
		return herr.NewErrInternal(derr)
	}

	return nil
}
