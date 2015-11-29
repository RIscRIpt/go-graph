package graph

import (
	"bytes"

	"github.com/RIscRIpt/bitarray"
)

type adjacencyMatrix struct {
	size   uint32
	matrix *bitarray.BitArray
}

func newAdjacencyMatrix(size uint32) (m *adjacencyMatrix) {
	m = &adjacencyMatrix{}
	m.size = size
	m.matrix = bitarray.NewBitArray(size*size, 1)
	return
}

func (m *adjacencyMatrix) adjacent(n uint32) (adj []uint32) {
	beg := m.size * n
	end := beg + m.size
	for i := beg; i < end; i++ {
		if m.matrix.GetB(i) != 0 {
			adj = append(adj, i-beg)
		}
	}
	return
}

func (m *adjacencyMatrix) connect(n1, n2 uint32) {
	if n1 >= m.size || n2 >= m.size {
		panic("Out of range")
	}
	m.matrix.SetB(n1*m.size+n2, 1)
	m.matrix.SetB(n2*m.size+n1, 1)
}

func (m *adjacencyMatrix) String() string {
	var buffer bytes.Buffer
	i := uint32(0)
	e := m.size * m.size
	for i < e {
		if m.matrix.GetB(i) != 0 {
			buffer.WriteString("1 ")
		} else {
			buffer.WriteString("0 ")
		}
		i++
		if i%m.size == 0 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}
