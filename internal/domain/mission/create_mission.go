package mission

import (
	"fmt"
	request "sca/internal/dto/request/mission"
	herr "sca/internal/error/http"
	"sca/internal/util"
)

type CreateMission struct {
	Targets []CreateTarget
}

type CreateTarget struct {
	Name    string
	Country string
	Notes   string
}

func NewCreateMissionFromRequest(
	req *request.CreateMission,
) (*CreateMission, herr.ErrorHttp) {
	return newCreateMission(req.Targets)
}

func newCreateMission(
	targets []request.CreateTarget,
) (*CreateMission, herr.ErrorHttp) {
	lenTargs := len(targets)
	if lenTargs > 3 || lenTargs < 1 {
		return nil, herr.NewErrBadRequest(
			fmt.Errorf(
				"a mission can have from 1 upto 3 targets, but %v were provided instead",
				lenTargs,
			),
		)
	}

	return &CreateMission{Targets: util.Map(targets, newCreateTargetFromRequest)}, nil
}

func newCreateTargetFromRequest(
	req request.CreateTarget,
) CreateTarget {
	return CreateTarget{
		Name:    req.Name,
		Country: req.Country,
		Notes:   req.Notes,
	}
}
