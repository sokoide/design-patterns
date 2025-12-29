package domain

// Lighting controls the room lights.
type Lighting interface {
	On()
	Off()
	Dim(level int)
}

// AudioSystem controls playback and audio settings.
type AudioSystem interface {
	On()
	Off()
	SetVolume(vol int)
	SetSource(source string)
}

// Projector controls display projection.
type Projector interface {
	On()
	Off()
	SetInput(input string)
	WideScreenMode()
}

// Screen controls the projection screen.
type Screen interface {
	Up()
	Down()
}

// CoffeeMaker controls coffee machine operations.
type CoffeeMaker interface {
	On()
	Off()
	Brew()
}

// Logger defines the interface for logging.
type Logger interface {
	Log(message string)
}
