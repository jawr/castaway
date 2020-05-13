package sprite

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jawr/castaway/internal/entity"
	"github.com/pkg/errors"
)

type Component struct {
	entity.Entity

	sprite *ebiten.Image
}

func NewComponent(e entity.Entity, path string) *Component {
	sprite, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		panic(errors.Wrapf(err, "Unable to open sprite '%s'", path))
	}

	return &Component{
		Entity: e,
		sprite: sprite,
	}
}

func (c *Component) GetEntity() entity.Entity { return c.Entity }
func (c *Component) Sprite() *ebiten.Image    { return c.sprite }
