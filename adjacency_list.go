package graph

import (
	"bytes"
	"strconv"
)

type adjacencyList struct {
	nodes map[uint32][]uint32
}

func newAdjacencyList(size uint32) *adjacencyList {
	return &adjacencyList{
		nodes: make(map[uint32][]uint32, size),
	}
}

func (l *adjacencyList) adjacent(n uint32) []uint32 {
	return l.nodes[n]
}

func (l *adjacencyList) connect(n1, n2 uint32) {
	l.nodes[n1] = append(l.nodes[n1], n2)
	l.nodes[n2] = append(l.nodes[n2], n1)
}

func (l *adjacencyList) String() string {
	var buffer bytes.Buffer
	for k := range l.nodes {
		buffer.WriteString(strconv.FormatUint(uint64(k), 10))
		buffer.WriteString(": ")
		for _, n := range l.nodes[k] {
			buffer.WriteString(strconv.FormatUint(uint64(n), 10))
			buffer.WriteRune(' ')
		}
		buffer.WriteRune('\n')
	}
	return buffer.String()
}
