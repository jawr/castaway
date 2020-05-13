package position

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/system/input"
)

type Component struct {
	entity.Entity

	x, y float64
}

func NewComponent(e entity.Entity, x, y float64) *Component {
	return &Component{
		Entity: e,
		x:      x,
		y:      y,
	}
}

func (c *Component) GetEntity() entity.Entity      { return c.Entity }
func (c *Component) GetFlag() entity.ComponentFlag { return entity.ComponentPosition }

func (c Component) X() float64 { return c.x }
func (c Component) Y() float64 { return c.y }

func (c *Component) Move(direction input.Direction, speed float64) {
	switch direction {
	case input.DirectionUp:
		c.y -= speed
	case input.DirectionDown:
		c.y += speed
	case input.DirectionLeft:
		c.x -= speed
	case input.DirectionRight:
		c.x += speed
	}
}
