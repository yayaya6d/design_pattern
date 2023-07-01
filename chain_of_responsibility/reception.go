package chain_of_responsibility

import "fmt"

type Reception struct {
	next  Department
	count int
}

func (r *Reception) Execute(patient *Patient) {
	if patient.RegisterDone {
		fmt.Println("Patient registration already done")
	} else {
		fmt.Println("Reciption registering patient")
		patient.RegisterDone = true
		r.count++
	}

	r.next.Execute(patient)
}

func (r *Reception) SetNext(nextDepartment Department) {
	r.next = nextDepartment
}

func (r *Reception) ExecuteCount() int {
	return r.count
}
