package models

import (
	"github.com/dhconnelly/rtreego"
)

//rtree node who stores the location info and user defined infos
type Node struct {
	Loc *rtreego.Rect
	Id  int64
	Ext string
}

//should implement Bounds interface
func (n *Node) Bounds() *rtreego.Rect {
	return n.Loc
}

//only two dim supported
func NewNode(x, y, w, h float64, id int64, ext string) *Node {
	p := rtreego.Point{x, y}
	r, ok := rtreego.NewRect(p, []float64{w, h})
	if ok != nil {
		return nil
	}
	return &Node{r, id, ext}

}
