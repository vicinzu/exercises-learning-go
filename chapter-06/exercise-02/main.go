package main

import "fmt"

func NewSlice() []string {
	return []string{"a", "b", "c", "d", "e"}
}

func UpdateSlice(slice []string, lastElement string) {
	slice[len(slice)-1] = lastElement
	fmt.Printf("Updated slice: %v\n", slice)
}

func GrowSlice(slice []string, lastElement string) {
	slice = append(slice, lastElement)
	fmt.Printf("Grown slice: %v\n", slice)
}

func main() {
	var slice = NewSlice()
	fmt.Printf("Original slice before UpdateSlice: %v\n", slice)
	UpdateSlice(slice, "z")
	fmt.Printf("Original slice after UpdateSlice: %v\n", slice)
	// Change visible in the original slice because the slice is a reference to the underlying array.
	// The length and capacity of the slice are the same.

	slice = NewSlice()
	fmt.Printf("Original slice before GrowSlice: %v\n", slice)
	GrowSlice(slice, "z")
	fmt.Printf("Original slice after GrowSlice: %v\n", slice)
	// Change not visible in the original slice because the slice is a reference to the underlying array,
	//  but the length and capacity of the slice are different. The outer slice has still the same length and capacity.
}
