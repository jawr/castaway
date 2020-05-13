package input

import (
	"github.com/jawr/castaway/internal/entity"
)

type Component struct {
	entity.Entity
}

func NewComponent(e entity.Entity) *Component {
	return &Component{
		Entity: e,
	}
}

func (c *Component) GetEntity() entity.Entity { return c.Entity }
