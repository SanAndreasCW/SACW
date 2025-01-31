package enums

import "strings"

type JobType int32

const (
	Unknown JobType = iota
	Delivery
)

func GetJobType(s string) JobType {
	s = strings.ToLower(strings.TrimSpace(s))
	switch s {
	case "delivery":
		return Delivery
	}
	return Unknown
}

func (t JobType) String() string {
	switch t {
	case Delivery:
		return "Delivery"
	default:
		return "Unknown"
	}
}
