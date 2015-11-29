package graph

type Node struct {
	Value interface{}
	id    uint32
}

func NewNode(id uint32, value interface{}) *Node {
	return &Node{
		Value: value,
		id:    id,
	}
}

func (n Node) Less(other interface{}) bool {
	return n.id < other.(Node).id
}

func (n Node) Equals(other interface{}) bool {
	return n.id == other.(Node).id
}
