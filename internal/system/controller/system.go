package controller

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
	"github.com/jawr/castaway/internal/system"
	"github.com/jawr/castaway/internal/system/animator"
	"github.com/jawr/castaway/internal/system/input"
	"github.com/jawr/castaway/internal/system/position"
)

type Controller struct {
	components []*Component
	lookup     map[entity.Entity]*Component
}

func NewController() *Controller {
	return &Controller{
		components: make([]*Component, 0),
		lookup:     make(map[entity.Entity]*Component, 0),
	}
}

// Add Entity and it's Component to the system
func (a *Controller) Add(c entity.Component) {
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
func (a *Controller) Get(e entity.Entity) (entity.Component, bool) {
	c, ok := a.lookup[e]
	return c, ok
}

// remove Entity from System
func (a *Controller) Remove(e entity.Entity) {
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
func (a *Controller) Update(emanager *entity.Manager, systems *system.Manager, publisher event.Publisher) error {
	for _, c := range a.components {
		if !emanager.Exists(c.Entity) {
			a.Remove(c.Entity)
			continue
		}
	}
	return nil
}

// initialise the system by setting up subscriptions to topics
func (a *Controller) SetupSubscriptions(systems *system.Manager, subscriber event.Subscriber) {
	subscriber(event.TopicDirectionKeyPressed, a.handleDirectionKeyPressed(systems))
}

// Used to check the type of this System
func (a *Controller) Type() system.SystemType {
	return system.SystemTypeController
}

func (a *Controller) handleDirectionKeyPressed(systems *system.Manager) event.Subscription {
	// how do we close these bad boys
	ch := make(event.Subscription, 1000)
	go func() {
		for ev := range ch {
			e, ok := ev.(*input.EventDirectionKey)
			if !ok {
				continue
			}

			// if we have an animator, update it
			for _, c := range a.components {
				com, ok := systems.Get(system.SystemTypeAnimator).Get(c.Entity)
				if ok {
					com.(*animator.Component).SetRow(int(e.Direction))
					com.(*animator.Component).NextFrame()
				}

				com, ok = systems.Get(system.SystemTypePosition).Get(c.Entity)
				if ok {
					com.(*position.Component).Move(e.Direction, c.speed)
				}
			}

		}
	}()
	return ch
}
