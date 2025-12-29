package domain

// Patient represents the workflow state as it moves through departments.
type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	PaymentDone       bool
}

// Department defines the chain interface.
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

// Logger abstracts logging for the domain.
type Logger interface {
	Log(message string)
}
