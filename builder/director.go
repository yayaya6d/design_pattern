package builder

type Director interface {
	SetCarBuilder(CarBuilder)
	GetCurCarBuilder() CarBuilder
	MakeSportCar() Car
	MakeSUVCar() Car
}

type director struct {
	builder CarBuilder
}

func (d *director) SetCarBuilder(carBulider CarBuilder) {
	d.builder = carBulider
}

func (d *director) GetCurCarBuilder() CarBuilder {
	return d.builder
}

func (d director) MakeSportCar() Car {
	d.builder.Reset()
	d.builder.SetSeats(2) // sport car only has 2 seats
	d.builder.SetEngine()
	d.builder.SetSkin()
	return d.builder.GetResultCar()
}

func (d director) MakeSUVCar() Car {
	d.builder.Reset()
	d.builder.SetSeats(6) // SUV car has 6 seats
	d.builder.SetEngine()
	d.builder.SetSkin()
	return d.builder.GetResultCar()
}

func NewDirector(carBuilder CarBuilder) Director {
	return &director{
		builder: carBuilder,
	}
}
