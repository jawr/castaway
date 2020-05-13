package event

import "sync"

type Topic uint8

type Event interface {
	Topic() Topic
}

// External callable functions
type Publisher func(Event)
type Subscriber func(Topic, Subscription)

// A callback function that is called with an Event for an Topic
type Subscription chan Event

type Manager struct {
	// map of Subscriptions for each Topic
	subscriptions    map[Topic][]Subscription
	subscriptionsMtx *sync.RWMutex
}

// Create a new Event Manager
func NewManager() *Manager {
	return &Manager{
		subscriptions:    make(map[Topic][]Subscription, 0),
		subscriptionsMtx: &sync.RWMutex{},
	}
}

// Handle published events
func (m *Manager) Publish(e Event) {
	m.subscriptionsMtx.RLock()
	if subs, ok := m.subscriptions[e.Topic()]; ok {
		// take a copy of our subscriptions so we can drop
		// the lock quicker, might want to remove this if
		// we care about memory usage
		chans := append([]Subscription{}, subs...)
		go func(chans []Subscription) {
			for _, ch := range subs {
				ch <- e
			}
		}(chans)
	}
	m.subscriptionsMtx.RUnlock()
}

// Caller can subscribe to specific Topic
func (m *Manager) Subscribe(topic Topic, ch Subscription) {
	m.subscriptionsMtx.Lock()
	defer m.subscriptionsMtx.Unlock()

	if _, ok := m.subscriptions[topic]; !ok {
		m.subscriptions[topic] = make([]Subscription, 0)
	}

	m.subscriptions[topic] = append(m.subscriptions[topic], ch)
}
