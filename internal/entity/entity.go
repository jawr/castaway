package entity

import (
	"log"

	"github.com/google/uuid"
)

// Entity is a UUID
type Entity uuid.UUID

type Manager struct {
	entities map[Entity]struct{}
}

// Create a new Entity, keep trying until a unique
// UUID is created
func NewEntity() Entity {
	return Entity(uuid.New())
}

// Create a new Entity Manager
func NewManager() *Manager {
	return &Manager{
		entities: make(map[Entity]struct{}, 0),
	}
}

// Check to see if an Entity exists
func (m *Manager) Exists(e Entity) bool {
	_, ok := m.entities[e]
	return ok
}

func (m *Manager) Add(e Entity) {
	if m.Exists(e) {
		log.Fatal("Entity already exists")
	}
	m.entities[e] = struct{}{}
}

// Remove an Entity from the manager
func (m *Manager) Remove(e Entity) {
	delete(m.entities, e)
}

func (e Entity) String() string {
	return uuid.UUID(e).String()
}
