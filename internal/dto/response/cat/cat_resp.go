package response

import (
	domain "sca/internal/domain/cat"
)

type Cat struct {
	Id string `json:"id"`

	Name        string       `json:"name"`
	YearsExp    uint         `json:"years_of_experience"`
	Breed       domain.Breed `json:"breed"`
	SalaryCents uint         `json:"salary_in_cents"`
}

func NewCatFromDomain(domCat *domain.Cat) *Cat {
	return &Cat{
		Id: domCat.Id.String(),

		Name:        domCat.Name,
		YearsExp:    domCat.YearsExp,
		Breed:       domCat.Breed,
		SalaryCents: domCat.SalaryCents,
	}
}
