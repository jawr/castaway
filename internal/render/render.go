package render

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jawr/castaway/internal/component/animator"
	"github.com/jawr/castaway/internal/component/position"
	"github.com/jawr/castaway/internal/component/sprite"
	"github.com/jawr/castaway/internal/entity"
)

type Manager struct {
	entities []entity.Entity
}

func NewManager() *Manager {
	return &Manager{
		entities: make([]entity.Entity, 0),
	}
}

func (m *Manager) AddEntity(c entity.Entity) {
	m.entities = append(m.entities, c)
}

func (m *Manager) Remove(e entity.Entity) {
	// hopefully this will remove all entities of entity, needs testing
	for idx := 0; idx < len(m.entities); idx++ {
		if m.entities[idx] == e {
			m.entities[idx] = m.entities[len(m.entities)-1]
			m.entities = m.entities[:len(m.entities)-1]
		}
	}
}

func (m *Manager) Draw(emanager *entity.Manager, screen *ebiten.Image) {
	for _, e := range m.entities {

		if !emanager.Exists(e) {
			m.Remove(e)
			continue
		}

		op := &ebiten.DrawImageOptions{}

		// get our sprite
		com, ok := emanager.GetComponent(e, entity.ComponentSprite)
		if !ok {
			continue
		}

		image := com.(*sprite.Component).Sprite()

		// try get frame
		com, ok = emanager.GetComponent(e, entity.ComponentAnimator)
		if ok {
			frame := com.(*animator.Component)
			image = image.SubImage(frame.GetFrameRect()).(*ebiten.Image)
		}

		// get position
		com, ok = emanager.GetComponent(e, entity.ComponentPosition)
		if ok {
			vec := com.(*position.Component)
			op.GeoM.Translate(vec.X(), vec.Y())
		}

		screen.DrawImage(image, op)
	}
}
