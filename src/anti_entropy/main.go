package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shehio/anti-entropy/src/anti_entropy/node"
)

func createAndConnectNodes(count int) []*node.Node {
    nodes := make([]*node.Node, count)
    for i := 0; i < count; i++ {
        nodes[i] = node.NewNode(uint64(i + 1))
    }
    
    for _, n := range nodes {
        for _, peer := range nodes {
            if n.GetID() != peer.GetID() {
                n.AddPeer(peer)
            }
        }
    }
    return nodes
}

type weatherUpdate struct {
    key   string
    value string
}

func getWeatherUpdates() []weatherUpdate {
    return []weatherUpdate{
        {"weather", "sunny"},
        {"temperature", "25°C"},
        {"humidity", "65%"},
        {"wind_speed", "12 km/h"},
        {"pressure", "1013 hPa"},
    }
}

func shuffleNodes(nodes []*node.Node) []*node.Node {
	shuffled := make([]*node.Node, len(nodes))
	copy(shuffled, nodes)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

func main() {
	rand.Seed(time.Now().UnixNano())
	gossipingRounds := 3
	nodes := createAndConnectNodes(5)
    updates := getWeatherUpdates()

	fmt.Println("Starting anti-entropy protocol...")
	fmt.Println("Initial states:")
	for _, n := range nodes {
		fmt.Printf("Node %d: %v\n", n.GetID(), n.GetState())
	}

	for i, update := range updates {
		updatingNode := nodes[rand.Intn(len(nodes))]
		updatingNode.UpdateState(update.key, update.value)
        
		fmt.Printf("\nIteration %d:\n", i+1)
		fmt.Printf("Node %d updated %s to %s\n", updatingNode.GetID(), update.key, update.value)

		for round := 0; round < gossipingRounds; round++ {
			fmt.Printf("\nGossip Round %d:\n", round+1)
			
			shuffledNodes := shuffleNodes(nodes)
			for _, n := range shuffledNodes {
				if rand.Float32() < 0.8 { // 80% chance to gossip // todo: show information propagation with different probabilities
					n.Gossip() // todo: add the number of peers to gossip with
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