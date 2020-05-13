package system

import (
	"github.com/google/uuid"
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
)

type Null struct{}
type NullComponent struct{}

func (nc *NullComponent) GetEntity() entity.Entity { return entity.Entity(uuid.Nil) }

func newNull() *Null {
	return &Null{}
}

// Add Entity and it's NullComponent to the system
func (a *Null) Add(e entity.Entity, f entity.ComponentFlags) {}

// get a component for an Entity
func (a *Null) Get(e entity.Entity) (entity.Component, bool) {
	return nil, false
}

// remove Entity from System
func (a *Null) Remove(e entity.Entity) {}

// update components in the system
func (a *Null) Update(emanager *entity.Manager, publisher event.Publisher) error {
	return nil
}

// initialise the system by setting up subscriptions to topics
func (a *Null) SetupSubscriptions(emanager *entity.Manager, subscriber event.Subscriber) {}

// Used to check the type of this System
func (a *Null) Type() SystemType {
	return SystemTypeNull
}
