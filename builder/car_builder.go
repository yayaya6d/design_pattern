package builder

type CarBuilder interface {
	Reset()
	SetSeats(int)
	SetEngine()
	SetSkin()
	GetResultCar() Car
}
