package merkle

import (
	"testing"
)

func TestEmptyTree(t *testing.T) {
	tree := NewMerkleTree(nil)
	if tree != nil {
		t.Error("Expected nil tree for empty data")
	}
}

func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
		[]byte("data4"),
	}

	tree := NewMerkleTree(data)
	if tree == nil {
		t.Error("Expected non-nil tree")
	}
	if tree.Root == nil {
		t.Error("Expected non-nil root")
	}
	if len(tree.Leaves) != len(data) {
		t.Errorf("Expected %d leaves, got %d", len(data), len(tree.Leaves))
	}
}

func TestOddNumberOfLeaves(t *testing.T) {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
	}

	tree := NewMerkleTree(data)
	if tree == nil {
		t.Error("Expected non-nil tree")
	}
	if tree.Root == nil {
		t.Error("Expected non-nil root")
	}
	if len(tree.Leaves) != len(data) {
		t.Errorf("Expected %d leaves, got %d", len(data), len(tree.Leaves))
	}
}

func TestSingleLeaf(t *testing.T) {
	data := [][]byte{[]byte("data1")}
	tree := NewMerkleTree(data)
	if tree == nil {
		t.Error("Expected non-nil tree")
	}
	if tree.Root == nil {
		t.Error("Expected non-nil root")
	}
	if !tree.Root.IsLeaf {
		t.Error("Expected root to be a leaf")
	}
	if len(tree.Leaves) != 1 {
		t.Error("Expected one leaf")
	}
} 

func TestVerify(t *testing.T) {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
		[]byte("data4"),
	}

	tree := NewMerkleTree(data)

	if !tree.Verify([]byte("data1")) {
		t.Error("Expected data1 to be verified")
	}

	if tree.Verify([]byte("nonexistent")) {
		t.Error("Expected nonexistent data to not be verified")
	}
}

func TestGetProof(t *testing.T) {
	data := [][]byte{
		[]byte("data1"),
		[]byte("data2"),
		[]byte("data3"),
		[]byte("data4"),
	}

	tree := NewMerkleTree(data)

	proof := tree.GetProof([]byte("data1"))
	if proof == nil {
		t.Error("Expected non-nil proof")
	}
	if len(proof) == 0 {
		t.Error("Expected non-empty proof")
	}

	proof = tree.GetProof([]byte("nonexistent"))
	if proof != nil {
		t.Error("Expected nil proof for nonexistent data")
	}
}