//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
)

type Cats struct {
	ID                uuid.UUID `sql:"primary_key"`
	Name              *string
	YearsOfExperience *int32
	Breed             *string
	SalaryInCents     *int64
}
