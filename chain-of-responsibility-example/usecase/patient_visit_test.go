package usecase_test

import (
	"testing"

	"chain-of-responsibility-example/domain"
	"chain-of-responsibility-example/usecase"
)

// MockDepartment for testing the UseCase independently of adapters
type MockDepartment struct {
	Executed bool
	Next     domain.Department
}

func (m *MockDepartment) Execute(p *domain.Patient) {
	m.Executed = true
	if m.Next != nil {
		m.Next.Execute(p)
	}
}

func (m *MockDepartment) SetNext(next domain.Department) {
	m.Next = next
}

func TestPatientVisitService_VisitHospital(t *testing.T) {
	// Arrange
	mockHead := &MockDepartment{}
	service := usecase.NewPatientVisitService(mockHead)
	patient := &domain.Patient{Name: "Test Patient"}

	// Act
	service.VisitHospital(patient)

	// Assert
	if !mockHead.Executed {
		t.Error("Expected the chain head to be executed")
	}
}
