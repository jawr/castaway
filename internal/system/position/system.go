package position

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
	"github.com/jawr/castaway/internal/system"
)

type Position struct {
	components []*Component
	lookup     map[entity.Entity]*Component
}

func NewPosition() *Position {
	return &Position{
		components: make([]*Component, 0),
		lookup:     make(map[entity.Entity]*Component, 0),
	}
}

// Add Entity and it's Component to the system
func (a *Position) Add(c entity.Component) {
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
func (a *Position) Get(e entity.Entity) (entity.Component, bool) {
	c, ok := a.lookup[e]
	return c, ok
}

// remove Entity from System
func (a *Position) Remove(e entity.Entity) {
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
func (a *Position) Update(emanager *entity.Manager, systems *system.Manager, publisher event.Publisher) error {
	for _, c := range a.components {
		if !emanager.Exists(c.Entity) {
			a.Remove(c.Entity)
			continue
		}
	}
	return nil
}

// initialise the system by setting up subscriptions to topics
func (a *Position) SetupSubscriptions(systems *system.Manager, subscriber event.Subscriber) {}

// Used to check the type of this System
func (a *Position) Type() system.SystemType {
	return system.SystemTypePosition
}
