package main

import "fmt"
import "time"


type Ball struct{ hits int }

func player(name string, table chan *Ball) {
	for {
		ball := <-table // réception de la balle
		ball.hits++    // enregistrement du nombre de coups
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond) // attente
		table <- ball // "renvoi" de la balle
	}
}

func main() {
 table := make(chan *Ball)

 go player("ping", table)
 go player("pong", table)

  table <- new(Ball) // game on; toss the ball // open goroutine
 time.Sleep(1 * time.Second) 
 <-table // game over; grab the ball        // close goroutine
}
