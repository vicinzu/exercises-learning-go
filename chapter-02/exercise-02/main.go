package main

import "fmt"

const (
	value = 20
)

func main() {
	var i int64 = value
	var f float64 = value

	fmt.Printf("i = %d, f = %f\n", i, f)
}
