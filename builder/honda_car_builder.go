package builder

type HondaCarBuilder struct {
	result Car
}

func (b *HondaCarBuilder) Reset() {
	b.result = Car{}
}

func (b *HondaCarBuilder) SetSeats(count int) {
	b.result.SeatCount = count
}

func (b *HondaCarBuilder) SetEngine() {
	hondaEngine := Engine{
		name:  "Honda 666",
		power: 8,
	}
	b.result.Engine = hondaEngine
}

func (b *HondaCarBuilder) SetSkin() {
	hondaSkin := Skin{
		color:    Red,
		hardness: 20,
	}
	b.result.Skin = hondaSkin
}

func (b *HondaCarBuilder) GetResultCar() Car {
	return b.result
}
