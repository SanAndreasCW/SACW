package enums

import "strings"

type JobGroup int16
type CheckpointType int16
type JobType int32

const (
	CheckpointFoot JobGroup = iota
	CheckpointVehicle
	Free
)

const (
	Target CheckpointType = iota
	Lookup
)

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
