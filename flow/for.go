package main

import(
	"fmt"
	"runtime"
	"time"
)

func main(){
	sum := 0
	for i := 0; i <10; i++{
		sum += i
	}

	fmt.Println("Sum of first 10 numbers is:", sum)

	fmt.Println("Conditional check:", sqrt(16), sqrt(-16))

	fmt.Println("Power calculation:", pow(2, 3, 10), pow(2, 4, 10))

	fmt.Println("Root calculation: ", Sqrt(16))

	// Swtich case example
	fmt.Print("Go is running on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	// Switch case evalutaion order
	fmt.Println("When is Thursday?")
	today := time.Now().Weekday()
	switch time.Thursday{
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch case with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Defer statement example
	defer fmt.Println("This will be printed last.")
	fmt.Println("This will be printed first.")

	// Defer multiple statements
	fmt.Println("Counting")
	for i := 0; i<10; i++{
		defer fmt.Println(i)
	}

	fmt.Println("Counting done.")
}