package response

import (
	domTarg "sca/internal/domain/target"
)

type Target struct {
	Id string `json:"id"`

	Name    string `json:"name"`
	Country string `json:"country"`
	Notes   string `json:"notes"`
	Status  string `json:"status"`
}

func NewTargetFromDomain(domTarg *domTarg.Target) *Target {
	return &Target{
		Id: domTarg.Id.String(),

		Name:    domTarg.Name,
		Country: domTarg.Country,
		Notes:   domTarg.Notes,
		Status:  string(domTarg.Status),
	}
}
