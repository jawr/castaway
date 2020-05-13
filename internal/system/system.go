package system

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
)

type SystemType uint8

type System interface {
	// Add an Entity with specific component data
	Add(entity.Entity, entity.Component)
	// Remove an Entity
	Remove(entity.Entity)
	// Called when the World updates
	Update(event.Publisher) error
	// Called to setup our Subscriptions
	SetupSubscriptions(event.Subscriber)
	// The Type of a System
	Type() SystemType
}

type Manager struct {
	// All added systems, currently only one of each can
	// be added
	systems map[SystemType]System
}

// Initialise a new Manager with listed systems
func NewManager(subscriber event.Subscriber, all ...System) *Manager {
	systems := make(map[SystemType]System, len(all))

	for _, s := range all {
		// make sure the system subscribes to the topics it wants
		s.SetupSubscriptions(subscriber)
		systems[s.Type()] = s
	}

	return &Manager{
		systems: systems,
	}
}

// Call Update on all systems
func (m *Manager) Update(publisher event.Publisher) error {
	for _, s := range m.systems {
		if err := s.Update(publisher); err != nil {
			return err
		}
	}
	return nil
}

// Try and get a System of a particular SystemType
func (m *Manager) Get(stype SystemType) (System, bool) {
	s, ok := m.systems[stype]
	return s, ok
}
