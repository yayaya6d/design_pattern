package observer

import "fmt"

type Customer struct {
	id string
}

func NewCustomer(id string) *Customer {
	return &Customer{
		id: id,
	}
}

func (c *Customer) Update(updateInfo UpdateInfo) {
	fmt.Printf("Sending email to customer %s for item %s, its avaliablity : %t\n", c.id, updateInfo.ItemName, updateInfo.Avaliable)
}

func (c *Customer) GetID() string {
	return c.id
}
