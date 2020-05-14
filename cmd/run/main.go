package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/jawr/castaway/internal/component/animator"
	"github.com/jawr/castaway/internal/component/position"
	"github.com/jawr/castaway/internal/component/speed"
	"github.com/jawr/castaway/internal/component/sprite"
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/system/controller"
	"github.com/jawr/castaway/internal/system/input"
	"github.com/jawr/castaway/internal/world"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %w\n", err)
		os.Exit(1)
	}
}

func run() error {

	// setup our world
	w := world.NewWorld(
		640,
		480,
		// add systems
		input.NewInput(),
		controller.NewController(),
	)

	island := entity.NewEntity()
	w.AddEntity(island)
	w.AddComponent(sprite.NewComponent(island, "./assets/island.png"))

	// create our player
	player := entity.NewEntity()
	w.AddEntity(player)
	w.AddComponent(animator.NewComponent(player, 16, 32, 4, 4, time.Millisecond*200))
	w.AddComponent(position.NewComponent(player, 0, 0))
	w.AddComponent(sprite.NewComponent(player, "./assets/wilson.png"))
	w.AddComponent(speed.NewComponent(player, 5.0))

	if err := ebiten.RunGame(w); err != nil {
		return err
	}

	return nil
}
