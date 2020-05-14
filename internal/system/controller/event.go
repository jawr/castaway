package controller

import (
	"github.com/jawr/castaway/internal/entity"
	"github.com/jawr/castaway/internal/event"
	"github.com/peterhellberg/gfx"
)

type EventMove struct {
	Entity entity.Entity
	Origin gfx.Vec
	New    gfx.Vec
}

func (e *EventMove) Topic() event.Topic { return event.TopicMove }
