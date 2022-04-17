package model

type ProjectState string

const (
	projectStatePlanned = "planned"
	projectStateActive  = "active"
	projectStateDone    = "done"
	projectStateFailed  = "failed"
)

func ValidateProjectState(str string) (ProjectState, bool) {
	switch str {
	case projectStatePlanned:
		return projectStatePlanned, true
	case projectStateActive:
		return projectStateActive, true
	case projectStateDone:
		return projectStateDone, true
	case projectStateFailed:
		return projectStateFailed, true
	case "":
		return "", true
	default:
		return ProjectState(str), false
	}
}
