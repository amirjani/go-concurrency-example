package main

import "fmt"

func sender(ch chan<- string, message string) {
	ch <- message
}

func receiver(ch <-chan string) string {
	message := <-ch
	return message
}

func direction(sender chan<- string, receiver <-chan string) {
	message := <-receiver
	sender <- message
}

func main() {
	message := "Spiderman: No Way Home was amazing"
	firstChan := make(chan string, 1)
	sender(firstChan, message)

	founded := receiver(firstChan)
	fmt.Println(founded)
}
