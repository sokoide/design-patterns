package facade

import (
	"facade-example/subsystems"
	"fmt"
)

// SmartHomeFacade hides the complexity of various home automation subsystems.
type SmartHomeFacade struct {
	light     *subsystems.Lighting
	audio     *subsystems.AudioSystem
	projector *subsystems.Projector
	screen    *subsystems.Screen
	coffee    *subsystems.CoffeeMaker
}

// NewSmartHomeFacade injects all dependencies.
// In a real app, these might be singleton instances.
func NewSmartHomeFacade(
	l *subsystems.Lighting,
	a *subsystems.AudioSystem,
	p *subsystems.Projector,
	s *subsystems.Screen,
	c *subsystems.CoffeeMaker,
) *SmartHomeFacade {
	return &SmartHomeFacade{
		light:     l,
		audio:     a,
		projector: p,
		screen:    s,
		coffee:    c,
	}
}

// StartMovieMode prepares the room for a movie experience.
func (f *SmartHomeFacade) StartMovieMode(movieName string) {
	fmt.Printf("\n=== Starting Movie Mode for '%s' ===\n", movieName)
	f.coffee.Off() // Safety first
	f.light.Dim(10)
	f.screen.Down()
	f.projector.On()
	f.projector.SetInput("BluRay")
	f.projector.WideScreenMode()
	f.audio.On()
	f.audio.SetSource("Projector")
	f.audio.SetVolume(35)
	fmt.Println(">>> POPCORN READY! <<<")
}

// EndMovieMode restores the room.
func (f *SmartHomeFacade) EndMovieMode() {
	fmt.Println("\n=== Shutting Down Movie Mode ===")
	f.audio.Off()
	f.projector.Off()
	f.screen.Up()
	f.light.On()
}

// GoodMorningMode prepares the house for the day.
func (f *SmartHomeFacade) GoodMorningMode() {
	fmt.Println("\n=== Good Morning! ===")
	f.light.On()
	f.audio.On()
	f.audio.SetSource("Spotify-Morning-Jazz")
	f.audio.SetVolume(15)
	f.coffee.On()
	f.coffee.Brew()
}
