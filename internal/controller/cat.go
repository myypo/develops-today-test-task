package controller

import (
	"sca/internal/context"
	domain "sca/internal/domain/cat"
	commReq "sca/internal/dto/request"
	request "sca/internal/dto/request/cat"
	response "sca/internal/dto/response/cat"
	herr "sca/internal/error/http"
	catRepo "sca/internal/repository/cat"
	"sca/internal/util"

	"github.com/google/uuid"
)

type CatController struct {
	catRepo catRepo.CatRepo
}

func NewCatController(catRepo catRepo.CatRepo) CatController {
	return CatController{
		catRepo,
	}
}

func (c *CatController) CreateCat(
	ctx *context.Context,
	req *request.CreateCat,
) (*response.Cat, herr.ErrorHttp) {
	dom, err := domain.NewCreateCatFromRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	domCat, derr := c.catRepo.CreateCat(ctx, dom)
	if derr != nil {
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewCatFromDomain(domCat), nil
}

func (c *CatController) UpdateCat(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
	req *request.UpdateCat,
) (*response.Cat, herr.ErrorHttp) {
	dom, err := domain.NewUpdateCatFromRequest(ctx, sel, req)
	if err != nil {
		return nil, err
	}

	domCat, derr := c.catRepo.UpdateCat(ctx, dom)
	if derr != nil {
		if derr.NotFound() {
			return nil, herr.NewErrNotFound(derr)
		}
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewCatFromDomain(domCat), nil
}

func (c *CatController) GetCat(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
) (*response.Cat, herr.ErrorHttp) {
	uuid, err := uuid.Parse(sel.ID)
	if err != nil {
		return nil, herr.NewErrBadRequest(err)
	}
	domCat, derr := c.catRepo.GetCat(ctx, uuid)
	if derr != nil {
		if derr.NotFound() {
			return nil, herr.NewErrNotFound(derr)
		}
		return nil, herr.NewErrInternal(derr)
	}

	return response.NewCatFromDomain(domCat), nil
}

func (c *CatController) ListCats(
	ctx *context.Context,
) ([]response.Cat, herr.ErrorHttp) {
	domCatList, derr := c.catRepo.ListCats(ctx)
	if derr != nil {
		return nil, herr.NewErrInternal(derr)
	}

	return util.MapRef(domCatList, response.NewCatFromDomain), nil
}

func (c *CatController) DeleteCat(
	ctx *context.Context,
	sel *commReq.InternalUUIDSelector,
) herr.ErrorHttp {
	uuid, err := uuid.Parse(sel.ID)
	if err != nil {
		return herr.NewErrBadRequest(err)
	}
	derr := c.catRepo.DeleteCat(ctx, uuid)
	if derr != nil {
		return herr.NewErrInternal(derr)
	}

	return nil
}
