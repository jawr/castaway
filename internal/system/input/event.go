package input

import "github.com/jawr/castaway/internal/event"

type Direction int

const (
	DirectionDown Direction = iota
	DirectionRight
	DirectionLeft
	DirectionUp
)

type EventDirectionKey struct {
	Direction Direction
}

func (e *EventDirectionKey) Topic() event.Topic { return event.TopicDirectionKeyPressed }
