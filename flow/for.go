package main

import "fmt"

func main(){
	sum := 0
	for i := 0; i <10; i++{
		sum += i
	}

	fmt.Println("Sum of first 10 numbers is:", sum)

	fmt.Println("Conditional check:", sqrt(16), sqrt(-16))

	fmt.Println("Power calculation:", pow(2, 3, 10), pow(2, 4, 10))
}