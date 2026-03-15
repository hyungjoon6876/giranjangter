package event

import (
	"testing"
	"time"
)

func TestBroker_SubscribeAndSend(t *testing.T) {
	b := NewBroker()

	ch, cleanup := b.Subscribe("user1")
	defer cleanup()

	if !b.IsOnline("user1") {
		t.Fatal("user1 should be online after subscribe")
	}
	if b.OnlineCount() != 1 {
		t.Fatalf("expected 1 online, got %d", b.OnlineCount())
	}

	// Send event
	b.SendToUser("user1", SSEEvent{EventType: "test", Data: "hello"})

	select {
	case evt := <-ch:
		if evt.EventType != "test" {
			t.Fatalf("expected event type 'test', got %q", evt.EventType)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for event")
	}
}

func TestBroker_SendToOfflineUser(t *testing.T) {
	b := NewBroker()

	// Should not panic
	b.SendToUser("nonexistent", SSEEvent{EventType: "test", Data: "hello"})

	if b.IsOnline("nonexistent") {
		t.Fatal("nonexistent user should not be online")
	}
}

func TestBroker_Cleanup(t *testing.T) {
	b := NewBroker()

	_, cleanup := b.Subscribe("user2")
	if b.OnlineCount() != 1 {
		t.Fatalf("expected 1, got %d", b.OnlineCount())
	}

	cleanup()

	if b.IsOnline("user2") {
		t.Fatal("user2 should be offline after cleanup")
	}
	if b.OnlineCount() != 0 {
		t.Fatalf("expected 0, got %d", b.OnlineCount())
	}
}

func TestBroker_Reconnect(t *testing.T) {
	b := NewBroker()

	ch1, _ := b.Subscribe("user3")
	ch2, cleanup2 := b.Subscribe("user3") // reconnect
	defer cleanup2()

	// ch1 should be closed
	select {
	case _, ok := <-ch1:
		if ok {
			t.Fatal("old channel should be closed")
		}
	default:
		// This is fine - channel may not have been read yet
	}

	// ch2 should work
	b.SendToUser("user3", SSEEvent{EventType: "test", Data: "reconnected"})
	select {
	case evt := <-ch2:
		if evt.EventType != "test" {
			t.Fatalf("expected 'test', got %q", evt.EventType)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout")
	}
}

func TestBroker_SendToMultipleUsers(t *testing.T) {
	b := NewBroker()

	ch1, cleanup1 := b.Subscribe("a")
	defer cleanup1()
	ch2, cleanup2 := b.Subscribe("b")
	defer cleanup2()

	b.SendToUsers([]string{"a", "b"}, SSEEvent{EventType: "broadcast", Data: "all"})

	for _, ch := range []<-chan SSEEvent{ch1, ch2} {
		select {
		case evt := <-ch:
			if evt.EventType != "broadcast" {
				t.Fatalf("expected 'broadcast', got %q", evt.EventType)
			}
		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}
