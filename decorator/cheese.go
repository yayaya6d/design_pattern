package decorator

type Cheese struct {
	next Indredient
}

func (c *Cheese) GetPrice() int {
	return 3 + c.next.GetPrice()
}

func (c *Cheese) GetName() string {
	return "cheese " + c.next.GetName()
}
