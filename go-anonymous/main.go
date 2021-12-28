package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", i)
			time.Sleep(time.Millisecond * 250)
		}
	}()

	go func() {
		for i := 'a'; i < 'j'; i++ {
			fmt.Printf("%c ", i)
			time.Sleep(time.Millisecond * 250)
		}
	}()

	<-time.After(time.Second * 5)
}
