package node

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
    id       uint64
    state    map[string]string
    version  map[string]int
    knownPeers    []*Node
    lastGossip time.Time
}

func NewNode(id uint64) *Node {
    node := &Node{
        id:       id,
        state:    make(map[string]string),
        version:  make(map[string]int),
        knownPeers:    make([]*Node, 0),
        lastGossip: time.Now(),
    }

    node.state["weather"] = "unknown"
    node.version["weather"] = 0
    return node
}

func (n *Node) AddPeer(peer *Node) {
    n.knownPeers = append(n.knownPeers, peer)
}

func (n *Node) GetID() uint64 {
    return n.id
}

func (n *Node) GetState() map[string]string {
    return n.state
}

func (n *Node) GetVersion() map[string]int {
    return n.version
}

func (n *Node) UpdateState(key, value string) {
    n.state[key] = value
    n.version[key]++
}

func (n *Node) Gossip() {
    if time.Since(n.lastGossip) < time.Second {
        return
    }
    n.lastGossip = time.Now()
    if len(n.knownPeers) == 0 {
        return
    }
    peer := n.knownPeers[rand.Intn(len(n.knownPeers))]
    
    fmt.Printf("\nNode %d gossiping with Node %d\n", n.id, peer.id)
    fmt.Printf("Node %d sending state: %v (versions: %v)\n", n.id, n.state, n.version)

    // Exchange states with peer
    for key, value := range peer.state {
        if n.version[key] < peer.version[key] {
            n.state[key] = value
            n.version[key] = peer.version[key]
            fmt.Printf("Node %d updated key %s to value %s (version %d)\n", n.id, key, value, n.version[key])
        }
    }

    // Also let peer learn from our state
    for key, value := range n.state {
        if peer.version[key] < n.version[key] {
            peer.state[key] = value
            peer.version[key] = n.version[key]
            fmt.Printf("Node %d updated key %s to value %s (version %d)\n", peer.id, key, value, peer.version[key])
        }
    }
}

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
