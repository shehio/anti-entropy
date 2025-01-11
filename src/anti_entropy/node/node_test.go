package node

import (
	"testing"
	"time"
)

func TestNewNode(t *testing.T) {
	node := NewNode(1)
	if node.id != 1 {
		t.Errorf("Expected node ID 1, got %d", node.id)
	}
	if node.state["food"] != "initial" {
		t.Errorf("Expected initial food value 'initial', got '%s'", node.state["food"])
	}
	if node.version["food"] != 0 {
		t.Errorf("Expected initial version 0, got %d", node.version["food"])
	}
}

func TestUpdateState(t *testing.T) {
	node := NewNode(1)
	
	// Test updating state
	node.UpdateState("food", "pizza")
	if node.state["food"] != "pizza" {
		t.Errorf("Expected food value 'pizza', got '%s'", node.state["food"])
	}
	if node.version["food"] != 1 {
		t.Errorf("Expected version 1, got %d", node.version["food"])
	}
	
	// Test updating multiple times
	node.UpdateState("food", "sushi")
	if node.state["food"] != "sushi" {
		t.Errorf("Expected food value 'sushi', got '%s'", node.state["food"])
	}
	if node.version["food"] != 2 {
		t.Errorf("Expected version 2, got %d", node.version["food"])
	}
}

func TestGossip(t *testing.T) {
	node := NewNode(1)
	
	// Test initial gossip
	node.Gossip()
	
	// Test gossip cooldown
	node.Gossip() // Should not update due to cooldown
	if node.state["food"] != "initial" {
		t.Errorf("Expected food value to remain 'initial' during cooldown, got '%s'", node.state["food"])
	}
	
	// Wait for cooldown and test again
	time.Sleep(200 * time.Millisecond)
	node.Gossip()
	
	// Note: We can't test the exact state after gossip because it's randomized
	// We can only verify that the version might have increased
	if node.version["food"] > 0 {
		t.Logf("Gossip updated version to %d", node.version["food"])
	}
}

func TestMultipleKeys(t *testing.T) {
	node := NewNode(1)
	
	// Test updating multiple keys
	node.UpdateState("food", "pizza")
	node.UpdateState("drink", "coffee")
	
	if node.state["food"] != "pizza" {
		t.Errorf("Expected food value 'pizza', got '%s'", node.state["food"])
	}
	if node.state["drink"] != "coffee" {
		t.Errorf("Expected drink value 'coffee', got '%s'", node.state["drink"])
	}
	if node.version["food"] != 1 {
		t.Errorf("Expected food version 1, got %d", node.version["food"])
	}
	if node.version["drink"] != 1 {
		t.Errorf("Expected drink version 1, got %d", node.version["drink"])
	}
} 