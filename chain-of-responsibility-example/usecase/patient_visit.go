package usecase

import "chain-of-responsibility-example/domain"

// PatientVisitService coordinates a visit through the department chain.
type PatientVisitService struct {
	chainHead domain.Department
}

// NewPatientVisitService builds a visit service with the chain head.
func NewPatientVisitService(chainHead domain.Department) *PatientVisitService {
	return &PatientVisitService{
		chainHead: chainHead,
	}
}

// VisitHospital starts the patient workflow at the chain head.
func (s *PatientVisitService) VisitHospital(p *domain.Patient) {
	if s.chainHead != nil {
		s.chainHead.Execute(p)
	}
}
