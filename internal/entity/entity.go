package entity

import (
	"github.com/google/uuid"
)

// Entity is a UUID
type Entity uuid.UUID

// Create a new Entity, keep trying until a unique
// UUID is created
func NewEntity() Entity {
	return Entity(uuid.New())
}

func (e Entity) String() string {
	return uuid.UUID(e).String()
}
