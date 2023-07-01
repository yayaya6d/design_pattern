package chain_of_responsibility

type Patient struct {
	Name              string
	RegisterDone      bool
	DoctorCheckupDone bool
	MedicineDone      bool
}
