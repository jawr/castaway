package entity

import "github.com/google/uuid"

// Entity is a UUID
type Entity uuid.UUID

type Manager struct {
	entities map[Entity]struct{}
}

// Create a new Entity Manager
func NewManager() *Manager {
	return &Manager{
		entities: make(map[Entity]struct{}, 0),
	}
}

// Create a new Entity, keep trying until a unique
// UUID is created
func (m *Manager) NewEntity() Entity {
	for {
		e := Entity(uuid.New())
		if _, ok := m.entities[e]; ok {
			continue
		}
		m.entities[e] = struct{}{}
		return e
	}
}

// Check to see if an Entity exists
func (m *Manager) Exists(e Entity) bool {
	_, ok := m.entities[e]
	return ok
}

// Remove an Entity from the manager
func (m *Manager) Remove(e Entity) {
	delete(m.entities, e)
}
