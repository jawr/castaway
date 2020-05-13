package controller

import (
	"github.com/jawr/castaway/internal/entity"
)

type Component struct {
	entity.Entity

	speed float64
}

func NewComponent(e entity.Entity, speed float64) *Component {
	return &Component{
		Entity: e,
		speed:  speed,
	}
}

func (c *Component) GetEntity() entity.Entity { return c.Entity }
