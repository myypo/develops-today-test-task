package request

type CreateMission struct {
	Targets []CreateTarget `json:"targets" binding:"dive,required"`
}

type CreateTarget struct {
	Name    string `json:"name"    binding:"required"`
	Country string `json:"country" binding:"required"`
	Notes   string `json:"notes"   binding:"required"`
}
