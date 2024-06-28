package cat

import (
	"sca/internal/context"
	request "sca/internal/dto/request/cat"
	herr "sca/internal/error/http"
)

type CreateCat struct {
	Name        string
	YearsExp    uint
	Breed       Breed
	SalaryCents uint
}

func NewCreateCatFromRequest(
	ctx *context.Context,
	req *request.CreateCat,
) (*CreateCat, herr.ErrorHttp) {
	breed, err := NewBreed(ctx, req.MaybeBreed)
	if err != nil {
		return nil, err
	}

	return newCreateCat(req.Name, req.YearsExp, breed, req.SalaryCents), nil
}

func newCreateCat(
	name string,
	yearsExp uint,
	breed Breed,
	salaryCents uint,
) *CreateCat {
	return &CreateCat{
		Name:        name,
		YearsExp:    yearsExp,
		Breed:       breed,
		SalaryCents: salaryCents,
	}
}
