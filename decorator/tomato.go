package decorator

type Tomato struct {
	next Indredient
}

func (t *Tomato) GetPrice() int {
	return 5 + t.next.GetPrice()
}

func (t *Tomato) GetName() string {
	return "tomato " + t.next.GetName()
}
