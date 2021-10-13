package decorator

type Pizza struct {
}

func (p *Pizza) GetPrice() int {
	return 10
}

func (p *Pizza) GetName() string {
	return "pizza"
}
