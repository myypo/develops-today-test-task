package response

type Base struct {
	Data  any    `json:"data"`
	Error string `json:"error,omitempty"`
}

func NewBaseResponse(data any, err string) Base {
	return Base{
		Data:  data,
		Error: err,
	}
}
