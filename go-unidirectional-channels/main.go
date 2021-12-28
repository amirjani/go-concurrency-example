package main

func sendData(senderCh chan<- int) {
	senderCh <- 10
}

func main() {
	sendCh := make(chan<- int)
	go sendData(sendCh)

	// * panic
	// fmt.Println(<-sendCh)
}
