package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jawr/castaway/internal/component/animator"
	"github.com/jawr/castaway/internal/component/position"
	"github.com/jawr/castaway/internal/component/sprite"
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
)

type Manager struct {
	entities []entity.Entity

	cam *camera
}

func NewManager(width, height int) *Manager {
	return &Manager{
		entities: make([]entity.Entity, 0),
		cam:      newCamera(width, height),
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

func (m *Manager) SetupSubscriptions(emanager *entity.Manager, publisher event.Publisher, subscriber event.Subscriber) {
	subscriber(event.TopicMove, m.cam.handleMove())
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
			op.GeoM.Translate(vec.X, vec.Y)
		}

		op.GeoM.Translate(-m.cam.X, -m.cam.Y)

		screen.DrawImage(image, op)
	}

	// draw our bound
	bound, err := ebiten.NewImage(int(m.cam.boundWidth), int(m.cam.boundHeight), ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.cam.bound.X, m.cam.bound.Y)
	screen.DrawImage(bound, op)

	ebitenutil.DrawRect(screen, m.cam.bound.X-m.cam.X, m.cam.bound.Y-m.cam.Y, m.cam.boundWidth, m.cam.boundHeight, color.NRGBA{0xff, 0x00, 0x00, 0x11})
}
