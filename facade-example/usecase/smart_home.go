package usecase

import (
	"facade-example/domain"
	"fmt"
)

// SmartHomeFacade hides the complexity of various home automation subsystems.
type SmartHomeFacade struct {
	light     domain.Lighting
	audio     domain.AudioSystem
	projector domain.Projector
	screen    domain.Screen
	coffee    domain.CoffeeMaker
	logger    domain.Logger
}

// NewSmartHomeFacade injects all dependencies.
func NewSmartHomeFacade(
	l domain.Lighting,
	a domain.AudioSystem,
	p domain.Projector,
	s domain.Screen,
	c domain.CoffeeMaker,
	logger domain.Logger,
) *SmartHomeFacade {
	return &SmartHomeFacade{
		light:     l,
		audio:     a,
		projector: p,
		screen:    s,
		coffee:    c,
		logger:    logger,
	}
}

// StartMovieMode prepares the room for a movie experience.
func (f *SmartHomeFacade) StartMovieMode(movieName string) {
	f.logger.Log(fmt.Sprintf("\n=== Starting Movie Mode for '%s' ===", movieName))
	f.coffee.Off() // Safety first
	f.light.Dim(10)
	f.screen.Down()
	f.projector.On()
	f.projector.SetInput("BluRay")
	f.projector.WideScreenMode()
	f.audio.On()
	f.audio.SetSource("Projector")
	f.audio.SetVolume(35)
	f.logger.Log(">>> POPCORN READY! <<<")
}

// EndMovieMode restores the room.
func (f *SmartHomeFacade) EndMovieMode() {
	f.logger.Log("\n=== Shutting Down Movie Mode ===")
	f.audio.Off()
	f.projector.Off()
	f.screen.Up()
	f.light.On()
}

// GoodMorningMode prepares the house for the day.
func (f *SmartHomeFacade) GoodMorningMode() {
	f.logger.Log("\n=== Good Morning! ===")
	f.light.On()
	f.audio.On()
	f.audio.SetSource("Spotify-Morning-Jazz")
	f.audio.SetVolume(15)
	f.coffee.On()
	f.coffee.Brew()
}
