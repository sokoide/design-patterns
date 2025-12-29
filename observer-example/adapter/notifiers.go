package adapter

import (
	"fmt"
	"observer-example/domain"
)

// Ensure implementation
var (
	_ domain.Observer = (*EmailNotifier)(nil)
	_ domain.Observer = (*SlackNotifier)(nil)
	_ domain.Observer = (*LogNotifier)(nil)
)

// --- 1. Email Notifier ---

type EmailNotifier struct {
	emailAddress string
	logger       domain.Logger
}

func NewEmailNotifier(email string, logger domain.Logger) *EmailNotifier {
	return &EmailNotifier{
		emailAddress: email,
		logger:       logger,
	}
}

func (e *EmailNotifier) OnUpdate(event string) {
	e.logger.Log(fmt.Sprintf("📧 [Email to %s] Received update: %s", e.emailAddress, event))
}

// --- 2. Slack Notifier ---

type SlackNotifier struct {
	webhookID string
	logger    domain.Logger
}

func NewSlackNotifier(webhookID string, logger domain.Logger) *SlackNotifier {
	return &SlackNotifier{
		webhookID: webhookID,
		logger:    logger,
	}
}

func (s *SlackNotifier) OnUpdate(event string) {
	s.logger.Log(fmt.Sprintf("💬 [Slack #%s] 🚨 Notification: %s", s.webhookID, event))
}

// --- 3. Logger Notifier ---
// Ideally this wraps the logger itself to log events as a distinct observer.

type LogNotifier struct {
	logger domain.Logger
}

func NewLogNotifier(logger domain.Logger) *LogNotifier {
	return &LogNotifier{logger: logger}
}

func (l *LogNotifier) OnUpdate(event string) {
	l.logger.Log(fmt.Sprintf("📝 [System Log] Event recorded: %s", event))
}
