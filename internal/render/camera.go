package render

import (
	"log"

	"github.com/jawr/castaway/internal/event"
	"github.com/jawr/castaway/internal/system/controller"
	"github.com/peterhellberg/gfx"
)

type camera struct {
	gfx.Vec

	width, height int

	bound                   gfx.Vec
	boundWidth, boundHeight float64
}

func newCamera(width, height int) *camera {
	boundWidth := float64(width) / 2.0
	boundHeight := float64(height) / 2.0

	boundX := (float64(width) - boundWidth) / 2.0
	boundY := (float64(height) - boundHeight) / 2.0

	return &camera{
		Vec: gfx.Vec{0.0, 0.0},

		width:  width,
		height: height,

		bound:       gfx.Vec{boundX, boundY},
		boundWidth:  boundWidth,
		boundHeight: boundHeight,
	}
}

func (c *camera) handleMove() event.Subscription {
	ch := make(event.Subscription, 1000)
	go func() {
		for in := range ch {
			ev, ok := in.(*controller.EventMove)
			if !ok {
				continue
			}

			if c.isInBounds(ev.Origin.X, ev.Origin.Y) {
				d := ev.Origin.Sub(ev.New)

				log.Printf(
					"%s moved from (%s) to (%s) d (%s)",
					ev.Entity, ev.Origin, ev.New, d,
				)

				// are we hitting the left wall
				// or the right wall
				if ev.New.X == c.bound.X {
					c.X -= d.X
					c.bound.X -= d.X

				} else if ev.New.X == c.bound.X+c.boundWidth {
					c.X -= d.X
					c.bound.X -= d.X
				}

				// are we hitting the top
				// or the bottom wall
				if ev.New.Y == c.bound.Y {
					c.Y -= d.Y
					c.bound.Y -= d.Y

				} else if ev.New.Y == c.bound.Y+c.boundHeight {
					c.Y -= d.Y
					c.bound.Y -= d.Y
				}
			}

		}
	}()
	return ch
}

func (c *camera) isInBounds(x, y float64) bool {
	return c.bound.X < x && x < c.bound.X+c.boundWidth && c.bound.Y < y && y < c.bound.Y+c.boundHeight
}
