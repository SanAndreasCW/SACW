package enums

type JobType uint32

const (
	Delivery JobType = 0
)

func (t JobType) String() string {
	switch t {
	case Delivery:
		return "Delivery"
	}
	return "unknown"
}
