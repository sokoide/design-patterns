package domain

type Lighting interface {
	On()
	Off()
	Dim(level int)
}

type AudioSystem interface {
	On()
	Off()
	SetVolume(vol int)
	SetSource(source string)
}

type Projector interface {
	On()
	Off()
	SetInput(input string)
	WideScreenMode()
}

type Screen interface {
	Up()
	Down()
}

type CoffeeMaker interface {
	On()
	Off()
	Brew()
}

// Logger defines the interface for logging.
type Logger interface {
	Log(message string)
}
