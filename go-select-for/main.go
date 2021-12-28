package main

import (
	"fmt"
	"time"
)

func firstChannelHandler(ch chan string) {
	ch <- "First Channel Handler"
}

func secondChannelHandler(ch chan string) {
	ch <- "Second Channel Handler"
}

func main() {
	chOne := make(chan string)
	chTwo := make(chan string)

	go firstChannelHandler(chOne)
	go secondChannelHandler(chTwo)

	for i := 0; i < 2; i++ {
		select {
		case messageOne := <-chOne:
			fmt.Print(messageOne)
		case messageTwo := <-chTwo:
			fmt.Println(messageTwo)
		}
	}

	<-time.After(time.Second * 2)
}
