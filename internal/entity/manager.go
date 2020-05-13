package entity

import (
	"log"
)

type Manager struct {
	// stored flags for an Entity
	flags map[Entity]ComponentFlags
	// stored components for an Entity
	components map[Entity][]Component
}

// Create a new Entity Manager
func NewManager() *Manager {
	return &Manager{
		flags:      make(map[Entity]ComponentFlags, 0),
		components: make(map[Entity][]Component, 0),
	}
}

// Check to see if an Entity exists
func (m *Manager) Exists(e Entity) bool {
	_, ok := m.flags[e]
	return ok
}

func (m *Manager) Add(e Entity) {
	if m.Exists(e) {
		log.Fatal("Tried adding an Entity that already exists")
	}

	m.flags[e] = newComponentFlags()
	m.components[e] = make([]Component, 0)
}

func (m *Manager) AddComponent(c Component) (Entity, ComponentFlags) {
	e := c.GetEntity()

	if !m.Exists(e) {
		log.Fatal("Tried adding a Component to an Entity that doesn't exist")
	}

	if m.flags[e].Contains(c.GetFlag()) {
		log.Fatal("Tried adding a Component to an Entity that already had the Component")
	}

	m.flags[e].Add(c.GetFlag())
	m.components[e] = append(m.components[e], c)

	return e, m.flags[e]
}

func (m *Manager) GetComponent(e Entity, flag ComponentFlag) (Component, bool) {
	if !m.Exists(e) {
		return nil, false
	}

	for _, c := range m.components[e] {
		if c.GetFlag() == flag {
			return c, true
		}
	}

	return nil, false
}

func (m *Manager) RemoveComponent(c Component) ComponentFlags {
	e := c.GetEntity()

	if !m.Exists(e) {
		log.Fatal("Tried removing a Component from an Entity that doesn't exist")
	}

	if !m.flags[e].Contains(c.GetFlag()) {
		log.Fatal("Tried removing a Component from an Entity that doesn't have the Component")
	}

	m.flags[e].Remove(c.GetFlag())
	for idx := 0; idx < len(m.components[e]); idx++ {
		if m.components[e][idx].GetFlag() == c.GetFlag() {
			m.components[e][idx] = m.components[e][len(m.components[e])-1]
			m.components[e] = m.components[e][:len(m.components[e])-1]
			break
		}
	}

	return m.flags[e]
}

// Remove an Entity from the manager
func (m *Manager) Remove(e Entity) {
	delete(m.flags, e)
	delete(m.components, e)
}
