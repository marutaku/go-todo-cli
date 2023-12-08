package todo

type Status = string

const (
	Todo       = Status("Todo")
	InProgress = Status("InProgress")
	Done       = Status("Done")
)

func ValidateStatus(status Status) bool {
	switch status {
	case Todo, InProgress, Done:
		return true
	}
	return false
}
