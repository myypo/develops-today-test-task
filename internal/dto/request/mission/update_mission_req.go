package request

type UpdateMission struct {
	Status *string `json:"status" binding:"omitempty,oneof=COMPLETED"`
	CatId  *string `json:"cat_id"`
}
