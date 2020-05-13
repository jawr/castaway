package input

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
	"github.com/jawr/castaway/internal/system"
)

type Input struct{}

func NewInput() *Input {
	return &Input{}
}

// Add Entity and it's Component to the system
func (a *Input) Add(c entity.Component) {}

// get a component for an Entity
func (a *Input) Get(e entity.Entity) (entity.Component, bool) {
	return nil, false
}

// remove Entity from System
func (a *Input) Remove(e entity.Entity) {}

// update components in the system
func (a *Input) Update(emanager *entity.Manager, systems *system.Manager, publisher event.Publisher) error {

	// TODO: be nice if this was more progmatic
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		publisher(&EventDirectionKey{DirectionUp})
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		publisher(&EventDirectionKey{DirectionDown})
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		publisher(&EventDirectionKey{DirectionLeft})
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		publisher(&EventDirectionKey{DirectionRight})
	}

	return nil
}

// initialise the system by setting up subscriptions to topics
func (a *Input) SetupSubscriptions(systems *system.Manager, subscriber event.Subscriber) {}

// Used to check the type of this System
func (a *Input) Type() system.SystemType {
	return system.SystemTypeInput
}
