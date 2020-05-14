package controller

import (
	"github.com/jawr/castaway/internal/component/animator"
	"github.com/jawr/castaway/internal/component/position"
	"github.com/jawr/castaway/internal/component/speed"
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
	"github.com/jawr/castaway/internal/system"
	"github.com/jawr/castaway/internal/system/input"
	"github.com/peterhellberg/gfx"
)

type Controller struct {
	entities []entity.Entity
}

func NewController() *Controller {
	return &Controller{
		entities: make([]entity.Entity, 0),
	}
}

// Add Entity and it's entity.Entity to the system
func (a *Controller) Add(e entity.Entity, flags entity.ComponentFlags) {
	if !flags.Contains(entity.ComponentAnimator) && !flags.Contains(entity.ComponentPosition, entity.ComponentSpeed) {
		return
	}

	for _, check := range a.entities {
		if check == e {
			return
		}
	}

	a.entities = append(a.entities, e)
}

// remove Entity from System
func (a *Controller) Remove(e entity.Entity) {
	for idx := 0; idx < len(a.entities); idx++ {
		if a.entities[idx] == e {
			a.entities[idx] = a.entities[len(a.entities)-1]
			a.entities = a.entities[:len(a.entities)-1]
			break
		}
	}
}

// update entities in the system
func (a *Controller) Update(emanager *entity.Manager, publish event.Publisher) error {
	for _, e := range a.entities {
		if !emanager.Exists(e) {
			a.Remove(e)
			continue
		}
	}
	return nil
}

// initialise the system by setting up subscriptions to topics
func (a *Controller) SetupSubscriptions(emanager *entity.Manager, publish event.Publisher, subscriber event.Subscriber) {
	subscriber(event.TopicDirectionKeyPressed, a.handleDirectionKeyPressed(emanager, publish))
}

// Used to check the type of this System
func (a *Controller) Type() system.SystemType {
	return system.SystemTypeController
}

func (a *Controller) handleDirectionKeyPressed(emanager *entity.Manager, publish event.Publisher) event.Subscription {
	// how do we close these bad boys
	ch := make(event.Subscription, 1000)
	go func() {
		for in := range ch {
			ev, ok := in.(*input.EventDirectionKey)
			if !ok {
				continue
			}

			// if we have an animator, update it
			for _, e := range a.entities {
				com, ok := emanager.GetComponent(e, entity.ComponentAnimator)
				if ok {
					com.(*animator.Component).SetRow(int(ev.Direction))
					com.(*animator.Component).NextFrame()
				}

				com, ok = emanager.GetComponent(e, entity.ComponentPosition)
				if ok {

					pos := com.(*position.Component)
					ox := pos.X
					oy := pos.Y

					scom, ok := emanager.GetComponent(e, entity.ComponentSpeed)
					if ok {
						pos.Move(ev.Direction, scom.(*speed.Component).Speed())

						// publish move event
						publish(&EventMove{
							Entity: e,
							Origin: gfx.Vec{ox, oy},
							New:    pos.Vec,
						})
					}
				}
			}

		}
	}()
	return ch
}
