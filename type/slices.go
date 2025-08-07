package main

import "fmt"

func Slices(){
	primes := []int{2,3,5,7,11,13}

	s := primes[3:5]
	fmt.Println("Slice of primes: ", s)
}