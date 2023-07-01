package chain_of_responsibility

import "fmt"

type Doctor struct {
	next  Department
	count int
}

func (r *Doctor) Execute(patient *Patient) {
	if patient.DoctorCheckupDone {
		fmt.Println("Doctor checkup already done")
	} else {
		fmt.Println("Doctor checkint patient")
		patient.DoctorCheckupDone = true
		r.count++
	}

	r.next.Execute(patient)
}

func (r *Doctor) SetNext(nextDepartment Department) {
	r.next = nextDepartment
}

func (r *Doctor) ExecuteCount() int {
	return r.count
}
