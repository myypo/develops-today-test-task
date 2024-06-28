package request

type UpdateTarget struct {
	Name    *string `json:"name"`
	Country *string `json:"country"`
	Notes   *string `json:"notes"`
	Status  *string `json:"status"  binding:"omitempty,oneof=COMPLETED"`
}
