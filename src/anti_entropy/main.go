package main

import (
	"fmt"
	"time"
	"github.com/shehio/anti-entropy/src/anti_entropy/node"
)

func main() {
    node := node.NewNode(1)
    fmt.Println("Starting anti-entropy protocol...")
    node.PrintState()
    
    foods := []string{"olives", "harissa", "shawerma", "baklava"}
    for i := 0; i < 6; i++ {
        // Update state
        node.UpdateState("food", foods[i%len(foods)])
        time.Sleep(500 * time.Millisecond)
        
        // Gossip
        node.Gossip()
        time.Sleep(500 * time.Millisecond)
    }
    
    fmt.Println("\nProtocol completed!")
} 