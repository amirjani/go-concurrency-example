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
	firstChannel := make(chan string)
	secondChannel := make(chan string)

	go firstChannelHandler(firstChannel)
	go secondChannelHandler(secondChannel)

	select {
	case message := <-firstChannel:
		fmt.Println(message)
	case messageTwo := <-secondChannel:
		fmt.Println(messageTwo)
	}

	<-time.After(time.Second * 2)
}
