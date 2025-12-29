package domain

// Computer represents a machine that accepts a lightning connector.
type Computer interface {
	InsertIntoLightningPort()
}

// Logger abstracts logging for the domain.
type Logger interface {
	Log(message string)
}
