package request

type UpdateCat struct {
	SalaryCents uint `json:"salary_in_cents" binding:"required"`
}
