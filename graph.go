package graph

import (
	"github.com/RIscRIpt/bitarray"
	"github.com/RIscRIpt/dllist"
)

type NodeHolder uint32

const (
	HT_None NodeHolder = iota
	HT_AdjacencyList
	HT_AdjacencyMatrix
)

type Graph struct {
	Root   *Node
	holder nodeHolder
	nodes  map[uint32]*Node
}

type nodeHolder interface {
	adjacent(n uint32) []uint32
	connect(n1, n2 uint32)
	String() string
}

func NewGraph(ht NodeHolder, initialSize uint32) (g *Graph) {
	g = &Graph{}
	switch ht {
	case HT_AdjacencyList:
		g.holder = newAdjacencyList(initialSize)
	case HT_AdjacencyMatrix:
		g.holder = newAdjacencyMatrix(initialSize)
	default:
		panic("Unsupported node holder type")
	}
	g.nodes = make(map[uint32]*Node)
	return
}

func (g *Graph) AddNode(value interface{}) uint32 {
	node := NewNode(uint32(len(g.nodes)), value)
	if g.Root == nil {
		g.Root = node
	}
	g.nodes[node.id] = node
	return node.id
}

func (g *Graph) AdjacencyList() (list map[*Node][]*Node) {
	list = make(map[*Node][]*Node)
	for _, n := range g.nodes {
		list[n] = append(list[n], g.Adjacent(n)...)
	}
	return
}

func (g *Graph) Adjacent(n *Node) (adj []*Node) {
	for _, i := range g.holder.adjacent(n.id) {
		adj = append(adj, g.nodes[i])
	}
	return
}

func (g *Graph) Connect(n1, n2 uint32) {
	g.holder.connect(n1, n2)
}

func (g *Graph) String() string {
	return g.holder.String()
}

func (g *Graph) ConnectNodes(n1, n2 *Node) {
	g.Connect(n1.id, n2.id)
	g.nodes[n1.id] = n1
	g.nodes[n2.id] = n2
}

func (g *Graph) BFS() (path []*Node) {
	visited := bitarray.NewBitArray(uint32(len(g.nodes)), 1)
	queue := dllist.New()
	queue.PushBack(g.Root)
	for queue.Length() > 0 {
		m, err := queue.PopFront()
		if err != nil {
			break
		}
		n := m.(*Node)
		if visited.GetB(n.id) == 0 {
			visited.SetB(n.id, 1)
			path = append(path, n)
			for _, c := range g.Adjacent(n) {
				queue.PushBack(c)
			}
		}
	}
	return
}

func (g *Graph) DFS() (path []*Node) {
	visited := bitarray.NewBitArray(uint32(len(g.nodes)), 1)
	stack := dllist.New()
	stack.PushFront(g.Root)
	for stack.Length() > 0 {
		m, err := stack.PopFront()
		if err != nil {
			panic(err)
		}
		n := m.(*Node)
		if visited.GetB(n.id) == 0 {
			path = append(path, n)
			visited.SetB(n.id, 1)
			for _, c := range g.Adjacent(n) {
				stack.PushFront(c)
			}
		}
	}
	return
}
