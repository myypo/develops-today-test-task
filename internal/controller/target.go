package controller

import (
	"errors"
	"fmt"
	"sca/internal/context"
	"sca/internal/domain/common"
	domain "sca/internal/domain/target"
	commReq "sca/internal/dto/request"
	request "sca/internal/dto/request/target"
	response "sca/internal/dto/response/target"
	herr "sca/internal/error/http"
	misRepo "sca/internal/repository/mission"
	targRepo "sca/internal/repository/target"

	"github.com/google/uuid"
)

type TargetController struct {
	targRepo targRepo.TargetRepo
	misRepo  misRepo.MissionRepo
}

func NewTargetController(
	targRepo targRepo.TargetRepo,
	misRepo misRepo.MissionRepo,
) TargetController {
	return TargetController{
		targRepo,
		misRepo,
	}
}

func (c *TargetController) AddTarget(
	ctx *context.Context,
	req *request.AddTarget,
) (*response.Target, herr.ErrorHttp) {
	misId, err := uuid.Parse(req.MissionId)
	if err != nil {
		return nil, herr.NewErrBadRequest(err)
	}
	misDom, derr := c.misRepo.GetMission(ctx, misId)
	if derr != nil {
		if derr.NotFound() {
			return nil, herr.NewErrNotFound(derr)
		}
		return nil, herr.NewErrInternal(derr)
	}
	if len(misDom.Targets) >= 3 {
		return nil, herr.NewErrBadRequest(
			fmt.Errorf(
				"a mission can have at most 3 targets, but you are trying to add the %v one",
				len(misDom.Targets)+1,
			),
		)
	}
	if misDom.Status == common.Completed {
		return nil, herr.NewErrBadRequest(
			errors.New("tried to add a target to an already completed mission"),
		)
	}

	dom := domain.NewAddTargetFromRequest(req)
	targDom, derr := c.targRepo.AddTarget(ctx, dom)
	if derr != nil {
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewTargetFromDomain(targDom), nil
}

func (c *TargetController) UpdateTarget(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
	req *request.UpdateTarget,
) (*response.Target, herr.ErrorHttp) {
	selId, err := uuid.Parse(sel.ID)
	if err != nil {
		return nil, herr.NewErrBadRequest(err)
	}
	targDom, derr := c.targRepo.GetTarget(ctx, selId)
	if derr != nil {
		if derr.NotFound() {
			return nil, herr.NewErrNotFound(derr)
		}
		return nil, herr.NewErrInternal(derr)
	}
	if targDom.Status == common.Completed {
		return nil, herr.NewErrBadRequest(
			errors.New("the requested target to be updated is already completed"),
		)
	}

	misDom, derr := c.misRepo.GetMission(ctx, targDom.Id)
	if derr != nil {
		return nil, herr.NewErrInternal(derr)
	}
	if misDom.Status == common.Completed {
		return nil, herr.NewErrBadRequest(
			errors.New(
				"the requested mission of the target to be updated has already been completed",
			),
		)
	}

	dom := domain.NewUpdateTargetFromRequest(sel, req)
	targDom, derr = c.targRepo.UpdateTarget(ctx, &dom)
	if derr != nil {
		if derr.NotFound() {
			return nil, herr.NewErrNotFound(derr)
		}
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewTargetFromDomain(targDom), nil
}

func (c *TargetController) DeleteTarget(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
) herr.ErrorHttp {
	targId, err := uuid.Parse(sel.ID)
	if err != nil {
		return herr.NewErrBadRequest(err)
	}

	targDom, derr := c.targRepo.GetTarget(ctx, targId)
	if derr != nil {
		if derr.NotFound() {
			return herr.NewErrNotFound(derr)
		}
		return herr.NewErrInternal(derr)
	}
	if targDom.Status == common.Completed {
		return herr.NewErrBadRequest(
			errors.New("the target requested for deletion has already been completed"),
		)
	}

	misDom, derr := c.misRepo.GetMission(ctx, targDom.MissionId)
	if derr != nil {
		return herr.NewErrInternal(derr)
	}
	if misDom.Status == common.Completed {
		return herr.NewErrBadRequest(
			errors.New(
				"the mission of the target requested for deleteion has already been completed",
			),
		)
	}

	derr = c.targRepo.DeleteTarget(ctx, targId)
	if derr != nil {
		if derr.NotFound() {
			return herr.NewErrNotFound(derr)
		}
		return herr.NewErrInternal(derr)
	}

	return nil
}
