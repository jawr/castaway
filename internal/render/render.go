package render

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/system"
	"github.com/jawr/castaway/internal/system/animator"
	"github.com/jawr/castaway/internal/system/position"
	"github.com/jawr/castaway/internal/system/sprite"
)

type Manager struct {
	components []entity.Component
}

func NewManager() *Manager {
	return &Manager{
		components: make([]entity.Component, 0),
	}
}

func (m *Manager) AddComponent(c entity.Component) {
	m.components = append(m.components, c)
}

func (m *Manager) Remove(e entity.Entity) {
	// hopefully this will remove all components of entity, needs testing
	for idx := 0; idx < len(m.components); idx++ {
		if m.components[idx].GetEntity() == e {
			m.components[idx] = m.components[len(m.components)-1]
			m.components = m.components[:len(m.components)-1]
		}
	}
}

func (m *Manager) Draw(entities *entity.Manager, systems *system.Manager, screen *ebiten.Image) {
	for _, c := range m.components {
		e := c.GetEntity()

		if !entities.Exists(e) {
			m.Remove(e)
			continue
		}

		op := &ebiten.DrawImageOptions{}

		// not safe
		image := c.(*sprite.Component).Sprite()

		// get frame
		com, ok := systems.Get(system.SystemTypeAnimator).Get(e)
		if ok {
			frame := com.(*animator.Component)
			image = image.SubImage(frame.GetFrameRect()).(*ebiten.Image)
		}

		// get position
		com, ok = systems.Get(system.SystemTypePosition).Get(e)
		if ok {
			vec := com.(*position.Component)
			op.GeoM.Translate(vec.X(), vec.Y())
		}

		screen.DrawImage(image, op)
	}
}
