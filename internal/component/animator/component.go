package animator

import (
	"image"
	"time"

	"github.com/jawr/castaway/internal/entity"
)

type Component struct {
	entity.Entity

	frameWidth, frameHeight int
	currentRow, currentCol  int
	totalRows, totalCols    int

	speed       time.Duration
	lastFrameAt time.Time
}

func NewComponent(e entity.Entity, frameWidth, frameHeight, totalRows, totalCols int, speed time.Duration) *Component {
	return &Component{
		Entity:      e,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
		totalRows:   totalRows,
		totalCols:   totalCols,
		speed:       speed,
	}
}

func (c *Component) GetEntity() entity.Entity      { return c.Entity }
func (c *Component) GetFlag() entity.ComponentFlag { return entity.ComponentAnimator }

func (c *Component) NextFrame() {

	if time.Since(c.lastFrameAt) < c.speed {
		return
	}
	c.lastFrameAt = time.Now()

	if c.currentCol+1 >= c.totalCols {
		c.currentCol = 0
	} else {
		c.currentCol++
	}
}

func (c *Component) SetRow(row int) {
	c.currentRow = row
}

func (c *Component) GetFrameRect() image.Rectangle {
	sx := c.currentCol * c.frameWidth
	sy := c.currentRow * c.frameHeight

	return image.Rect(sx, sy, sx+c.frameWidth, sy+c.frameHeight)
}
