package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := [3]string{"Ore1", "Ore2", "Ore3"}
	oreChan := make(chan string)

	go func(mine [3]string) {
		for _, item := range theMine {
			oreChan <- item
		}
	}(theMine)

	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}
	}()

	<-time.After(time.Second * 5)

}
