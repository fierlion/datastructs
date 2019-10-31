package board

import (
    "fmt"
    "github.com/fierlion/datastructs/shape"
)

type Board struct {
    state [][]*shape.Shape
}

func New(size int) *Board {
    b := new(Board)
    b.state = make([][]*shape.Shape, size)
    for i := range b.state {
        b.state[i] = make([]*shape.Shape, size)
    }
    return b
}

func (b Board) Add(x, y int, s *shape.Shape) {
    b.state[x][y] = s
}

func (b Board) Print() {
    for i := range b.state {
        fmt.Printf("Row: %v\n", i)
        for j := range b.state[i] {
            fmt.Printf("%v\n", b.state[i][j])
        }
    }
}
