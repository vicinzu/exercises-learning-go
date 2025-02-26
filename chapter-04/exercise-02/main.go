package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randomNumbers = make([]uint8, 100)
	for i := 0; i < 100; i++ {
		randomNumbers[i] = uint8(rand.Intn(101))
	}

	for i, v := range randomNumbers {
		fmt.Printf("%d (%d): ", i, v)
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind!")
		}
	}
}
