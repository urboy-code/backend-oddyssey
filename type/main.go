package main

import "fmt"

func main(){
	fmt.Println("Struct example: ", Vertex{1, 2})

	v := Vertex{3, 4}
	fmt.Println("Accessing fields: ", v.x, v.y)

	// Array example
	Array()

	// Slice example
	Slices()

	// Slice pointers example
	SlicePointers()
}