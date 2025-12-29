package adapter_test

import (
	"testing"

	"chain-of-responsibility-example/adapter"
	"chain-of-responsibility-example/domain"
)

type MockLogger struct {
	Logs []string
}

func (m *MockLogger) Log(message string) {
	m.Logs = append(m.Logs, message)
}

func TestHospitalChain(t *testing.T) {
	logger := &MockLogger{}

	cashier := adapter.NewCashier(logger)
	doctor := adapter.NewDoctor(logger)
	reception := adapter.NewReception(logger)

	// Chain
	reception.SetNext(doctor)
	doctor.SetNext(cashier)

	patient := &domain.Patient{Name: "TestPatient"}

	// Run
	reception.Execute(patient)

	// Verify Patient State
	if !patient.RegistrationDone {
		t.Error("Registration not done")
	}
	if !patient.DoctorCheckUpDone {
		t.Error("Doctor checkup not done")
	}
	if !patient.PaymentDone {
		t.Error("Payment not done")
	}

	// Verify Logs
	expectedLogs := []string{
		"Reception registering patient",
		"Doctor checking patient",
		"Cashier getting money",
	}

	if len(logger.Logs) != 3 {
		t.Fatalf("expected 3 logs, got %d", len(logger.Logs))
	}

	for i, log := range logger.Logs {
		if log != expectedLogs[i] {
			t.Errorf("log %d mismatch: got %q, want %q", i, log, expectedLogs[i])
		}
	}

	// Run again (Already done)
	logger.Logs = nil // clear logs
	reception.Execute(patient)

	expectedLogsDone := []string{
		"Patient registration already done",
		"Doctor checkup already done",
		"Payment already done",
	}

	if len(logger.Logs) != 3 {
		t.Fatalf("expected 3 logs (2nd run), got %d", len(logger.Logs))
	}

	for i, log := range logger.Logs {
		if log != expectedLogsDone[i] {
			t.Errorf("log %d mismatch (2nd run): got %q, want %q", i, log, expectedLogsDone[i])
		}
	}
}
