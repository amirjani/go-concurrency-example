package main

import (
	"fmt"
	"time"
)

func main() {
	marvelHeros := []string{"IronMan", "Hulk", "Thor", "Captain America", "Black Widow", "Hawkeye", "SpiderMan"}
	DCHeros := []string{"Batman", "Superman", "WonderWoman", "Flash", "GreenLantern", "Aquaman", "Cyborg"}

	go listOfMarvelHeroes(marvelHeros)
	go listOfDCHeroes(DCHeros)

	<-time.After(time.Second * 10)
}

func listOfMarvelHeroes(marvelHeros []string) {
	for _, hero := range marvelHeros {
		fmt.Println("marvel hero is: ", hero)
		time.Sleep(time.Second)
	}
}

func listOfDCHeroes(dcHeros []string) {
	for _, hero := range dcHeros {
		fmt.Println("dc hero is: ", hero)
		time.Sleep(time.Second)
	}
}
