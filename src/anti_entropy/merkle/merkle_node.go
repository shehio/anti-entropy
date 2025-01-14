package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

type MerkleNode struct {
	Hash     string
	Left     *MerkleNode
	Right    *MerkleNode
	Data     []byte
	IsLeaf   bool
}

func calculateHash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func NewLeafNode(data []byte) *MerkleNode {
	return &MerkleNode{
		Hash:   calculateHash(data),
		Data:   data,
		IsLeaf: true,
	}
}

func NewParentNode(left, right *MerkleNode) *MerkleNode {
	var hashData []byte
	if right != nil {
		hashData = append([]byte(left.Hash), []byte(right.Hash)...)
	} else {
		hashData = []byte(left.Hash)
	}

	return &MerkleNode{
		Hash:   calculateHash(hashData),
		Left:   left,
		Right:  right,
		IsLeaf: false,
	}
}

func (n *MerkleNode) GetSibling() *MerkleNode {
	if n == nil {
		return nil
	}
	if n == n.Left {
		return n.Right
	}
	return n.Left
} 