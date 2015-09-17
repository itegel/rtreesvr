package models

import (
	"testing"
)

func TestNode(t *testing.T) {
	var node *Node
	node = NewNode(1, 1, 2, 2, 1000, "hello")
	t.Log(*node)
}
