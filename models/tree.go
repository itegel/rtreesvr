package models

import (
	"github.com/dhconnelly/rtreego"
	"sync"
)

var DTree *DBufferdTree

type Tree struct {
	Rtree       *rtreego.Rtree
	Mutex       sync.RWMutex //rwlock for modification of the rtree is not thread safe
	dim         int
	minChildren int
	maxChildren int
}

type DBufferdTree struct {
	trees   [2]*Tree
	working int
}

func init() {
	DTree = nil
}

func NewDTree(MinChildren, MaxChildren int) *DBufferdTree {
	DTree = &DBufferdTree{}
	DTree.trees[0] = NewTree(MinChildren, MaxChildren)
	DTree.trees[1] = NewTree(MinChildren, MaxChildren)
	DTree.working = 0 //init working buffer
	return DTree
}

func NewTree(MinChildren, MaxChildren int) *Tree {
	tree := Tree{}
	tree.Rtree = rtreego.NewTree(2, MinChildren, MaxChildren)
	tree.Mutex = sync.RWMutex{}
	tree.dim = 2 //only 2-dim is supported
	tree.minChildren = MinChildren
	tree.maxChildren = MaxChildren

	return &tree
}

//insert a node to tree
func (t *Tree) Insert(node *Node) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Rtree.Insert(node)
}

func (dt *DBufferdTree) Insert(node *Node) {
	//insert into idle buffer
	insert_index := (dt.working + 1) % 2
	dt.trees[insert_index].Insert(node)

	//switch working buffer
	dt.working = insert_index

	//insert into current idle buffer
	insert_index = (dt.working + 1) % 2
	dt.trees[insert_index].Insert(node)
}

//search all nodes which intersect with [(x,y),(x+w,y+h)]
//if you need rlock you should set need_lock to true,
//it always works without lock in many case
func (t *Tree) SearchIntersect(x, y, w, h float64, need_lock bool) []*Node {
	if need_lock {
		t.Mutex.RLock()
		defer t.Mutex.RUnlock()
	}
	p := rtreego.Point{x, y}
	r, ok := rtreego.NewRect(p, []float64{w, h})
	if ok != nil {
		return nil
	}
	results := t.Rtree.SearchIntersect(r)

	nodes := []*Node{}
	for _, res := range results {
		if node, ok := res.(*Node); ok {
			///TODO: filters here
			nodes = append(nodes, node)
		} else {
			panic("convert result to *Node fail")
		}
	}
	return nodes
}

func (dt *DBufferdTree) SearchIntersect(x, y, w, h float64, need_lock bool) []*Node {
	return dt.trees[dt.working].SearchIntersect(x, y, w, h, need_lock)
}

///TODO: deletes here
