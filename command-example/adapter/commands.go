package adapter

import (
	"command-example/domain"
	"fmt"
)

// Ensure implementation
var (
	_ domain.Command = (*InsertCommand)(nil)
	_ domain.Command = (*DeleteCommand)(nil)
)

// --- 1. Insert Command ---

type InsertCommand struct {
	textToInsert string
	logger       domain.Logger
}

func NewInsertCommand(text string, logger domain.Logger) *InsertCommand {
	return &InsertCommand{
		textToInsert: text,
		logger:       logger,
	}
}

func (c *InsertCommand) Do(b *domain.Buffer) {
	b.Content += c.textToInsert
	c.logger.Log(fmt.Sprintf("[CMD] Inserted: '%s'", c.textToInsert))
}

func (c *InsertCommand) Undo(b *domain.Buffer) {
	length := len(c.textToInsert)
	currentLen := len(b.Content)
	if currentLen >= length {
		b.Content = b.Content[:currentLen-length]
		c.logger.Log(fmt.Sprintf("[CMD] Undid Insert: Removed '%s'", c.textToInsert))
	}
}

// --- 2. Delete Command (Stateful) ---

// DeleteCommand removes the last N characters.
// It remembers what was deleted to support Undo.
type DeleteCommand struct {
	count       int
	deletedText string // State to save for Undo
	logger      domain.Logger
}

func NewDeleteCommand(count int, logger domain.Logger) *DeleteCommand {
	return &DeleteCommand{
		count:  count,
		logger: logger,
	}
}

func (c *DeleteCommand) Do(b *domain.Buffer) {
	currentLen := len(b.Content)
	if currentLen < c.count {
		c.count = currentLen // Adjust to available length
	}

	// Save state for Undo
	start := currentLen - c.count
	c.deletedText = b.Content[start:]

	// Execute
	b.Content = b.Content[:start]
	c.logger.Log(fmt.Sprintf("[CMD] Deleted last %d chars: '%s'", c.count, c.deletedText))
}

func (c *DeleteCommand) Undo(b *domain.Buffer) {
	// Restore state
	b.Content += c.deletedText
	c.logger.Log(fmt.Sprintf("[CMD] Undid Delete: Restored '%s'", c.deletedText))
}
