package main

import (
	"fmt"
	"time"
)

func number(start chan bool) {
	<-start
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("%d ", i)
	}
}

func alphabets(start chan bool) {
	<-start
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("%c ", i)
	}
}

func main() {
	start := make(chan bool)

	go number(start)
	go alphabets(start)

	time.Sleep(time.Second * 2)
	close(start)
	fmt.Println("Channel closed after 2 second")

	<-time.After(time.Second * 5)
}
