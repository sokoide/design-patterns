package domain

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
