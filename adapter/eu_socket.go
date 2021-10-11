package adapter

var EUPowerVal float64 = 110.0

type EUSocket interface {
	OutputEUPower() float64
}

func NewEUSocket() EUSocket {
	return &euSocket{}
}

type euSocket struct {
}

func (s *euSocket) OutputEUPower() float64 {
	return EUPowerVal
}
