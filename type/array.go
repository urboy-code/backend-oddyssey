package main

import "fmt"
func Array(){
	var arr [2]string
	arr[0] = "I"
	arr[1] = "am"
	fmt.Println("Array example: ", arr[0], arr[1])

	primes := [5]int{2,3,5,7,11}
	fmt.Println("Array of primes: ", primes)
}