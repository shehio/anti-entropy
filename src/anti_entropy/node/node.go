package node

import (
	"fmt"
	"math/rand"
	"time"
)

// Node represents a node in the anti-entropy protocol
type Node struct {
    id       uint64
    state    map[string]string
    version  map[string]int
    peers    []uint64
    lastGossip time.Time
}

// NewNode creates a new node with the given ID
func NewNode(id uint64) *Node {
    node := &Node{
        id:       id,
        state:    make(map[string]string),
        version:  make(map[string]int),
        peers:    []uint64{2, 3},
        lastGossip: time.Now(),
    }
    // Initialize with some state
    node.state["food"] = "initial"
    node.version["food"] = 0
    return node
}

// PrintState displays the current state of the node
func (n *Node) PrintState() {
    fmt.Printf("\n=== Node %d State ===\n", n.id)
    if len(n.state) == 0 {
        fmt.Println("State is empty")
        return
    }
    for key, value := range n.state {
        fmt.Printf("Key: %s, Value: %s, Version: %d\n", key, value, n.version[key])
    }
    fmt.Println("==================")
}

// Gossip performs a gossip operation with a random peer
func (n *Node) Gossip() {
    if time.Since(n.lastGossip) <= 100*time.Millisecond {
        return
    }
    n.lastGossip = time.Now()
    peer := n.peers[rand.Intn(len(n.peers))]
    
    fmt.Printf("\nNode %d gossiping with Node %d\n", n.id, peer)
    
    // Simulate receiving state from peer
    if rand.Float32() < 0.5 { // 50% chance of peer having different state
        n.state["food"] = "pizza"
        n.version["food"]++
        fmt.Printf("Node %d updated key food to value pizza (version %d)\n", n.id, n.version["food"])
        n.PrintState()
    }
}

// UpdateState updates the state with a new value
func (n *Node) UpdateState(key string, value string) {
    n.state[key] = value
    n.version[key]++
    fmt.Printf("\nNode %d manually updated key %s to value %s (version %d)\n", n.id, key, value, n.version[key])
    n.PrintState()
} 