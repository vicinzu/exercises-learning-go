package main

import "fmt"

func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var s1 = greetings[:2]
	var s2 = greetings[1:4]
	var s3 = greetings[3:]

	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s3 = %v\n", s3)
}
