package postgres

import (
	domCat "sca/internal/domain/cat"

	model "sca/.gen/jet/sca/public/model"
)

func newModelFromCreate(create *domCat.CreateCat) model.Cats {
	return model.Cats{
		Name: &create.Name,
		YearsOfExperience: func() *int32 {
			i := int32(create.YearsExp)
			return &i
		}(),
		Breed: (*string)(&create.Breed),
		SalaryInCents: func() *int64 {
			i := int64(create.SalaryCents)
			return &i
		}(),
	}
}

func newModelFromUpdate(update *domCat.UpdateCat) model.Cats {
	return model.Cats{
		SalaryInCents: func() *int64 {
			i := int64(update.SalaryCents)
			return &i
		}(),
	}
}

func newDomainFromModel(modCat *model.Cats) *domCat.Cat {
	return &domCat.Cat{
		Id:          modCat.ID,
		Name:        *modCat.Name,
		YearsExp:    uint(*modCat.YearsOfExperience),
		Breed:       domCat.Breed(*modCat.Breed),
		SalaryCents: uint(*modCat.SalaryInCents),
	}
}
