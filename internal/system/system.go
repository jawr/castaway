package system

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
)

type SystemType uint8

type System interface {
	// Add an Entity with specific component data
	Add(entity.Entity, entity.ComponentFlags)
	// Remove an Entity
	Remove(entity.Entity)
	// Called when the World updates
	Update(*entity.Manager, event.Publisher) error
	// Called to setup our Subscriptions
	SetupSubscriptions(*entity.Manager, event.Subscriber)
	// The Type of a System
	Type() SystemType
}

type Manager struct {
	// All added systems, currently only one of each can
	// be added
	systems map[SystemType]System

	// null system
	null System
}

// Initialise a new Manager with listed systems
func NewManager(entityManager *entity.Manager, subscriber event.Subscriber, all ...System) *Manager {
	manager := &Manager{
		systems: make(map[SystemType]System, len(all)),
		null:    newNull(),
	}

	for _, s := range all {
		// make sure the system subscribes to the topics it wants
		s.SetupSubscriptions(entityManager, subscriber)
		manager.systems[s.Type()] = s
	}

	return manager
}

// Call AddComponent on all Systems
func (m *Manager) AddEntity(e entity.Entity, flags entity.ComponentFlags) {
	for _, s := range m.systems {
		s.Add(e, flags)
	}
}

// Call Update on all systems
func (m *Manager) Update(entityManager *entity.Manager, publisher event.Publisher) error {
	for _, s := range m.systems {
		if err := s.Update(entityManager, publisher); err != nil {
			return err
		}
	}
	return nil
}

// Try and get a System of a particular SystemType
func (m *Manager) Get(stype SystemType) System {
	s, ok := m.systems[stype]
	if !ok {
		return m.null
	}
	return s
}
