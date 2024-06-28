//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type MissionStatus string

const (
	MissionStatus_InProgress MissionStatus = "IN_PROGRESS"
	MissionStatus_Completed  MissionStatus = "COMPLETED"
)

func (e *MissionStatus) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "IN_PROGRESS":
		*e = MissionStatus_InProgress
	case "COMPLETED":
		*e = MissionStatus_Completed
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for MissionStatus enum")
	}

	return nil
}

func (e MissionStatus) String() string {
	return string(e)
}
