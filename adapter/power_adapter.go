package adapter

type EU2USPowerAdapter interface {
	OutputUSPower() float64
}

func NewEU2USPowerAdapter(euSocket EUSocket) EU2USPowerAdapter {
	return &eu2usPowerAdapter{
		euSocket: euSocket,
	}
}

type eu2usPowerAdapter struct {
	euSocket EUSocket
}

func (a *eu2usPowerAdapter) OutputUSPower() float64 {
	return a.euSocket.OutputEUPower() * 2
}
