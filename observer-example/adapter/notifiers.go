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
}

func NewEmailNotifier(email string) *EmailNotifier {
	return &EmailNotifier{emailAddress: email}
}

func (e *EmailNotifier) OnUpdate(event string) {
	// Simulate sending an email
	fmt.Printf("📧 [Email to %s] Received update: %s\n", e.emailAddress, event)
}

// --- 2. Slack Notifier ---

type SlackNotifier struct {
	webhookID string
}

func NewSlackNotifier(webhookID string) *SlackNotifier {
	return &SlackNotifier{webhookID: webhookID}
}

func (s *SlackNotifier) OnUpdate(event string) {
	// Simulate posting to Slack
	fmt.Printf("💬 [Slack #%s] 🚨 Notification: %s\n", s.webhookID, event)
}

// --- 3. Logger Notifier ---

type LogNotifier struct{}

func NewLogNotifier() *LogNotifier {
	return &LogNotifier{}
}

func (l *LogNotifier) OnUpdate(event string) {
	fmt.Printf("📝 [System Log] Event recorded: %s\n", event)
}
