package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)

	// producer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		i := 1
		for _ = range ticker.C {
			fmt.Printf("Write data into messages channel: %d\n", i)
			messages <- i
			i++
		}
	}()

	// consumer
	go func() {
		ticker1 := time.NewTicker(1 * time.Second)
		for _ = range ticker1.C {
			select {
			case <-messages:
				a := <-messages
				fmt.Printf("            Read data from messages channel: %d\n", a)
			default:
				fmt.Printf("messages channel is null...\n")
			}
		}
	}()

	time.Sleep(time.Second * 20)
	fmt.Println("main process exit!")
}
