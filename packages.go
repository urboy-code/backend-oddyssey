package main

import (
	"fmt"
	"math/rand"
	
)

func main(){
	fmt.Println("Welcome to the Oddyssey package!")
	// rand.Seed(42) // Seed the random number generator
	fmt.Println("Random number generated:", rand.Intn(100)) // Generate a random number between
	// 0 and 99
	hasil := Add(5,7)
	fmt.Println("The result of adding 5 and 7 is:", hasil)

	a,b := swap("Hello", "World");
	fmt.Println("Swapped values:", a, b)

	result1, result2 := split(100)
	fmt.Println("Split results:", result1, result2)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}