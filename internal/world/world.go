package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
	"github.com/jawr/castaway/internal/render"
	"github.com/jawr/castaway/internal/system"
	"github.com/jawr/castaway/internal/system/sprite"
)

// World represents ... the world, maybe one day we will have
// multiple worlds
type World struct {
	// constants
	screenWidth, screenHeight int

	// managers
	entities *entity.Manager
	events   *event.Manager
	systems  *system.Manager
	renderer *render.Manager
}

// Create a new World
func NewWorld(screenWidth, screenHeight int, allSystems ...system.System) *World {
	// create our entity manager
	entities := entity.NewManager()

	// create our event manager
	events := event.NewManager()

	// create our systems
	systems := system.NewManager(events.Subscribe, allSystems...)

	// create our renderer
	renderer := render.NewManager()

	return &World{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,

		// managers
		entities: entities,
		events:   events,
		systems:  systems,
		renderer: renderer,
	}
}

func (w *World) AddEntity(e entity.Entity) {
	w.entities.Add(e)
}

// check if this component is a Sprite

func (w *World) AddComponent(c entity.Component) {
	w.systems.AddComponent(c)
	// check if this can be rendered
	switch c.(type) {
	case *sprite.Component:
		w.renderer.AddComponent(c)
	}
}

// implement ebiten Game interface

// Update all Systems
func (w *World) Update(screen *ebiten.Image) error {
	return w.systems.Update(w.entities, w.events.Publish)
}

// Render!
func (w *World) Draw(screen *ebiten.Image) {
	w.renderer.Draw(w.entities, w.systems, screen)
}

func (w *World) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w.screenWidth, w.screenHeight
}
