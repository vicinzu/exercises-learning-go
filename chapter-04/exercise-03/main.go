package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i // Bug: variable total is shadowed by the new variable total; Fix: total += i
		fmt.Println(total)
	}

	fmt.Printf("Total: %d\n", total)
}
