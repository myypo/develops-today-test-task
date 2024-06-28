package common

type Status string

const (
	InProgress Status = "IN_PROGRESS"
	Completed  Status = "COMPLETED"
)

func NewStatus(maybeStatus string) (Status, bool) {
	switch maybeStatus {
	case string(InProgress):
		{
			return InProgress, true
		}
	case string(Completed):
		{
			return Completed, true
		}
	}

	return "", false
}
