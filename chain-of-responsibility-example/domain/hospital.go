package domain

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	PaymentDone       bool
}

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Logger interface {
	Log(message string)
}
