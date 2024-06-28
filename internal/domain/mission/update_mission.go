package mission

import (
	"fmt"
	"sca/internal/context"
	"sca/internal/domain/common"
	reqComm "sca/internal/dto/request"
	reqMission "sca/internal/dto/request/mission"
	herr "sca/internal/error/http"

	"github.com/google/uuid"
)

type UpdateMission struct {
	Id uuid.UUID

	Status *common.Status
	CatId  *uuid.UUID
}

func NewUpdateMissionFromRequest(
	ctx *context.Context,
	sel *reqComm.InternalUUIDSelector,
	req *reqMission.UpdateMission,
) (*UpdateMission, herr.ErrorHttp) {
	id, err := uuid.Parse(sel.ID)
	if err != nil {
		return nil, herr.NewErrBadRequest(err)
	}
	catId, err := func() (*uuid.UUID, error) {
		if req.CatId == nil {
			return nil, nil
		}

		u, err := uuid.Parse(*req.CatId)
		return &u, err
	}()
	if err != nil {
		return nil, herr.NewErrBadRequest(err)
	}

	if req.Status != nil {
		status, ok := common.NewStatus(*req.Status)
		if !ok {
			return nil, herr.NewErrBadRequest(
				fmt.Errorf("inalid mission status provided: %v", req.Status),
			)
		}
		return newUpdateMission(id, &status, catId), nil
	}

	return newUpdateMission(id, nil, catId), nil
}

func newUpdateMission(
	id uuid.UUID,
	status *common.Status,
	catId *uuid.UUID,
) *UpdateMission {
	return &UpdateMission{
		Id:     id,
		Status: status,
		CatId:  catId,
	}
}
