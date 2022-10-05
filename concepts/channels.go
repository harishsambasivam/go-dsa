package concepts

import (
	"fmt"
	"time"
)

func LoopWithRoutines() {
	// sheep()
	// goat()

	// go sheep()
	// goat()

	sheeps := make(chan string)
	goats := make(chan string)

	go goat(goats)
	go sheep(sheeps)

	for {
		select {
		case g, open := <-goats:
			fmt.Println(g, open)
		case s, open := <-sheeps:
			fmt.Println(s, open)
		}
	}
}

func sheep(sheeps chan string) {
	for i := 0; i < 5; i++ {
		sheeps <- "sheep"
		time.Sleep(5 * time.Second)
	}
	close(sheeps)
}

func goat(goats chan string) {
	for i := 0; i < 5; i++ {
		goats <- "goat"
		time.Sleep(1 * time.Second)
	}
	close(goats)
}
