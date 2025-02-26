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

	fmt.Printf("Random Numbers (%d): %v\n", len(randomNumbers), randomNumbers)
}
