package main

import (
    "github.com/fierlion/datastructs/shape"
    "github.com/fierlion/datastructs/board"
)

const boardSize = 10

func main() {
    testShape1 := &shape.Shape{U : shape.Green}
    testShape2 := &shape.Shape{D : shape.Blue}
    //testShape3 := &shape.Shape{R : shape.Red}
    //testShape4 := &shape.Shape{L : shape.Green}

    myBoard := board.New(boardSize)
    myBoard.Add(0, 1, testShape1)
    myBoard.Add(1, 0, testShape2)
    //myBoard.Add(9, 9, testShape3)
    //myBoard.Add(5, 5, testShape4)
    myBoard.Print()
}
