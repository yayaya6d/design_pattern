package adapter

type USPower interface {
	OutputPower() float64
}

func NewUSPower(usSocket USSocket) USPower {
	return &usPower{
		usSocket: usSocket,
	}
}

type usPower struct {
	usSocket USSocket
}

func (p *usPower) OutputPower() float64 {
	return p.usSocket.OutputUSPower()
}
