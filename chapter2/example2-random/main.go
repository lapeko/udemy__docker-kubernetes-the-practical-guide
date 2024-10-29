package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	var minNum int
	var maxNum int
	fmt.Println("Enter min number")
	_, err := fmt.Scanln(&minNum)
	if err != nil {
		log.Fatalln(fmt.Errorf("error reading input: %w", err))
		return
	}
	fmt.Println("Enter max number")
	_, err = fmt.Scanln(&maxNum)
	if err != nil {
		log.Fatalln(fmt.Errorf("error reading input: %w", err))
		return
	}
	fmt.Printf("Random number: %d\n", minNum+rand.Intn(maxNum-minNum+1))
}
