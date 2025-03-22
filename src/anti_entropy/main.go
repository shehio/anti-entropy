package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shehio/anti-entropy/src/anti_entropy/node"
)

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	node1 := node.NewNode(1)
	node2 := node.NewNode(2)
	node3 := node.NewNode(3)
	node4 := node.NewNode(4)
	node5 := node.NewNode(5)

	nodes := []*node.Node{node1, node2, node3, node4, node5}
	for _, n := range nodes {
		for _, peer := range nodes {
			if n.GetID() != peer.GetID() {
				n.AddPeer(peer)
			}
		}
	}

	fmt.Println("Starting anti-entropy protocol...")
	fmt.Println("Initial states:")
	for _, n := range nodes {
		fmt.Printf("Node %d: %v\n", n.GetID(), n.GetState())
	}

	updates := []struct {
		key   string
		value string
	}{
		{"weather", "sunny"},
		{"temperature", "25°C"},
		{"humidity", "65%"},
		{"wind_speed", "12 km/h"},
		{"pressure", "1013 hPa"},
	}

	for i, update := range updates {
		updatingNode := nodes[rand.Intn(len(nodes))]
		updatingNode.UpdateState(update.key, update.value)
		fmt.Printf("\nIteration %d:\n", i+1)
		fmt.Printf("Node %d updated %s to %s\n", updatingNode.GetID(), update.key, update.value)

		numRounds := 3
		for round := 0; round < numRounds; round++ {
			fmt.Printf("\nGossip Round %d:\n", round+1)
			
			shuffledNodes := make([]*node.Node, len(nodes))
			copy(shuffledNodes, nodes)
			rand.Shuffle(len(shuffledNodes), func(i, j int) {
				shuffledNodes[i], shuffledNodes[j] = shuffledNodes[j], shuffledNodes[i]
			})

			for _, n := range shuffledNodes {
				if rand.Float32() < 0.8 { // 80% chance to gossip
					n.Gossip()
				}
			}

			fmt.Println("\nCurrent states:")
			for _, n := range nodes {
				fmt.Printf("Node %d: %v\n", n.GetID(), n.GetState())
			}

			time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
		}
	}
} 