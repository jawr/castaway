package sprite

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
	"github.com/jawr/castaway/internal/system"
)

type Sprite struct {
	components []*Component
	lookup     map[entity.Entity]*Component
}

func NewSprite() *Sprite {
	return &Sprite{
		components: make([]*Component, 0),
		lookup:     make(map[entity.Entity]*Component, 0),
	}
}

// Add Entity and it's Component to the system
func (a *Sprite) Add(c entity.Component) {
	// not thread safe

	ac, ok := c.(*Component)
	if !ok {
		return
	}

	if _, ok := a.lookup[ac.Entity]; ok {
		return
	}

	a.components = append(a.components, ac)
	a.lookup[ac.Entity] = ac
}

// get a component for an Entity
func (a *Sprite) Get(e entity.Entity) (entity.Component, bool) {
	c, ok := a.lookup[e]
	return c, ok
}

// remove Entity from System
func (a *Sprite) Remove(e entity.Entity) {
	if _, ok := a.lookup[e]; ok {
		for idx := 0; idx < len(a.components); idx++ {
			if a.components[idx].Entity == e {
				a.components[idx] = a.components[len(a.components)-1]
				a.components = a.components[:len(a.components)-1]
				break
			}
		}

		delete(a.lookup, e)
	}
}

// update components in the system
func (a *Sprite) Update(emanager *entity.Manager, systems *system.Manager, publisher event.Publisher) error {
	for _, c := range a.components {
		if !emanager.Exists(c.Entity) {
			a.Remove(c.Entity)
			continue
		}
	}
	return nil
}

// initialise the system by setting up subscriptions to topics
func (a *Sprite) SetupSubscriptions(systems *system.Manager, subscriber event.Subscriber) {}

// Used to check the type of this System
func (a *Sprite) Type() system.SystemType {
	return system.SystemTypeSprite
}
