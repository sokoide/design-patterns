package usecase

import "chain-of-responsibility-example/domain"

type PatientVisitService struct {
	chainHead domain.Department
}

func NewPatientVisitService(chainHead domain.Department) *PatientVisitService {
	return &PatientVisitService{
		chainHead: chainHead,
	}
}

func (s *PatientVisitService) VisitHospital(p *domain.Patient) {
	if s.chainHead != nil {
		s.chainHead.Execute(p)
	}
}
