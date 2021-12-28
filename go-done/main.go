package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello to go routine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)

	<-done
	fmt.Println("main function")
}
