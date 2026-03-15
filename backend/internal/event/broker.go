package event

import (
	"log"
	"sync"
)

// SSEEvent represents a server-sent event.
type SSEEvent struct {
	EventType string      `json:"eventType"`
	Data      interface{} `json:"data"`
}

// Broker manages SSE client connections and broadcasts events.
type Broker struct {
	mu      sync.RWMutex
	clients map[string]chan SSEEvent // userId → event channel
}

// NewBroker creates a new SSE broker.
func NewBroker() *Broker {
	return &Broker{
		clients: make(map[string]chan SSEEvent),
	}
}

// Subscribe registers a user for SSE events. Returns a channel and cleanup func.
func (b *Broker) Subscribe(userID string) (<-chan SSEEvent, func()) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Close existing channel if reconnecting
	if old, ok := b.clients[userID]; ok {
		close(old)
	}

	ch := make(chan SSEEvent, 64)
	b.clients[userID] = ch

	cleanup := func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		if current, ok := b.clients[userID]; ok && current == ch {
			delete(b.clients, userID)
			close(ch)
		}
	}

	return ch, cleanup
}

// SendToUser sends an event to a specific user.
func (b *Broker) SendToUser(userID string, evt SSEEvent) {
	b.mu.RLock()
	ch, ok := b.clients[userID]
	b.mu.RUnlock()

	if !ok {
		return // User not connected
	}

	select {
	case ch <- evt:
	default:
		log.Printf("SSE: dropping event for user %s (buffer full)", userID)
	}
}

// SendToUsers sends an event to multiple users.
func (b *Broker) SendToUsers(userIDs []string, evt SSEEvent) {
	for _, uid := range userIDs {
		b.SendToUser(uid, evt)
	}
}

// IsOnline checks if a user has an active SSE connection.
func (b *Broker) IsOnline(userID string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	_, ok := b.clients[userID]
	return ok
}

// OnlineCount returns the number of connected clients.
func (b *Broker) OnlineCount() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.clients)
}

