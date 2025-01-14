package merkle

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type MerkleNode struct {
	Hash     string
	Left     *MerkleNode
	Right    *MerkleNode
	Data     []byte
	IsLeaf   bool
}

type MerkleTree struct {
	Root   *MerkleNode
	Leaves []*MerkleNode
	leafMap map[string]int
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	if len(data) == 0 {
		return nil
	}

	leaves := make([]*MerkleNode, len(data))
	leafMap := make(map[string]int)
	for i, d := range data {
		hash := calculateHash(d)
		leaves[i] = &MerkleNode{
			Hash:   hash,
			Data:   d,
			IsLeaf: true,
		}
		leafMap[hash] = i
	}

	root := buildTree(leaves)

	return &MerkleTree{
		Root:     root,
		Leaves:   leaves,
		leafMap:  leafMap,
	}
}

func buildTree(leaves []*MerkleNode) *MerkleNode {
	if len(leaves) == 0 {
		return nil
	}
	if len(leaves) == 1 {
		return leaves[0]
	}

	parents := make([]*MerkleNode, 0, (len(leaves)+1)/2)
	for i := 0; i < len(leaves); i += 2 {
		var parent *MerkleNode
		if i+1 < len(leaves) {
			parent = &MerkleNode{
				Left:  leaves[i],
				Right: leaves[i+1],
				Hash:  calculateHash(append([]byte(leaves[i].Hash), []byte(leaves[i+1].Hash)...)),
			}
		} else {
			parent = &MerkleNode{
				Left:  leaves[i],
				Hash:  calculateHash([]byte(leaves[i].Hash)),
			}
		}
		parents = append(parents, parent)
	}

	return buildTree(parents)
}

func calculateHash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func (t *MerkleTree) Verify(data []byte) bool {
	if t == nil || t.Root == nil {
		return false
	}

	targetHash := calculateHash(data)
	return verifyNode(t.Root, targetHash)
}

func verifyNode(node *MerkleNode, targetHash string) bool {
	if node == nil {
		return false
	}

	if node.IsLeaf {
		return node.Hash == targetHash
	}

	return verifyNode(node.Left, targetHash) || verifyNode(node.Right, targetHash)
}

func (t *MerkleTree) GetRootHash() string {
	if t == nil || t.Root == nil {
		return ""
	}
	return t.Root.Hash
}

func (t *MerkleTree) String() string {
	if t == nil || t.Root == nil {
		return "Empty Tree"
	}
	return fmt.Sprintf("MerkleTree(Root: %s)", t.Root.Hash)
}

func (t *MerkleTree) GetProof(data []byte) []string {
	if t == nil || t.Root == nil {
		return nil
	}

	targetHash := calculateHash(data)
	leafIndex, exists := t.leafMap[targetHash]
	if !exists {
		return nil
	}

	proof := make([]string, 0)
	currentIndex := leafIndex
	for currentIndex > 0 {
		siblingIndex := currentIndex + 1
		if currentIndex%2 == 0 {
			siblingIndex = currentIndex - 1
		}
		if siblingIndex < len(t.Leaves) {
			proof = append(proof, t.Leaves[siblingIndex].Hash)
		}
		currentIndex = (currentIndex - 1) / 2
	}

	return proof
}