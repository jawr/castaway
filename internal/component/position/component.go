package position

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/system/input"
	"github.com/peterhellberg/gfx"
)

type Component struct {
	entity.Entity

	gfx.Vec
}

func NewComponent(e entity.Entity, x, y float64) *Component {
	return &Component{
		Entity: e,
		Vec:    gfx.Vec{x, y},
	}
}

func (c *Component) GetEntity() entity.Entity      { return c.Entity }
func (c *Component) GetFlag() entity.ComponentFlag { return entity.ComponentPosition }

func (c *Component) Move(direction input.Direction, speed float64) {
	switch direction {
	case input.DirectionUp:
		c.Y -= speed
	case input.DirectionDown:
		c.Y += speed
	case input.DirectionLeft:
		c.X -= speed
	case input.DirectionRight:
		c.X += speed
	}
}
