package chain_of_responsibility

func BuildHositital() Department {
	reception := new(Reception)
	doctor := new(Doctor)
	medical := new(Madical)

	doctor.SetNext(medical)
	reception.SetNext(doctor)

	return reception
}
