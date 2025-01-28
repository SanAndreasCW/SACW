package commons

type Refuse struct {
	error
}

func NewRefuse() *Refuse {
	return &Refuse{}
}
