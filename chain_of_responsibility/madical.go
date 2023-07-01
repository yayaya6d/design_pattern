package chain_of_responsibility

import "fmt"

type Madical struct {
	next  Department
	count int
}

func (r *Madical) Execute(patient *Patient) {
	if patient.MedicineDone {
		fmt.Println("Madicine already given to patient")
	} else {
		fmt.Println("Give medicine to patient")
		patient.MedicineDone = true
		r.count++
	}
}

func (r *Madical) SetNext(nextDepartment Department) {
	r.next = nextDepartment
}

func (r *Madical) ExecuteCount() int {
	return r.count
}
