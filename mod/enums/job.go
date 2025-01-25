package enums

type JobType int32

const (
	Delivery JobType = 0
)

func (t JobType) String() string {
	switch t {
	case Delivery:
		return "delivery"
	}
	return "unknown"
}
