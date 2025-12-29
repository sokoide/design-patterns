package server

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	t.Run("Default Configuration", func(t *testing.T) {
		s := New()
		if s.host != "localhost" {
			t.Errorf("expected host localhost, got %s", s.host)
		}
		if s.port != 8080 {
			t.Errorf("expected port 8080, got %d", s.port)
		}
		if s.timeout != 30*time.Second {
			t.Errorf("expected timeout 30s, got %v", s.timeout)
		}
		if s.maxConn != 100 {
			t.Errorf("expected maxConn 100, got %d", s.maxConn)
		}
		if s.tls != false {
			t.Error("expected tls false, got true")
		}
	})

	t.Run("Custom Configuration", func(t *testing.T) {
		s := New(
			WithHost("example.com"),
			WithPort(9000),
			WithTimeout(10*time.Second),
			WithMaxConn(500),
			WithTLS(true),
		)

		if s.host != "example.com" {
			t.Errorf("expected host example.com, got %s", s.host)
		}
		if s.port != 9000 {
			t.Errorf("expected port 9000, got %d", s.port)
		}
		if s.timeout != 10*time.Second {
			t.Errorf("expected timeout 10s, got %v", s.timeout)
		}
		if s.maxConn != 500 {
			t.Errorf("expected maxConn 500, got %d", s.maxConn)
		}
		if s.tls != true {
			t.Error("expected tls true, got false")
		}
	})
}
