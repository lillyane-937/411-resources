package main

import (
	"log"
	"net/rpc"
	"time"
)

type Move struct {
	Color int
	Col   int
}

type Board struct {
	BoardString string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()

	ch := make(chan int)
	for i := range 10 {
		go func() {
			var reply int
			moveWhite := Move{0, i % 5}

			// log.Println("Client is sending a White color move")
			errWhite := client.Call("ConnectGame.Move", &moveWhite, &reply)
			if errWhite != nil {
				log.Println("RPC error:", errWhite)
				ch <- 0
			} else {
				ch <- 1
			}
		}()
	}
	sum := 0
	for range 10 {
		// receive on channela nd update sum correctly to count succesful moves
		val := <-ch
		sum += val
	}
	var replyB int
	moveBlack := Move{1, 0}

	// log.Println("Client is sending a Black color move")
	errBlack := client.Call("ConnectGame.Move", &moveBlack, &replyB)
	if errBlack != nil {
		log.Println("RPC error:", errBlack)
	}
	time.Sleep(10 * time.Second)
}
