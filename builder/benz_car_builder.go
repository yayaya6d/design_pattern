package builder

type BenzCarBuilder struct {
	result Car
}

func (b *BenzCarBuilder) Reset() {
	b.result = Car{}
}

func (b *BenzCarBuilder) SetSeats(count int) {
	b.result.SeatCount = count
}

func (b *BenzCarBuilder) SetEngine() {
	benzEngine := Engine{
		name:  "Benz 777",
		power: 10,
	}
	b.result.Engine = benzEngine
}

func (b *BenzCarBuilder) SetSkin() {
	benzSkin := Skin{
		color:    White,
		hardness: 15,
	}
	b.result.Skin = benzSkin
}

func (b *BenzCarBuilder) GetResultCar() Car {
	return b.result
}
