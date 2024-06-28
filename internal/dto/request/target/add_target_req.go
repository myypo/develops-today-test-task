package request

type AddTarget struct {
	MissionId string `json:"mission_id" binding:"required,uuid"`

	Name    string `json:"name"    binding:"required"`
	Country string `json:"country" binding:"required"`
	Notes   string `json:"notes"   binding:"required"`
}
