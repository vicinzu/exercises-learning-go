package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	fmt.Printf("Before:\tb = %d, smallI = %d, bigI = %d\n", b, smallI, bigI)

	b++
	smallI++
	bigI++
	fmt.Printf("After:\tb = %d, smallI = %d, bigI = %d\n", b, smallI, bigI)
}
