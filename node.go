package dumbo

import "net/http"

 
type edge struct {
	key string
	n   *Node
}


type Node struct {
	handler  http.HandlerFunc
	edges    []*edge
	priority int
	depth    int
}


func (n *Node) IsLeaf() bool {
	length := len(n.edges)
	if length == 2 {
		return n.edges[0] == nil && n.edges[1] == nil
	}
	return length == 0
}


func (n *Node) incrDepth() {
	n.depth++
	for _, e := range n.edges {
		e.n.incrDepth()
	}
}


func (n *Node) clone() *Node {
	c := *n
	c.incrDepth()
	return &c
}