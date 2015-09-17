package models

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	var tree *Tree
	tree = NewTree(25, 25)
	var n1, n2 *Node
	n1 = NewNode(1, 1, 2, 2, 1001, "airport")
	n2 = NewNode(4, 4, 2, 2, 1002, "train station")

	tree.Insert(n1)
	tree.Insert(n2)
	//t.Log(n1, n2, tree)

	nodes := tree.SearchIntersect(2.5, 2.5, 1, 1, true)
	if len(nodes) != 1 || nodes[0].Id != 1001 {
		t.Error("search failed!")
	}

	nodes = tree.SearchIntersect(2.5, 2.5, 3, 3, false)
	if len(nodes) != 2 {
		t.Error("search failed!")
	}
}

func TestDTree(t *testing.T) {
	NewDTree(30, 30)
	t.Log(DTree.working)

	var n1, n2, n3 *Node
	n1 = NewNode(1, 1, 2, 2, 1001, "airport")
	n2 = NewNode(4, 4, 2, 2, 1002, "train station")
	DTree.Insert(n1)
	t.Log(DTree.working)
	DTree.Insert(n2)
	t.Log(DTree.working)
	nodes := DTree.SearchIntersect(2.5, 2.5, 1, 1, true)
	if len(nodes) != 1 || nodes[0].Id != 1001 {
		t.Error("search failed!")
	}

	n3 = NewNode(3, 3, 2, 2, 1003, "train station")
	DTree.Insert(n3)
	t.Log(DTree.working)
	nodes = DTree.SearchIntersect(2.5, 2.5, 1, 1, true)
	if len(nodes) != 2 {
		t.Error("search failed!")
	}
}
