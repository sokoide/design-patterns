package main

import (
	"fmt"
	"time"
)

// Server represents a complex server with many configuration options.
type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
	tls     bool
}

// Option is a function type that modifies the Server configuration.
type Option func(*Server)

// NewServer creates a new Server with default settings, applying any provided Options.
func NewServer(options ...Option) *Server {
	// 1. Set default values
	srv := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30 * time.Second,
		maxConn: 100,
		tls:     false,
	}

	// 2. Apply options
	for _, opt := range options {
		opt(srv)
	}

	return srv
}

// WithHost sets the host.
func WithHost(h string) Option {
	return func(s *Server) {
		s.host = h
	}
}

// WithPort sets the port.
func WithPort(p int) Option {
	return func(s *Server) {
		s.port = p
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.timeout = t
	}
}

// WithMaxConn sets the maximum number of connections.
func WithMaxConn(n int) Option {
	return func(s *Server) {
		s.maxConn = n
	}
}

// WithTLS enables TLS.
func WithTLS(enable bool) Option {
	return func(s *Server) {
		s.tls = enable
	}
}

// Start simulates starting the server.
func (s *Server) Start() {
	scheme := "http"
	if s.tls {
		scheme = "https"
	}
	fmt.Printf("[Server] Starting at %s://%s:%d\n", scheme, s.host, s.port)
	fmt.Printf("         Config: Timeout=%v, MaxConn=%d\n", s.timeout, s.maxConn)
}
