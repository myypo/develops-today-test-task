package request

type CreateCat struct {
	Name        string `json:"name"                binding:"required"`
	YearsExp    uint   `json:"years_of_experience" binding:"required"`
	MaybeBreed  string `json:"breed"               binding:"required"`
	SalaryCents uint   `json:"salary_in_cents"     binding:"required"`
}
