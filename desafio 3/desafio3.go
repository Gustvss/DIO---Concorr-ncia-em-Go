package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string, 2)

	for {
	canal <- "Ping"
	canal <- "Pong"

	msg1 := <-canal
	msg2 := <-canal

	fmt.Println(msg1)
	time.Sleep(time.Second * 1)
	fmt.Println(msg2)
	time.Sleep(time.Second * 1)
	}
	close(canal)

	fmt.Println("Programa encerrado!")
}
