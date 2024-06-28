package cat

import (
	"sca/internal/context"
	reqComm "sca/internal/dto/request"
	reqCat "sca/internal/dto/request/cat"
	herr "sca/internal/error/http"

	"github.com/google/uuid"
)

type UpdateCat struct {
	Id          uuid.UUID
	SalaryCents uint
}

func NewUpdateCatFromRequest(
	ctx *context.Context,
	sel *reqComm.InternalUUIDSelector,
	req *reqCat.UpdateCat,
) (*UpdateCat, herr.ErrorHttp) {
	uuid, err := uuid.Parse(sel.ID)
	if err != nil {
		return nil, herr.NewErrBadRequest(err)
	}

	return newUpdateCat(uuid, req.SalaryCents), nil
}

func newUpdateCat(
	id uuid.UUID,
	salaryCents uint,
) *UpdateCat {
	return &UpdateCat{
		Id:          id,
		SalaryCents: salaryCents,
	}
}
