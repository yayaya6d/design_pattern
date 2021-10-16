package factory_method

type Bow struct {
	arrows int
}

func (b *Bow) GetName() string {
	return "bow"
}

func (b *Bow) Attack() int {
	if b.arrows == 0 {
		return 0
	}
	b.arrows--
	return 3
}

func (b *Bow) AttackDistance() int {
	return 10
}
