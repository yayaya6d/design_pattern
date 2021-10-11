package adapter

var USPowerVal float64 = 220.0

type USSocket interface {
	OutputUSPower() float64
}

func NewUSSocket() USSocket {
	return &usSocket{}
}

type usSocket struct {
}

func (s *usSocket) OutputUSPower() float64 {
	return USPowerVal
}
