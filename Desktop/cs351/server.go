package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Move struct {
	Color int
	Col   int
}

type Board struct {
	BoardString string
}

var previousColor int = -1
type ConnectGame int
func (t *ConnectGame) Move(args *Move, reply *int) error {
	log.Println("Server received a Move request for", args.Color, "for column", args.Col)
	currentColor := args.Color
	if currentColor == previousColor {
		return errors.New("Turn Order Violation")
	}
	previousColor = currentColor
	return nil
}

func (t *ConnectGame) Get(args *int, reply *Board) error {
	reply.BoardString = "Hello World"
	return nil
}

const rows = 6
const cols = 7

var gameBoard [][]int 

func main() {
	gameBoard = make([][]int, rows)

	for i := range gameBoard {
		gameBoard[i] = make([]int, cols)
	}
	cg := new(ConnectGame)
	rpc.Register(cg)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

 

	log.Println("Serving on PORT 1234")
	http.Serve(l, nil)
}
