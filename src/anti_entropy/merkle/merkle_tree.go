package merkle

import (
	"fmt"
)

type MerkleTree struct {
	Root     *MerkleNode
	Leaves   []*MerkleNode
	leafMap  map[string]int
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	if len(data) == 0 {
		return nil
	}

	leaves := make([]*MerkleNode, len(data))
	leafMap := make(map[string]int)
	for i, d := range data {
		hash := calculateHash(d)
		leaves[i] = NewLeafNode(d)
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
			parent = NewParentNode(leaves[i], leaves[i+1])
		} else {
			parent = NewParentNode(leaves[i], nil)
		}
		parents = append(parents, parent)
	}

	return buildTree(parents)
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
	proof := make([]string, 0)
	if !buildProof(t.Root, targetHash, &proof) {
		return nil
	}
	return proof
}

func buildProof(node *MerkleNode, targetHash string, proof *[]string) bool {
	if node == nil {
		return false
	}

	if node.IsLeaf {
		return node.Hash == targetHash
	}

	// Try left subtree
	if buildProof(node.Left, targetHash, proof) {
		if node.Right != nil {
			*proof = append(*proof, node.Right.Hash)
		}
		return true
	}

	// Try right subtree
	if buildProof(node.Right, targetHash, proof) {
		if node.Left != nil {
			*proof = append(*proof, node.Left.Hash)
		}
		return true
	}

	return false
} 